package middleware

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"reflect"
)

func ValidateUser(next http.HandlerFunc, permissions []string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("access")
		if authorizationHeader != "" {
			claims := jwt.MapClaims{}
			_, err := jwt.ParseWithClaims(authorizationHeader, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte("secret"), nil
			})
			if err_handling(err, "bad token", w) != nil {
				return
			}
			if decodedPermissions, ok := claims["permissions"]; ok &&
				(decodedPermissions == "god permissions" || SliceExists(permissions, decodedPermissions)) {
				next(w, r)
			} else {
				json.NewEncoder(w).Encode(error_response{Err_type: "bad token",
					Message: "no access"})
			}
		} else {
			json.NewEncoder(w).Encode(error_response{Err_type: "invalid authorization header",
				Message: "an authorization header is required"})
		}
	})
}

func SliceExists(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("SliceExists() given a non-slice type")
	}
	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true
		}
	}
	return false
}
