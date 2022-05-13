package middleware

import (
	"encoding/json"
	"errors"
	"go-postgres/datetime"
	"go-postgres/logger"
	"go-postgres/models"
	"go-postgres/postgres"
	"math"
	"net/http"
	"time"
)

func GetAllProfessions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	professions, err := postgres.GetAllProfessions()
	if err_handling(err, "can`t get all professions", w) != nil {
		return
	}
	logger.Info.Println("get all professions response send")
	json.NewEncoder(w).Encode(professions)
}

func GetProfessionCostsForPeriod(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var data models.ProfessionForPeriod
	err := json.NewDecoder(r.Body).Decode(&data)
	if err_handling(err, "unable to decode the request body", w) != nil {
		return
	}
	// check profession
	profession, err := postgres.GetProfessionWithID(data.Profession.ID)
	if err_handling(err, "profession does not exists", w) != nil {
		return
	}
	// get all labor_costs
	labor_costs, err := postgres.GetAllLaborCostsWithProfessionID(profession.ID)
	if err_handling(err, "can`t get labor costs for profession", w) != nil {
		return
	}
	// get all products for profession
	professions_need, err := postgres.GetAllProfessionsNeedWithProfessionID(profession.ID)
	if err_handling(err, "can`t get professions need for progession", w) != nil {
		return
	}
	var products_with_profession []models.Product
	for _, profession_need := range professions_need {
		product, err := postgres.GetProductWithID(profession_need.ProductID)
		if err_handling(err, "can`t get product with profession", w) != nil {
			return
		}
		products_with_profession = append(products_with_profession, product)
	}
	// all calculations
	calculations, err := postgres.GetAllCalculations()
	if err_handling(err, "problem during calculations passing", w) != nil {
		return
	}

	// for each labor cost in range
	var resp_data []models.LaborCostForProduct
	var start_date, end_date time.Time
	if data.StartDate == "" {
		start_date = datetime.GetLowDateLimit()
	} else {
		start_date, err = datetime.DateStringToTime(data.StartDate)
		if err_handling(err, "invalid start date", w) != nil {
			return
		}
	}
	if data.EndDate == "" {
		end_date = datetime.GetHighDateLimit()
	} else {
		end_date, err = datetime.DateStringToTime(data.EndDate)
		if err_handling(err, "invalid end date", w) != nil {
			return
		}
	}
	for _, labor_cost := range labor_costs {
		date, err := datetime.DateStringToTime(labor_cost.ActualDate)
		if err_handling(err, "bad labor cost date in database, try to ask administrator", w) != nil {
			return
		}
		if datetime.IsInRange(start_date, end_date, date) {
			for _, product_with_profession := range products_with_profession {
				is_write := false
				for _, calculation := range calculations {
					if calculation.StartDate[:10] == labor_cost.ActualDate[:10] &&
						calculation.ProductID == product_with_profession.ID &&
						(product_with_profession.ID == data.Product.ID || data.Product.ID == 0) {
						is_write = true
						break
					}
				}
				if is_write {
					resp_data = append(resp_data, models.LaborCostForProduct{ProductName: product_with_profession.Name,
						ActualDate: labor_cost.ActualDate, Cost: labor_cost.Cost})
				}
			}
		}
	}
	logger.Info.Println("profession costs for product for period send")
	json.NewEncoder(w).Encode(resp_data)
}

func professionCostsProcessing(packed_calculation models.PackedCalculation) error {
	for _, profession := range packed_calculation.Product.Professions {
		var updated_id int64 = -1
		labor_costs, err := postgres.GetAllLaborCostsWithProfessionID(profession.ID)
		if err != nil {
			return err
		}
		for _, labor_cost := range labor_costs {
			if labor_cost.ActualDate[:10] == packed_calculation.StartDate[:10] {
				updated_id = labor_cost.ID
				break
			}
		}
		// TODO: fix magic constant NormID = 2
		labor_cost := models.LaborCost{ActualDate: packed_calculation.StartDate, Cost: profession.Cost, NormID: 2, ProfessionID: profession.ID}
		if updated_id == -1 { // create new cost for profession
			_, err = postgres.InsertLaborCost(labor_cost)
			if err != nil {
				return err
			}
		} else { // update cost for product
			_, err = postgres.UpdateLaborCost(updated_id, labor_cost)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func getActualLaborCost(profession_id int64, start_date_str string) (models.LaborCost, error) {
	// all costs
	labor_costs, err := postgres.GetAllLaborCostsWithProfessionID(profession_id)
	if err != nil {
		return models.LaborCost{}, err
	}
	// search for the nearest date to the beginning of the period
	start_date, err := time.Parse(datetime.Layout, start_date_str[:10])
	if err != nil {
		return models.LaborCost{}, err
	}
	nearest_labor_cost_index, min_delta := -1, -1.0
	for labor_cost_index, labor_cost := range labor_costs {
		actual_date, err := time.Parse(datetime.Layout, labor_cost.ActualDate[:10])
		if err != nil {
			return models.LaborCost{}, err
		}
		delta := math.Abs(start_date.Sub(actual_date).Hours())
		if nearest_labor_cost_index == -1 || delta < min_delta {
			nearest_labor_cost_index = labor_cost_index
			min_delta = delta
		}
	}
	if nearest_labor_cost_index == -1 {
		return models.LaborCost{}, errors.New("there is no cost for professions")
	}
	return labor_costs[nearest_labor_cost_index], nil
}

func getAllProfessionsForProduct(product_id int64) ([]models.Profession, error) {
	professions_need, err := postgres.GetAllProfessionsNeedWithProductID(product_id)
	if err != nil {
		return nil, err
	}
	var professions []models.Profession
	for _, profession_need := range professions_need {
		profession, err := postgres.GetProfessionWithID(profession_need.ProfessionID)
		if err != nil {
			return nil, err
		}
		professions = append(professions, profession)
	}
	return professions, nil
}
