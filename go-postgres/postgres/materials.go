package postgres

import (
	"go-postgres/logger"
	"go-postgres/models"
)

func UpdateMaterialCost(id int64, material_cost models.MaterialCost) (int64, error) {
	db := createConnection()
	defer db.Close()
	sqlStatement := `UPDATE material_cost SET actual_date=$2, cost=$3, id_norm=$4, id_material=$5 WHERE id=$1`
	res, err := db.Exec(sqlStatement, id, material_cost.ActualDate, material_cost.Cost, material_cost.NormID, material_cost.MaterialID)
	if err != nil {
		return -1, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return -1, err
	}
	logger.Info.Println(rowsAffected, " rows was affected")
	return rowsAffected, nil
}

func InsertMaterialCost(material_cost models.MaterialCost) (int64, error) {
	db := createConnection()
	defer db.Close()
	sqlStatement := `INSERT INTO material_cost (actual_date, cost, id_norm, id_material) VALUES ($1, $2, $3, $4) RETURNING id`
	var id int64
	err := db.QueryRow(sqlStatement, material_cost.ActualDate, material_cost.Cost, material_cost.NormID, material_cost.MaterialID).Scan(&id)
	if err != nil {
		return -1, err
	}
	logger.Info.Printf("inserted a single material_cost record %v", id)
	return id, nil
}

func GetAllMaterialsForProduct(product_id int64) ([]models.Material, error) {
	materials_need, err := GetAllMaterialsNeedWithProductID(product_id)
	if err != nil {
		return nil, err
	}
	var materials []models.Material
	for _, material_need := range materials_need {
		material, err := GetMaterialWithID(material_need.MaterialID)
		if err != nil {
			return nil, err
		}
		materials = append(materials, material)
	}
	return materials, nil
}

func GetAllMaterialCostsWithMaterialID(material_id int64) ([]models.MaterialCost, error) {
	db := createConnection()
	defer db.Close()
	var material_costs []models.MaterialCost
	sqlStatement := `SELECT * FROM material_cost WHERE id_material=$1`
	rows, err := db.Query(sqlStatement, material_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var material_cost models.MaterialCost
		err = rows.Scan(&material_cost.ID, &material_cost.ActualDate, &material_cost.Cost, &material_cost.NormID, &material_cost.MaterialID)
		if err != nil {
			return nil, err
		}
		material_costs = append(material_costs, material_cost)
	}
	return material_costs, nil
}

func GetMaterialWithID(id int64) (models.Material, error) {
	db := createConnection()
	defer db.Close()
	var material models.Material
	sqlStatement := `SELECT * FROM material WHERE id=$1`
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&material.ID, &material.Cipher, &material.Name)
	return material, err
}

func GetAllMaterialsNeedWithProductID(product_id int64) ([]models.MaterialsNeed, error) {
	db := createConnection()
	defer db.Close()
	var materials_need []models.MaterialsNeed
	sqlStatement := `SELECT * FROM materials_need WHERE id_product=$1`
	rows, err := db.Query(sqlStatement, product_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var material_need models.MaterialsNeed
		err = rows.Scan(&material_need.Consumption, &material_need.NormID, &material_need.MaterialID, &material_need.ProductID)
		if err != nil {
			return nil, err
		}
		materials_need = append(materials_need, material_need)
	}
	return materials_need, nil
}

func GetAllMaterials() ([]models.Material, error) {
	db := createConnection()
	defer db.Close()

	var materials []models.Material
	sqlStatement := `SELECT * FROM material`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var material models.Material
		err = rows.Scan(&material.ID, &material.Cipher, &material.Name)
		if err != nil {
			return nil, err
		}
		materials = append(materials, material)
	}
	return materials, nil
}
