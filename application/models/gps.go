package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/local/app-gps/utils"
)

// struct rooms
type Gps struct {
	Brand       string `json:"brand,omitemp"`
	Models      string `json:"models,omitemp"`
	Name        string `json:"name,omitemp"`
	Waranty     string `json:"waranty,omitemp"`
	DateBuy     string `json:"date_buy,omitemp"`
	DateSold    string `json:"date_sold,omitemp"`
	SoldTo      string `json:"sold_to,omitemp"`
	Foto        string `json:"foto,omitemp"`
	Description string `json:"description,omitemp"`
	gorm.Model
}

// strtuck for many room array
type ManyGps []Gps

// GetRooms function for get all data from table
func GetAllGPS(w http.ResponseWriter, r *http.Request) ([]*Gps, error) {
	// qulimit := ""
	// quShort := ""
	// qulimit = LimitOffset(r.URL.Query().Get("limit"), r.URL.Query().Get("offset"))
	// quShort = ShortBy(r.URL.Query().Get("sort_by"))

	value := make([]*Gps, 0)
	err := GetDB.Table("gps").Find(&value).Error
	if err != nil {
		saveError := fmt.Sprintf("Error query data, and %s", err)
		return nil, errors.New(saveError)
	}
	return value, nil
}

// // GetRoom functionfor get perRooms
func GetGPS(w http.ResponseWriter, r *http.Request) (*Gps, error) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	url := strings.Split((r.URL.String()), "/")
	if id == 0 {
		id, _ = strconv.Atoi(url[5])
	}
	if id < 1 {
		return nil, errors.New("ID Only Positive and Integer")
	}
	var value Gps

	// qulimit := ""
	// quShort := ""
	// qulimit = LimitOffset(r.URL.Query().Get("limit"), r.URL.Query().Get("offset"))
	// quShort = ShortBy(r.URL.Query().Get("sort_by"))

	GetDB.Where("id = ?", id).First(&value)
	return &value, nil
}

// UpdateRooms function for update data rooms
func UpdateGPS(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	value := &Gps{}
	err := json.NewDecoder(r.Body).Decode(&value)
	if err != nil {
		utils.FailError(err, "Error Parsing Data ")
	}
	value.ID = uint(id)

	GetDB.Save(&value)
	return nil, nil
}

// InsertGPS function for insert data in table GPS
func InsertGPS(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {
	var value Gps
	err := json.NewDecoder(r.Body).Decode(&value)
	utils.FailError(err, "Convert Error")
	GetDB.Create(&value)
	return nil, nil
}

// DeleteRoom function for delete one data in table rooms
func DeleteGPS(w http.ResponseWriter, r *http.Request) (string, error) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	url := strings.Split((r.URL.String()), "/")
	if id == 0 {
		id, _ = strconv.Atoi(url[5])
	}
	if id < 1 {
		return "", errors.New("ID Only Positive and Integer")
	}
	var value Gps
	GetDB.Where("id = ?", id).Delete(&value)

	return "Berhasil dihapus", nil
}

// CekExistData function for validate data exist
// func CekExistData(id int) error {
// 	rm_id := ""
// 	sql := "SELECT rm_id FROM rooms WHERE rm_id = $1"
// 	statement, err := DB.Prepare(sql)
// 	defer statement.Close()
// 	if err != nil {
// 		return err
// 	}
// 	err = statement.QueryRow(id).Scan(&rm_id)
// 	if err != nil {
// 		return errors.New("ID Doesn't Exist")
// 	}
// 	return nil
// }
