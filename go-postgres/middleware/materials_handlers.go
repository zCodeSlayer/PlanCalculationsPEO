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

func GetAllMaterials(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	materials, err := postgres.GetAllMaterials()
	if err_handling(err, "can`t get all materials", w) != nil {
		return
	}
	logger.Info.Println("get all materials response send")
	json.NewEncoder(w).Encode(materials)
}

func GetMaterialCostsForPeriod(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var data models.MaterialForPeriod
	err := json.NewDecoder(r.Body).Decode(&data)
	if err_handling(err, "unable to decode the request body", w) != nil {
		return
	}

	material, err := postgres.GetMaterialWithID(data.Material.ID)
	if err_handling(err, "product does not exists", w) != nil {
		return
	}
	material_costs, err := postgres.GetAllMaterialCostsWithMaterialID(material.ID)
	if err_handling(err, "can`t get material costs for material", w) != nil {
		return
	}

	var resp_data []models.PackedMaterialCost
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
	for _, material_cost := range material_costs {
		date, err := datetime.DateStringToTime(material_cost.ActualDate)
		if err_handling(err, "bad labor cost date in database, try to ask administrator", w) != nil {
			return
		}
		if datetime.IsInRange(start_date, end_date, date) {
			resp_data = append(resp_data, models.PackedMaterialCost{
				MaterialName: material.Name,
				ActualDate:   material_cost.ActualDate,
				Cost:         material_cost.Cost})
		}
	}
	logger.Info.Println("material costs for period send")
	json.NewEncoder(w).Encode(resp_data)
}

func materialCostsProcessing(packed_calculation models.PackedCalculation) error {
	for _, material := range packed_calculation.Product.Materials {
		var updated_id int64 = -1
		material_costs, err := postgres.GetAllMaterialCostsWithMaterialID(material.ID)
		if err != nil {
			return err
		}
		for _, material_cost := range material_costs {
			if material_cost.ActualDate[:10] == packed_calculation.StartDate[:10] {
				updated_id = material_cost.ID
				break
			}
		}
		// TODO: fix magic constant NormID = 3
		material_cost := models.MaterialCost{ActualDate: packed_calculation.StartDate, Cost: material.Cost, NormID: 3, MaterialID: material.ID}
		if updated_id == -1 { // create new cost for profession
			_, err = postgres.InsertMaterialCost(material_cost)
			if err != nil {
				return err
			}
		} else { // update cost for product
			_, err = postgres.UpdateMaterialCost(updated_id, material_cost)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func getActualMaterialCost(material_id int64, start_date_str string) (models.MaterialCost, error) {
	// all costs
	material_costs, err := postgres.GetAllMaterialCostsWithMaterialID(material_id)
	if err != nil {
		return models.MaterialCost{}, err
	}
	// search for the nearest date to the beginning of the period
	start_date, err := time.Parse(datetime.Layout, start_date_str[:10])
	if err != nil {
		return models.MaterialCost{}, err
	}
	nearest_material_cost_index, min_delta := -1, -1.0
	for material_cost_index, material_cost := range material_costs {
		actual_date, err := time.Parse(datetime.Layout, material_cost.ActualDate[:10])
		if err != nil {
			return models.MaterialCost{}, err
		}
		delta := math.Abs(start_date.Sub(actual_date).Hours())
		if nearest_material_cost_index == -1 || delta < min_delta {
			nearest_material_cost_index = material_cost_index
			min_delta = delta
		}
	}
	if nearest_material_cost_index == -1 {
		return models.MaterialCost{}, errors.New("there is no cost for material")
	}
	return material_costs[nearest_material_cost_index], nil
}
