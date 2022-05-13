package middleware

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"go-postgres/logger"
	"go-postgres/models"
	"go-postgres/postgres"
	"net/http"
	"strconv"
)

func UpdateCalculation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err_handling(err, "bad calculation id", w) != nil {
		return
	}
	var packed_calculation models.PackedCalculation
	err = json.NewDecoder(r.Body).Decode(&packed_calculation)
	if err_handling(err, "bad request body", w) != nil {
		return
	}
	_, err = updateCalculationWithPackedCalculation(int64(id), packed_calculation)
	if err_handling(err, "bad calculation update", w) != nil {
		return
	}
	logger.Info.Println("calculation updated successfully")
	json.NewEncoder(w).Encode(with_id_response{
		ID:      int64(id),
		Message: "calculation updated successfully",
	})
}

func updateCalculationWithPackedCalculation(id int64, packed_calculation models.PackedCalculation) (int64, error) {
	// verify product
	product, err := postgres.GetProductWithID(packed_calculation.Product.ID)
	if err != nil {
		return -1, err
	}

	// process material and labor costs
	err = professionCostsProcessing(packed_calculation)
	if err != nil {
		return -1, err
	}
	err = materialCostsProcessing(packed_calculation)
	if err != nil {
		return -1, err
	}

	// process cost_items
	// 1. delete intermediate tables
	periods_for_item, err := postgres.GetAllPeriodForItemWithCalculationID(packed_calculation.ID)
	if err != nil {
		return -1, err
	}
	for _, period_for_item := range periods_for_item {
		_, err := postgres.DeletePeriodForItem(packed_calculation.ID, period_for_item.CostItemID)
		if err != nil {
			return -1, errors.New("fatal error during database working, check the safety of your data and contact your system administrator")
		}
	}
	// 2. rewrite period for items
	for _, packed_period_for_item := range packed_calculation.CostItems {
		new_period_for_item := models.PeriodForItem{
			Cost:          packed_period_for_item.Cost,
			CostItemID:    packed_period_for_item.ID,
			CalculationID: packed_calculation.ID}
		err := postgres.InsertPeriodForItem(new_period_for_item, true)
		if err != nil { // if bad insert period_for_item then
			_, err := postgres.DeleteCalculation(packed_calculation.ID, true) // delete all editable calculation
			if err != nil {
				panic(err)
			}
			return -1, errors.New("database data was corrupted with bad period for item, check database, calculation was removed with intermediate tables")
		}
	}
	rowsAffected, err := postgres.UpdateCalculation(id, models.Calculation{ID: id, StartDate: packed_calculation.StartDate,
		EndDate: packed_calculation.EndDate, ProductID: product.ID})
	if err != nil {
		return -1, err
	}
	return rowsAffected, nil
}

func CreateNewCalculation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var packed_calculation models.PackedCalculation
	err := json.NewDecoder(r.Body).Decode(&packed_calculation)
	if err_handling(err, "unable to decode the request body", w) != nil {
		return
	}

	// 1. process costs
	err = professionCostsProcessing(packed_calculation)
	if err_handling(err, "bad labor costs processing", w) != nil {
		return
	}
	err = materialCostsProcessing(packed_calculation)
	if err_handling(err, "bad material costs processing", w) != nil {
		return
	}

	// new calculation
	new_calculation := models.Calculation{ID: -1, StartDate: packed_calculation.StartDate,
		EndDate: packed_calculation.EndDate, ProductID: packed_calculation.Product.ID}
	new_calculation_id, err := postgres.InsertCalculation(new_calculation, true)
	if err_handling(err, "can`t insert calculation", w) != nil {
		return
	}

	// new cost_item
	cost_items := packed_calculation.CostItems
	for _, cost_item := range cost_items {
		new_period_for_item := models.PeriodForItem{Cost: cost_item.Cost, CostItemID: cost_item.ID, CalculationID: new_calculation_id}
		err := postgres.InsertPeriodForItem(new_period_for_item, true)
		if err != nil { // if bad insert cost_item then
			_, err := postgres.DeleteCalculation(new_calculation_id, true) // delete the one you just added calculation
			if err_handling(err, "fatal error, the server will be shut down, please contact your system administrator", w) != nil {
				panic(err)
			}
			err_handling(errors.New("bad period for item"), "can`t create new cost_item, operation canceled", w) // return error response
			return
		}
	}

	// response part
	logger.Info.Println("calculation created successfully")
	json.NewEncoder(w).Encode(with_id_response{
		ID:      new_calculation_id,
		Message: "calculation created successfully",
	})
}

func DeleteCalculation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err_handling(err, "unable to convert the string into int", w) != nil {
		return
	}
	deletedRows, err := postgres.DeleteCalculation(int64(id), true)
	if err_handling(err, "deleting have special error (maybe not critical)", w) != nil {
		return
	}
	logger.Info.Printf("calculation deleted successfully, total rows/record affected %v", deletedRows)
	json.NewEncoder(w).Encode(with_id_response{
		ID:      int64(id),
		Message: "calculation deleted successfully",
	})
}

func GetAllPackedCalculations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var packed_calculations []models.PackedCalculation
	calculations, err := postgres.GetAllCalculations()
	if err_handling(err, "couldn't request a list of all the calculations", w) != nil {
		return
	}

	// for each calculation
	for _, calculation := range calculations {
		packed_calculation, err := getPackedCalculationWithID(calculation.ID)
		if err_handling(err, "can`t get packed calculation", w) != nil {
			return
		}
		packed_calculations = append(packed_calculations, packed_calculation)
	}

	logger.Info.Println("get all packed calculations response send")
	json.NewEncoder(w).Encode(packed_calculations)
}

func GetPackedCalculationWithID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err_handling(err, "unable to convert the string into int", w) != nil {
		return
	}
	packed_calculation, err := getPackedCalculationWithID(int64(id))
	if err_handling(err, "can`t get packed calculation", w) != nil {
		return
	}
	logger.Info.Println("get packed calculation response send")
	json.NewEncoder(w).Encode(packed_calculation)
}

func getPackedCalculationWithID(id int64) (models.PackedCalculation, error) {
	var packed_calculation models.PackedCalculation
	calculation, err := postgres.GetCalculationWithID(id)
	if err != nil {
		return packed_calculation, err
	}

	// calculation part
	packed_calculation.ID = calculation.ID
	packed_calculation.StartDate = calculation.StartDate
	packed_calculation.EndDate = calculation.EndDate

	// product part
	calculation_product, err := postgres.GetProductWithID(calculation.ProductID)
	if err != nil {
		return packed_calculation, err
	}
	packed_product := models.PackedProductWithActualCosts{ID: calculation_product.ID, Name: calculation_product.Name}

	// product-materials part
	total__materials_cost := 0.0
	materials_need, err := postgres.GetAllMaterialsNeedWithProductID(packed_product.ID)
	if err != nil {
		return packed_calculation, err
	}
	var materials_with_actual_cost []models.MaterialWithActualCost
	for _, material_need := range materials_need {
		material, err := postgres.GetMaterialWithID(material_need.MaterialID)
		if err != nil {
			return packed_calculation, err
		}
		material_with_actual_cost := models.MaterialWithActualCost{ID: material.ID, Name: material.Name}

		material_cost, err := getActualMaterialCost(material.ID, calculation.StartDate)
		if err != nil {
			return packed_calculation, err
		}
		material_with_actual_cost.Cost = material_cost.Cost
		materials_with_actual_cost = append(materials_with_actual_cost, material_with_actual_cost)
		total__materials_cost += material_with_actual_cost.Cost * material_need.Consumption
	}
	packed_product.Materials = materials_with_actual_cost

	// product-professions part
	total_professions_cost := 0.0
	professions_need, err := postgres.GetAllProfessionsNeedWithProductID(packed_product.ID)
	if err != nil {
		return packed_calculation, err
	}
	var professions_with_actual_cost []models.ProfessionWithActualCost
	for _, profession_need := range professions_need {
		profession, err := postgres.GetProfessionWithID(profession_need.ProfessionID)
		if err != nil {
			return packed_calculation, err
		}
		profession_with_actual_cost := models.ProfessionWithActualCost{ID: profession.ID, Description: profession.Description}

		labor_cost, err := getActualLaborCost(profession.ID, calculation.StartDate)
		if err != nil {
			return packed_calculation, err
		}
		profession_with_actual_cost.Cost = labor_cost.Cost
		professions_with_actual_cost = append(professions_with_actual_cost, profession_with_actual_cost)
		total_professions_cost += profession_with_actual_cost.Cost
	}
	packed_product.Professions = professions_with_actual_cost

	packed_calculation.Product = packed_product

	// cost items part
	total_cost := 0.0
	periods_for_item, err := postgres.GetAllPeriodForItemWithCalculationID(calculation.ID)
	if err != nil {
		return packed_calculation, err
	}
	var packed_cost_items []models.PackedCostItem
	for _, period_for_item := range periods_for_item {
		cost_item, err := postgres.GetCostItemWithID(period_for_item.CostItemID)
		if err != nil {
			return packed_calculation, err
		}
		packed_cost_item := models.PackedCostItem{ID: cost_item.ID, Name: cost_item.Name, Cost: period_for_item.Cost}
		total_cost += period_for_item.Cost
		packed_cost_items = append(packed_cost_items, packed_cost_item)
	}
	packed_calculation.CostItems = packed_cost_items

	// calculated cost item part
	calculated_cost_items := []models.CalculatedCostItem{
		{Name: "затраты на материалы", Cost: total__materials_cost},
		{Name: "затраты на заработную плату", Cost: total_professions_cost},
	}
	packed_calculation.CalculatedCostItems = calculated_cost_items

	// full_cost part
	packed_calculation.FullCost = total_cost + total__materials_cost + total_professions_cost

	return packed_calculation, nil
}
