package middleware

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"go-postgres/logger"
	"go-postgres/models"
	"go-postgres/postgres"
	"net/http"
)

type JwtToken struct {
	ID    int64  `json:"id,string"`
	Token string `json:"token"`
}

func Authorize(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err_handling(err, "invalid request body", w) != nil {
		return
	}
	user, err = postgres.GetUserWithNameAndPassword(user.Login, user.Password)
	if access_err_handling(err, "user not found", w) != nil {
		return
	}
	user_group, _ := postgres.GetGroupWithID(user.Role)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"permissions": user_group.Permissions,
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err_handling(err, "bad token string", w) != nil {
		return
	}
	logger.Info.Println("user " + user.Login + " successfully authorized")
	json.NewEncoder(w).Encode(JwtToken{ID: user.ID, Token: tokenString})
}
