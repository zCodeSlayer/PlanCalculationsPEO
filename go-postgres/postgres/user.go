package postgres

import (
	"go-postgres/logger"
	"go-postgres/models"
)

func InsertGroup(group models.Group) (int64, error) {
	db := createConnection()
	defer db.Close()

	sqlStatement := `INSERT INTO groups (name, permissions) VALUES ($1, $2) RETURNING id`
	var id int64
	err := db.QueryRow(sqlStatement, group.Name, group.Permissions).Scan(&id)

	if err != nil {
		return -1, err
	}
	logger.Info.Printf("Inserted a single record %v", id)
	return id, nil
}

func GetGroupWithName(name string) (models.Group, error) {
	db := createConnection()
	defer db.Close()

	var group models.Group
	sqlStatement := `SELECT * FROM groups WHERE name=$1`
	row := db.QueryRow(sqlStatement, name)
	err := row.Scan(&group.ID, &group.Name, &group.Permissions)
	return group, err
}

func GetGroupWithID(id int64) (models.Group, error) {
	db := createConnection()
	defer db.Close()

	var group models.Group
	sqlStatement := `SELECT * FROM groups WHERE id=$1`
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&group.ID, &group.Name, &group.Permissions)
	return group, err
}

func InsertUser(user models.User) (int64, error) {
	db := createConnection()
	defer db.Close()

	sqlStatement := `INSERT INTO users (login, password, id_role) VALUES ($1, $2, $3) RETURNING id`
	var id int64
	err := db.QueryRow(sqlStatement, user.Login, user.Password, user.Role).Scan(&id)

	if err != nil {
		return -1, err
	}
	logger.Info.Printf("Inserted a single record %v", id)
	return id, nil
}

func UpdateUser(id int64, user models.User) (int64, error) {
	db := createConnection()
	defer db.Close()
	sqlStatement := `UPDATE users SET login=$2, password=$3 WHERE id=$1`
	res, err := db.Exec(sqlStatement, id, user.Login, user.Password)
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

func GetUserWithNameAndPassword(login string, password string) (models.User, error) {
	db := createConnection()
	defer db.Close()

	var user models.User
	sqlStatement := `SELECT * FROM users WHERE login=$1 AND password=$2`
	row := db.QueryRow(sqlStatement, login, password)
	err := row.Scan(&user.ID, &user.Login, &user.Password, &user.Role)
	return user, err
}

func GetUserWithID(id int64) (models.User, error) {
	db := createConnection()
	defer db.Close()

	var user models.User
	sqlStatement := `SELECT * FROM users WHERE id=$1`
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&user.ID, &user.Login, &user.Password, &user.Role)
	return user, err
}
