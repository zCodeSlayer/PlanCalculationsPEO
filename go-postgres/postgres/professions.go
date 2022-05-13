package postgres

import (
	"go-postgres/logger"
	"go-postgres/models"
)

func UpdateLaborCost(id int64, labor_cost models.LaborCost) (int64, error) {
	db := createConnection()
	defer db.Close()
	sqlStatement := `UPDATE labor_cost SET actual_date=$2, cost=$3, id_norm=$4, id_profession=$5 WHERE id=$1`
	res, err := db.Exec(sqlStatement, id, labor_cost.ActualDate, labor_cost.Cost, labor_cost.NormID, labor_cost.ProfessionID)
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

func InsertLaborCost(labor_cost models.LaborCost) (int64, error) {
	db := createConnection()
	defer db.Close()
	sqlStatement := `INSERT INTO labor_cost (actual_date, cost, id_norm, id_profession) VALUES ($1, $2, $3, $4) RETURNING id`
	var id int64
	err := db.QueryRow(sqlStatement, labor_cost.ActualDate, labor_cost.Cost, labor_cost.NormID, labor_cost.ProfessionID).Scan(&id)
	if err != nil {
		return -1, err
	}
	logger.Info.Printf("inserted a single labor_cost record %v", id)
	return id, nil
}

func GetAllLaborCostsWithProfessionID(profession_id int64) ([]models.LaborCost, error) {
	db := createConnection()
	defer db.Close()
	var labor_costs []models.LaborCost
	sqlStatement := `SELECT * FROM labor_cost WHERE id_profession=$1`
	rows, err := db.Query(sqlStatement, profession_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var labor_cost models.LaborCost
		err = rows.Scan(&labor_cost.ID, &labor_cost.ActualDate, &labor_cost.Cost, &labor_cost.NormID, &labor_cost.ProfessionID)
		if err != nil {
			return nil, err
		}
		labor_costs = append(labor_costs, labor_cost)
	}
	return labor_costs, nil
}

func GetProfessionWithID(id int64) (models.Profession, error) {
	db := createConnection()
	defer db.Close()
	var profession models.Profession
	sqlStatement := `SELECT * FROM profession WHERE id=$1`
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&profession.ID, &profession.Description, &profession.Cipher)
	return profession, err
}

func GetAllProfessionsNeedWithProfessionID(profession_id int64) ([]models.ProfessionsNeed, error) {
	db := createConnection()
	defer db.Close()
	var professions_need []models.ProfessionsNeed
	sqlStatement := `SELECT * FROM professions_need WHERE id_profession=$1`
	rows, err := db.Query(sqlStatement, profession_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var profession_need models.ProfessionsNeed
		err = rows.Scan(&profession_need.ProductID, &profession_need.ProfessionID)
		if err != nil {
			return nil, err
		}
		professions_need = append(professions_need, profession_need)
	}
	return professions_need, nil
}

func GetAllProfessionsNeedWithProductID(product_id int64) ([]models.ProfessionsNeed, error) {
	db := createConnection()
	defer db.Close()
	var professions_need []models.ProfessionsNeed
	sqlStatement := `SELECT * FROM professions_need WHERE id_product=$1`
	rows, err := db.Query(sqlStatement, product_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var profession_need models.ProfessionsNeed
		err = rows.Scan(&profession_need.ProductID, &profession_need.ProfessionID)
		if err != nil {
			return nil, err
		}
		professions_need = append(professions_need, profession_need)
	}
	return professions_need, nil
}

func GetAllProfessions() ([]models.Profession, error) {
	db := createConnection()
	defer db.Close()

	var professions []models.Profession
	sqlStatement := `SELECT * FROM profession`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var profession models.Profession
		err = rows.Scan(&profession.ID, &profession.Description, &profession.Cipher)
		if err != nil {
			return nil, err
		}
		professions = append(professions, profession)
	}
	return professions, nil
}
