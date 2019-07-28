package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/local/app-gps/application/models"
	"github.com/spf13/viper"
)

//
func GetAllUserPanel(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	data, err := models.GetAllUserPanel(w, r)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetUserPanel(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	data, err := models.GetUserPanel(w, r)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func UpdateUserPanel(r http.ResponseWriter, h *http.Request) (interface{}, error) {
	_, err := models.UpdateUserPanel(r, h)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func DeleteUserPanel(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	message, err := models.DeleteUserPanel(w, r)
	if err != nil {
		return nil, err
	}
	var convert = map[string]interface{}{"message": message}
	return convert, nil
}

func RegisterUserPanel(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		return map[string]interface{}{"status": "invalid", "message": "invalid parse data body"}, err
	}
	response := user.RegisterUserPanel()
	return response, nil
}

func Login(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		fmt.Println(err)
		return map[string]interface{}{"status": "invalid", "message": "invalid parse data body"}, err
	}
	response := user.Login()
	return response, nil
}

// Me function for get information about user login in JWT
func Me(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	tokenHeader := r.Header.Get("Authorization")

	headerAuthorizationString := strings.Split(tokenHeader, " ")
	token := headerAuthorizationString[1]
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("api.secret_key")), nil
	})

	if err != nil {
		log.Fatalln("errornya is : ", err)
	}
	return claims, nil
}
