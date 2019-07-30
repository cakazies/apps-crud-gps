package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/local/app-gps/utils"
	"github.com/spf13/viper"
)

var (
	GetDB *gorm.DB // connection about DB
)

func Connect() {
	host := viper.GetString("configDB.host")
	port := viper.GetString("configDB.port")
	user := viper.GetString("configDB.user")
	password := viper.GetString("configDB.password")
	dbname := viper.GetString("configDB.dbname")

	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
	conn, err := gorm.Open("postgres", dbUri)
	utils.FailError(err, "Connection Error")
	GetDB = conn
}

func ReturnDBGorm() *gorm.DB {
	return GetDB
}
