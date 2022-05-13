package middleware

import (
	"encoding/json"
	"go-postgres/logger"
	"go-postgres/postgres"
	"net/http"
)

func GetAllCostItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	cost_items, err := postgres.GetAllCostItems()
	if err_handling(err, "can`t get all cost items", w) != nil {
		return
	}
	logger.Info.Println("all cost items response send")
	json.NewEncoder(w).Encode(cost_items)
}
