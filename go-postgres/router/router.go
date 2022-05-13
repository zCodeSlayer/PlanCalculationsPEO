package router

import (
	"github.com/gorilla/mux"
	"go-postgres/middleware"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// common
	router.HandleFunc("/api/ping", middleware.Ping).Methods("GET", "OPTIONS")
	// authenticate
	router.HandleFunc("/api/auth", middleware.Authorize).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/super_user/ping",
		middleware.ValidateUser(middleware.Ping, []string{})).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/reader/ping",
		middleware.ValidateUser(middleware.Ping, []string{"read only permissions"})).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/redactor/ping",
		middleware.ValidateUser(middleware.Ping, []string{"limited crud"})).Methods("GET", "OPTIONS")
	// roles
	router.HandleFunc("/api/roles/create_user",
		middleware.ValidateUser(middleware.CreateUser, []string{})).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/roles/update_user/{id}",
		middleware.ValidateUser(middleware.UpdateUser, []string{})).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/roles/get_user_id",
		middleware.ValidateUser(middleware.GetUserIDWithNameAndPassword, []string{})).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/roles/check_readonly/{id}", middleware.CheckReadonly).Methods("GET", "OPTIONS")
	// calculations
	router.HandleFunc("/api/calculations/get", middleware.GetAllPackedCalculations).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/calculations/get/{id}", middleware.GetPackedCalculationWithID).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/calculations/create",
		middleware.ValidateUser(middleware.CreateNewCalculation, []string{"limited crud"})).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/calculations/update/{id}",
		middleware.ValidateUser(middleware.UpdateCalculation, []string{"limited crud"})).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/calculations/delete/{id}",
		middleware.ValidateUser(middleware.DeleteCalculation, []string{"limited crud"})).Methods("DELETE", "OPTIONS")
	// products
	router.HandleFunc("/api/products/get", middleware.GetAllPackedProducts).Methods("GET", "OPTIONS")
	// cost items
	router.HandleFunc("/api/cost_items/get", middleware.GetAllCostItems).Methods("GET", "OPTIONS")
	// directories
	router.HandleFunc("/api/directories/get", middleware.GetAllDirectories).Methods("GET", "OPTIONS")
	// professions
	router.HandleFunc("/api/professions/get", middleware.GetAllProfessions).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/professions/get/range", middleware.GetProfessionCostsForPeriod).Methods("POST", "OPTIONS")
	// materials
	router.HandleFunc("/api/materials/get", middleware.GetAllMaterials).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/materials/get/range", middleware.GetMaterialCostsForPeriod).Methods("POST", "OPTIONS")
	return router
}
