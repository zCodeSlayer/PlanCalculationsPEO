package postgres

import "go-postgres/models"

func GetAllNorms() ([]models.Norm, error) {
	db := createConnection()
	defer db.Close()

	var norms []models.Norm
	sqlStatement := `SELECT * FROM norm`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var norm models.Norm
		err = rows.Scan(&norm.ID, &norm.Cipher, &norm.Name)
		if err != nil {
			return nil, err
		}
		norms = append(norms, norm)
	}
	return norms, nil
}
