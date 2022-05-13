package postgres

import (
	"errors"
	"go-postgres/logger"
	"go-postgres/models"
)

func InsertPeriodForItem(period_for_item models.PeriodForItem, isCheckForeignKeys bool) error {
	if isCheckForeignKeys {
		_, err := GetCostItemWithID(period_for_item.CostItemID)
		if err != nil {
			return errors.New("cost_item foreign key does not exists")
		}
		_, err = GetCalculationWithID(period_for_item.CalculationID)
		if err != nil {
			return errors.New("calculation foreign key does not exists")
		}
	}

	db := createConnection()
	defer db.Close()

	sqlStatement := `INSERT INTO period_for_item (cost, id_cost_item, id_calc) VALUES ($1, $2, $3) RETURNING cost`
	var cost int64
	err := db.QueryRow(sqlStatement, period_for_item.Cost, period_for_item.CostItemID, period_for_item.CalculationID).Scan(&cost)

	if err != nil {
		return err
	}
	logger.Info.Printf("Inserted a new period for item")
	return nil
}

func GetCostItemWithID(id int64) (models.CostItem, error) {
	db := createConnection()
	defer db.Close()
	var cost_item models.CostItem
	sqlStatement := `SELECT * FROM cost_item WHERE id=$1`
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&cost_item.ID, &cost_item.Name, &cost_item.Cipher)
	return cost_item, err
}

func GetAllCostItems() ([]models.CostItem, error) {
	db := createConnection()
	defer db.Close()

	var cost_items []models.CostItem
	sqlStatement := `SELECT * FROM cost_item`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var cost_item models.CostItem
		err = rows.Scan(&cost_item.ID, &cost_item.Name, &cost_item.Cipher)
		if err != nil {
			return nil, err
		}
		cost_items = append(cost_items, cost_item)
	}
	return cost_items, nil
}
