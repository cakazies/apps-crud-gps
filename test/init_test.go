package test

import (
	"database/sql"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/local/app-gps/application/api"

	mw "github.com/local/app-gps/application/middleware"
	conf "github.com/local/app-gps/application/models"
	"github.com/local/app-gps/routes"
	"github.com/local/app-gps/utils"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	DB      *sql.DB
	Cfg     testCaseParams
)

type testCaseParams struct {
	ID  string
	URL string
}

func initCOnfig() {
	viper.SetConfigFile("toml")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("../configs")
		viper.SetConfigName("config")
	}
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	utils.FailError(err, "Error Viper config")
}

func TestInit(t *testing.T) {
	initCOnfig()
	conf.Connect()
	Cfg = testCaseParams{
		ID:  viper.GetString("testing.room_id"),
		URL: viper.GetString("app.host"),
	}
}

func getRequest(url, path string, handler routes.HandlerFunc) *httptest.ResponseRecorder {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	r := mux.NewRouter()
	r.Handle(path, routes.HandlerFunc(api.GetGPS))
	r.Handle(path, routes.HandlerFunc(handler))
	r.Use(mw.JwtAuthentication)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}

func postRequest(url, path string, handler routes.HandlerFunc) *httptest.ResponseRecorder {
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	r := mux.NewRouter()
	r.Handle(path, routes.HandlerFunc(api.GetGPS))
	r.Handle(path, routes.HandlerFunc(handler))
	r.Use(mw.JwtAuthentication)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}
