package routes

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/local/app-gps/application/api"
	"github.com/local/app-gps/application/middleware"
	"github.com/spf13/viper"
)

func Route() {
	r := mux.NewRouter()

	routers := r.PathPrefix("/api").Subrouter()

	// cek for middleware
	routers.Use(middleware.JwtAuthentication)

	// modul rooms
	routers.Handle("/gps", HandlerFunc(api.GetAllGPS)).Methods(http.MethodGet)
	routers.Handle("/gps", HandlerFunc(api.InsertGPS)).Methods(http.MethodPost)
	routers.Handle("/gps/{id}", HandlerFunc(api.GetGPS)).Methods(http.MethodGet)
	routers.Handle("/gps/{id}", HandlerFunc(api.UpdateGPS)).Methods(http.MethodPut)
	routers.Handle("/gps/{id}", HandlerFunc(api.DeleteGPS)).Methods(http.MethodDelete)

	// modul User
	routers.Handle("/user", HandlerFunc(api.GetAllUserPanel)).Methods(http.MethodGet)
	routers.Handle("/user", HandlerFunc(api.RegisterUserPanel)).Methods(http.MethodPost)
	routers.Handle("/user/{id}", HandlerFunc(api.GetUserPanel)).Methods(http.MethodGet)
	routers.Handle("/user/{id}", HandlerFunc(api.UpdateUserPanel)).Methods(http.MethodPut)
	routers.Handle("/user/{id}", HandlerFunc(api.DeleteUserPanel)).Methods(http.MethodDelete)
	routers.Handle("/user/login", HandlerFunc(api.Login)).Methods(http.MethodPost)
	routers.Handle("/me", HandlerFunc(api.Me)).Methods(http.MethodGet)

	host := fmt.Sprintf(viper.GetString("app.host"))

	srv := &http.Server{
		Handler:      routers,
		Addr:         host,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
