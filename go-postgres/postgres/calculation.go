package postgres

import (
	"errors"
	"go-postgres/logger"
	"go-postgres/models"
)

func UpdateCalculation(id int64, calculation models.Calculation) (int64, error) {
	db := createConnection()
	defer db.Close()
	sqlStatement := `UPDATE calculation SET start_date=$2, end_date=$3, id_product=$4 WHERE id=$1`
	res, err := db.Exec(sqlStatement, id, calculation.StartDate, calculation.EndDate, calculation.ProductID)
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

func InsertCalculation(calculation models.Calculation, isCheckForeignKeys bool) (int64, error) {
	if isCheckForeignKeys {
		_, err := GetProductWithID(calculation.ProductID)
		if err != nil {
			return -1, errors.New("product foreign key does not exists")
		}
	}

	db := createConnection()
	defer db.Close()

	sqlStatement := `INSERT INTO calculation (start_date, end_date, id_product) VALUES ($1, $2, $3) RETURNING id`
	var id int64
	err := db.QueryRow(sqlStatement, calculation.StartDate, calculation.EndDate, calculation.ProductID).Scan(&id)
	if err != nil {
		return -1, err
	}
	logger.Info.Printf("Inserted a single record %v", id)
	return id, nil
}

func DeleteCalculation(id int64, cascade bool) (int64, error) {
	db := createConnection()
	defer db.Close()

	sqlStatement := `DELETE FROM calculation WHERE id=$1`
	if cascade {
		sqlStatement = `DELETE FROM calculation CASCADE WHERE id=$1`
	}
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		return -1, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return -1, err
	}
	logger.Info.Printf("Total calculation rows/record affected %v", rowsAffected)
	return rowsAffected, nil
}

func DeletePeriodForItem(calculation_id int64, cost_item_id int64) (int64, error) {
	db := createConnection()
	defer db.Close()

	sqlStatement := `DELETE FROM period_for_item WHERE id_cost_item=$1 AND id_calc=$2`
	res, err := db.Exec(sqlStatement, cost_item_id, calculation_id)
	if err != nil {
		return -1, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return -1, err
	}
	logger.Info.Printf("Total calculation rows/record affected %v", rowsAffected)
	return rowsAffected, nil
}

func GetCalculationWithID(id int64) (models.Calculation, error) {
	db := createConnection()
	defer db.Close()
	var calculation models.Calculation
	sqlStatement := `SELECT * FROM calculation WHERE id=$1`
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&calculation.ID, &calculation.StartDate, &calculation.EndDate, &calculation.ProductID)
	return calculation, err
}

func GetAllPeriodForItemWithCalculationID(calculation_id int64) ([]models.PeriodForItem, error) {
	db := createConnection()
	defer db.Close()
	var periodForItems []models.PeriodForItem
	sqlStatement := `SELECT * FROM period_for_item WHERE id_calc=$1`
	rows, err := db.Query(sqlStatement, calculation_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var periodForItem models.PeriodForItem
		err = rows.Scan(&periodForItem.Cost, &periodForItem.CostItemID, &periodForItem.CalculationID)
		if err != nil {
			return nil, err
		}
		periodForItems = append(periodForItems, periodForItem)
	}
	return periodForItems, nil
}

func GetPeriodForItemWithCalculationIDAndCostItemID(calculation_id int64, cost_item_id int64) (models.PeriodForItem, error) {
	db := createConnection()
	defer db.Close()
	var period_for_item models.PeriodForItem
	sqlStatement := `SELECT * FROM period_for_item WHERE id_cost_item=$1 AND id_calc=$2`
	row := db.QueryRow(sqlStatement, cost_item_id, calculation_id)
	err := row.Scan(&period_for_item.Cost, &period_for_item.CostItemID, &period_for_item.CalculationID)
	return period_for_item, err
}

func GetAllCalculations() ([]models.Calculation, error) {
	db := createConnection()
	defer db.Close()

	var calculations []models.Calculation
	sqlStatement := `SELECT * FROM calculation`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var calculation models.Calculation
		err = rows.Scan(&calculation.ID, &calculation.StartDate, &calculation.EndDate, &calculation.ProductID)
		if err != nil {
			return nil, err
		}
		calculations = append(calculations, calculation)
	}
	return calculations, nil
}
