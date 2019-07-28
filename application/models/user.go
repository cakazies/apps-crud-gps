package models

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/local/app-gps/utils"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email    string `json:"email,omitemp"`
	Username string `json:"username,omitemp"`
	Password string `json:"password,omitemp"`
	Status   string `json:"status,omitemp"`
	gorm.Model
}

type ManyUser []User

const table = "users"

// GetRooms function for get all data from table
func GetAllUserPanel(w http.ResponseWriter, r *http.Request) ([]*User, error) {
	var value []*User
	err := GetDB.Where("status = ?", "2").Find(&value).Error
	utils.FailError(err, "Error Query Data ")
	return value, nil
}

// GetUserPanel functionfor get perUser
func GetUserPanel(w http.ResponseWriter, r *http.Request) (*User, error) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	url := strings.Split((r.URL.String()), "/")
	if id == 0 {
		id, _ = strconv.Atoi(url[5])
	}
	if id < 1 {
		return nil, errors.New("ID Only Positive and Integer")
	}
	var value User

	// qulimit := ""
	// quShort := ""
	// qulimit = LimitOffset(r.URL.Query().Get("limit"), r.URL.Query().Get("offset"))
	// quShort = ShortBy(r.URL.Query().Get("sort_by"))

	GetDB.Where("id = ? AND deleted_at IS NULL", id).First(&value)
	return &value, nil
}

// UpdateUserPanel function for update data rooms
func UpdateUserPanel(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	value := &User{}
	err := GetDB.Table(table).Where("id = ?", id).First(value).Error
	utils.FailError(err, "Error query row Data ")

	err = json.NewDecoder(r.Body).Decode(&value)
	utils.FailError(err, "Error Parsing Data ")

	if value.Password != "" {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(value.Password), bcrypt.DefaultCost)
		value.Password = string(hashedPassword)
	}

	value.ID = uint(id)
	value.Status = "2"
	GetDB.Save(&value)
	return nil, nil
}

// DeleteRoom function for delete one data in table rooms
func DeleteUserPanel(w http.ResponseWriter, r *http.Request) (string, error) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	url := strings.Split((r.URL.String()), "/")
	if id == 0 {
		id, _ = strconv.Atoi(url[5])
	}
	if id < 1 {
		return "", errors.New("ID Only Positive and Integer")
	}
	var value User
	GetDB.Where("id = ?", id).Delete(&value)

	return "Berhasil dihapus", nil
}

func (user *User) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(user.Email, "@") && !strings.Contains(user.Email, ".") {
		return map[string]interface{}{"status": "invalid", "message": "Email address format is incorrect"}, false
	}
	if len(user.Password) < 6 {
		return map[string]interface{}{"status": "invalid", "message": "Password is minimum 6 character"}, false
	}
	var value User
	err := GetDB.Where("email = ?", user.Email).First(&value).Error
	if err.Error() != "record not found" {
		return map[string]interface{}{"status": "invalid", "message": "Something went wrong, please contact admin or developer."}, false
	}

	if value.Email != "" {
		return map[string]interface{}{"status": "invalid", "message": "Email address already in use by another user."}, false
	}

	return map[string]interface{}{"status": "Valid", "message": "Requirement passed"}, true
}

// asds
func (user *User) RegisterUserPanel() map[string]interface{} {
	if rsp, status := user.Validate(); !status {
		return rsp
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	user.Status = "2"

	err := GetDB.Create(user).Error
	if err != nil {
		return map[string]interface{}{"status": "invalid", "message": "Insert Errors call admin or developer "}
	}

	Hours, _ := strconv.Atoi(viper.GetString("expired.hours"))
	Mins, _ := strconv.Atoi(viper.GetString("expired.minutes"))
	timein := time.Now().Local().Add(time.Hour*time.Duration(Hours) + time.Minute*time.Duration(Mins))

	tk := &Token{UserId: uint(2), Email: user.Email, TimeExp: timein}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(viper.GetString("api.secret_key")))

	return map[string]interface{}{"status": "valid", "message": "Account is successfully created ", "token": tokenString}
}

func (registeredUser *User) Login() map[string]interface{} {
	var value User
	GetDB.Where("email = ?", registeredUser.Email).First(&value)
	if value.Email == "" {
		return map[string]interface{}{"status": "invalid", "message": "Email Invalid please try again."}
	}

	err := bcrypt.CompareHashAndPassword([]byte(value.Password), []byte(registeredUser.Password))

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return map[string]interface{}{"status": "invalid", "message": "Password Invalid."}
	}
	Hours, _ := strconv.Atoi(viper.GetString("expired.hours"))
	Mins, _ := strconv.Atoi(viper.GetString("expired.minutes"))
	timein := time.Now().Local().Add(time.Hour*time.Duration(Hours) +
		time.Minute*time.Duration(Mins))

	tk := &Token{UserId: uint(value.ID), Email: value.Email, TimeExp: timein}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(viper.GetString("api.secret_key")))

	return map[string]interface{}{"status": "valid", "message": "Login is Success", "token": tokenString}
}
