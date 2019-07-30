package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	cf "github.com/local/app-gps/application/models"
	"github.com/local/app-gps/cmd"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	cmd.InitViper()
	var limit int
	limit = 10
	cf.Connect()
	MigrationGPS(limit)
	MigrationUser(limit)
}

// MigrationGPS function for migration tabel GPS and insert dummy data
func MigrationGPS(limit int) {
	tableName := "gps"
	drop := fmt.Sprintf("DROP TABLE IF EXISTS %s;", tableName)

	cf.GetDB.Exec(drop)
	queryCreate := fmt.Sprintf(`
					CREATE TABLE public.%s
					(
						id SERIAL NOT NULL,
						brand character varying(100) COLLATE pg_catalog."default" NOT NULL,
						models character varying(100) COLLATE pg_catalog."default" NOT NULL,
						name character varying(400) COLLATE pg_catalog."default" NOT NULL,
						waranty character varying(400) COLLATE pg_catalog."default" NOT NULL,
						date_buy timestamp without time zone NULL,
						date_sold timestamp without time zone NULL,
						sold_to character varying(200) COLLATE pg_catalog."default" NOT NULL,
						foto character varying(200) COLLATE pg_catalog."default" NULL,
						description character varying(1000) COLLATE pg_catalog."default" NULL,
						created_at timestamp without time zone NOT NULL,
						updated_at timestamp without time zone ,
						deleted_at timestamp without time zone ,
						CONSTRAINT %s_pk PRIMARY KEY (id)
					);`, tableName, tableName)
	cf.GetDB.Exec(queryCreate)
	log.Println(fmt.Sprintf("Import Table %s Succesfull", tableName))

	for i := 1; i <= limit; i++ {
		brand := "brand-" + strconv.Itoa(i)
		model := "model-" + strconv.Itoa(i)
		name := "name-" + strconv.Itoa(i)
		waranty := "waranty-" + strconv.Itoa(i)
		sold_to := "sold_to-" + strconv.Itoa(i)
		foto := "foto-" + strconv.Itoa(i)
		description := "description-" + strconv.Itoa(i)
		date_buy := time.Now().Format("2006-01-02 15:04:05")
		date_sold := time.Now().Format("2006-01-02 15:04:05")
		createdAt := time.Now().Format("2006-01-02 15:04:05")
		sql := fmt.Sprintf("INSERT INTO %s ( brand, models, name, waranty, date_buy, date_sold, sold_to, foto, description , created_at) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s'); ",
			tableName, brand, model, name, waranty, date_buy, date_sold, sold_to, foto, description, createdAt)
		cf.GetDB.Exec(sql)
		time.Sleep(time.Second / 10)
	}
	log.Println(fmt.Sprintf("Insert Data Dummy table %s successfull", tableName))
}

// MigrationUser function for migration table user
func MigrationUser(limit int) {
	tableName := "users"
	drop := fmt.Sprintf("DROP TABLE IF EXISTS %s;", tableName)
	cf.GetDB.Exec(drop)
	queryCreate := fmt.Sprintf(`
					CREATE TABLE public.%s
					(
						id SERIAL NOT NULL,
						email character varying(200) COLLATE pg_catalog."default" NOT NULL,
						username character varying(100) COLLATE pg_catalog."default" NOT NULL,
						password character varying(250) COLLATE pg_catalog."default" NOT NULL,
						created_at timestamp without time zone,
						updated_at timestamp without time zone,
						deleted_at timestamp without time zone,
						status integer ,
						CONSTRAINT %s_pk PRIMARY KEY (id)
					);`, tableName, tableName)
	cf.GetDB.Exec(queryCreate)
	log.Println(fmt.Sprintf("Import Table %s Succesfull", tableName))

	for i := 1; i <= limit; i++ {
		email := "admin@gmail.com"
		username := "admin"
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		password := string(hashedPassword)
		createdAt := time.Now().Format("2006-01-02 15:04:05")
		status := "1"
		if i != 1 {
			status = "2"
			email = "email_" + strconv.Itoa(i) + "@gmail.com"
			username = "username-" + strconv.Itoa(i)
		}
		sql := fmt.Sprintf("INSERT INTO %s ( email, username, password, created_at, status) VALUES ( '%s', '%s', '%s', '%s', '%s'); ",
			tableName, email, username, password, createdAt, status)
		cf.GetDB.Exec(sql)
		time.Sleep(time.Second / 10)
	}
	log.Println(fmt.Sprintf("Insert Data Dummy table %s successfull", tableName))
}
