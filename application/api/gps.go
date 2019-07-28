package api

import (
	"net/http"

	"github.com/local/app-gps/application/models"
)

func GetAllGPS(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	data, err := models.GetAllGPS(w, r)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetGPS(r http.ResponseWriter, h *http.Request) (interface{}, error) {
	data, err := models.GetGPS(r, h)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func InsertGPS(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	_, err := models.InsertGPS(w, r)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func UpdateGPS(r http.ResponseWriter, h *http.Request) (interface{}, error) {
	_, err := models.UpdateGPS(r, h)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func DeleteGPS(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	message, err := models.DeleteGPS(w, r)
	if err != nil {
		return nil, err
	}
	var convert = map[string]interface{}{"message": message}
	return convert, nil
}
