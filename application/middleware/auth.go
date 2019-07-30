package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/local/app-gps/application/models"
	"github.com/spf13/viper"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		noAuthPath := []string{"/api/user/register"}
		noAuthPath = append(noAuthPath, "/api/user/login")
		noAuthPath = append(noAuthPath, "/api/user/logout")
		requestPath := r.URL.Path
		// looping for check pathnya
		for _, path := range noAuthPath {
			if path == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}
		tokenHeader := r.Header.Get("Authorization")

		// check for token present or not
		if tokenHeader == "" {
			rsp := map[string]interface{}{"status": "invalid", "message": "Token is not Present ;"}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(rsp)
			return
		}

		// check format auth token valid or not
		headerAuthorizationString := strings.Split(tokenHeader, " ")
		if len(headerAuthorizationString) != 2 {
			rsp := map[string]interface{}{"status": "invalid", "message": "Invalid/Format Auth Token ;"}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(rsp)
			return
		}

		// type token barier or not
		barier := headerAuthorizationString[0]
		if barier != "Barier" {
			rsp := map[string]interface{}{"status": "invalid", "message": "Token is not Barier ;"}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(rsp)
			return
		}

		tk := &models.Token{}
		tokenValue := headerAuthorizationString[1]
		token, err := jwt.ParseWithClaims(tokenValue, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("api.secret_key")), nil
		})

		// somthing went wrong, about token, please reset token
		if err != nil {
			rsp := map[string]interface{}{"status": "invalid", "message": "Malformed Authentication Token Please Login Again;"}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(rsp)
			return
		}

		// check admin or not
		checkAdmin := make(map[string]string)
		checkAdmin["GET"] = "/api/user/1"
		checkAdmin["GET"] = "/api/user"
		checkAdmin["DELETE"] = "/api/user"

		for key, value := range checkAdmin {
			if key == r.Method && value == requestPath && tk.UserId != 1 {
				rsp := map[string]interface{}{"status": "invalid", "message": "This is for admin users"}
				w.Header().Add("Content-Type", "application/json")
				json.NewEncoder(w).Encode(rsp)
				return
			}
		}

		// check for time expired
		diff := tk.TimeExp.Sub(time.Now())
		if diff < 0 {
			rsp := map[string]interface{}{"status": "invalid", "message": "Time Expired, please login again;"}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(rsp)
			return
		}

		// check format token valid or not
		if !token.Valid {
			rsp := map[string]interface{}{"status": "invalid", "message": "Invalid/Format Auth Token ;"}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(rsp)
			return
		}

		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)

	})
}
