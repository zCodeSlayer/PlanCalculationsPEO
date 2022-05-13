package postgres

import "go-postgres/models"

func GetProductWithID(id int64) (models.Product, error) {
	db := createConnection()
	defer db.Close()
	var product models.Product
	sqlStatement := `SELECT * FROM product WHERE id=$1`
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&product.ID, &product.Name, &product.Cipher)
	return product, err
}

func GetAllProducts() ([]models.Product, error) {
	db := createConnection()
	defer db.Close()

	var products []models.Product
	sqlStatement := `SELECT * FROM product`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var product models.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Cipher)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
