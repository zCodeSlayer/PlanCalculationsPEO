package middleware

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-postgres/logger"
	"go-postgres/models"
	"go-postgres/postgres"
	"go-postgres/validators"
	"net/http"
	"strconv"
)

func CreateGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var group models.Group

	err := json.NewDecoder(r.Body).Decode(&group)
	if err_handling(err, "unable to decode the request body", w) != nil ||
		err_handling(validators.GroupValidation(group), "invalid the request body", w) != nil {
		return
	}

	insertID, err := postgres.InsertGroup(group)
	if err_handling(err, "unable to create a group", w) != nil {
		return
	}

	json.NewEncoder(w).Encode(with_id_response{
		ID:      insertID,
		Message: "group created successfully",
	})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err_handling(err, "unable to decode the request body", w) != nil ||
		err_handling(validators.UserValidation(user), "invalid the request body", w) != nil {
		return
	}

	insertID, err := postgres.InsertUser(user)
	if err_handling(err, "unable to create a user", w) != nil {
		return
	}

	json.NewEncoder(w).Encode(with_id_response{
		ID:      insertID,
		Message: "user created successfully",
	})
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err_handling(err, "unable to convert the string into int", w) != nil {
		return
	}
	var user models.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err_handling(err, "unable to decode the request body", w) != nil {
		return
	}
	updatedRows, err := postgres.UpdateUser(int64(id), user)
	if err_handling(err, "unable to execute the query", w) != nil {
		return
	}
	logger.Info.Println("updated ", updatedRows, " rows")
	json.NewEncoder(w).Encode(with_id_response{
		ID:      int64(id),
		Message: "successful",
	})
}

func CheckReadonly(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err_handling(err, "unable to convert the string into int", w) != nil {
		return
	}
	user, err := postgres.GetUserWithID(int64(id))
	if err_handling(err, "couldn't get user by id", w) != nil {
		return
	}
	user_group, _ := postgres.GetGroupWithID(user.Role)
	readonly := 1
	if user_group.Name != "reader" {
		readonly = 0
	}
	json.NewEncoder(w).Encode(flag_response{F: readonly})
}

func GetUserIDWithNameAndPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err_handling(err, "unable to decode the request body", w) != nil {
		return
	}
	user, err = postgres.GetUserWithNameAndPassword(user.Login, user.Password)
	if err_handling(err, "not found", w) != nil {
		return
	}
	json.NewEncoder(w).Encode(with_id_response{
		ID:      user.ID,
		Message: "successfully",
	})
}
