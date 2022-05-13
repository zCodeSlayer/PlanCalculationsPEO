package middleware

import (
	"encoding/json"
	"go-postgres/logger"
	"go-postgres/models"
	"go-postgres/postgres"
	"net/http"
)

func GetAllPackedProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	products, err := postgres.GetAllProducts()
	if err_handling(err, "can`t get all products", w) != nil {
		return
	}

	var packed_products []models.PackedProduct
	for _, product := range products {
		packed_product, err := getPackedProductWithID(product.ID)
		if err_handling(err, "can`t get packed product with id", w) != nil {
			return
		}
		packed_products = append(packed_products, packed_product)
	}

	logger.Info.Println("all packed products response send")
	json.NewEncoder(w).Encode(packed_products)
}

func getPackedProductWithID(id int64) (models.PackedProduct, error) {
	var packed_product models.PackedProduct
	product, err := postgres.GetProductWithID(id)
	if err != nil {
		return packed_product, err
	}
	packed_product = models.PackedProduct{ID: product.ID, Name: product.Name}
	materials_for_product, err := postgres.GetAllMaterialsForProduct(id)
	if err != nil {
		return packed_product, err
	}
	packed_product.Materials = materials_for_product
	professions_for_product, err := getAllProfessionsForProduct(id)
	if err != nil {
		return packed_product, err
	}
	packed_product.Professions = professions_for_product
	return packed_product, nil
}
