package middleware

import (
	"encoding/json"
	_ "github.com/lib/pq" // postgres golang driver
	"go-postgres/logger"
	"go-postgres/postgres"
	"net/http"
)

var Port string

func Ping(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(basic_response{Message: "ping OK, port " + Port}) // TODO: added runserver configurations
	logger.Info.Println("ping OK, port "+Port, r.RemoteAddr)
}

func GetAllDirectories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var directories []directory

	// 1. cost item
	cost_item_d := directory{Name: "Статьи затрат", Columns: []string{"Наименование", "Шифр"}, Data: [][]string{}}
	cost_items, err := postgres.GetAllCostItems()
	if err_handling(err, "can`t get all cost items", w) != nil {
		return
	}
	for _, cost_item := range cost_items {
		row := []string{cost_item.Name, cost_item.Cipher}
		cost_item_d.Data = append(cost_item_d.Data, row)
	}
	directories = append(directories, cost_item_d)

	// 2. product
	product_d := directory{Name: "Изделие", Columns: []string{"Наименование", "Шифр"}, Data: [][]string{}}
	products, err := postgres.GetAllProducts()
	if err_handling(err, "can`t get all products", w) != nil {
		return
	}
	for _, product := range products {
		row := []string{product.Name, product.Cipher}
		product_d.Data = append(product_d.Data, row)
	}
	directories = append(directories, product_d)

	// 3. profession
	profession_d := directory{Name: "Профессия", Columns: []string{"Описание", "Шифр"}, Data: [][]string{}}
	professions, err := postgres.GetAllProfessions()
	if err_handling(err, "can`t get all professions", w) != nil {
		return
	}
	for _, profession := range professions {
		row := []string{profession.Description, profession.Cipher}
		profession_d.Data = append(profession_d.Data, row)
	}
	directories = append(directories, profession_d)

	// 4. norm
	norm_d := directory{Name: "Единица измерения", Columns: []string{"Наименование", "Шифр"}, Data: [][]string{}}
	norms, err := postgres.GetAllNorms()
	if err_handling(err, "can`t get all norms", w) != nil {
		return
	}
	for _, norm := range norms {
		row := []string{norm.Name, norm.Cipher}
		norm_d.Data = append(norm_d.Data, row)
	}
	directories = append(directories, norm_d)

	// 5. material
	material_d := directory{Name: "Материал", Columns: []string{"Наименование", "Шифр"}, Data: [][]string{}}
	materials, err := postgres.GetAllMaterials()
	if err_handling(err, "can`t get all norms", w) != nil {
		return
	}
	for _, material := range materials {
		row := []string{material.Name, material.Cipher}
		material_d.Data = append(material_d.Data, row)
	}
	directories = append(directories, material_d)

	logger.Info.Println("all directories successfully send with response")
	json.NewEncoder(w).Encode(directories)
}
