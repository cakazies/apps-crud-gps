# APPS-CRUD-GPS

This repo code with golang Rest API , database Postgresql and some package in github and 2 level auth with Middleware, 

  - Cobra
  - Viper
  - Mux
  - JWT-go
  - Gorm
  - Postgres

# Feature

  - 2 Level User admin and user
  - Register New User  
  - Read all user register for admin (only) 
  - Update Password for user
  - Create new data GPS 
  - Read all data GPS 
  - Delete data GPS 
  - Update data GPS 

### Installation.

Install the package before start the server.

```sh
$ go get github.com/dgrijalva/jwt-go
$ go get github.com/spf13/viper 
$ go get golang.org/x/crypto/bcrypt
$ go get github.com/lib/pq
$ go get github.com/jinzhu/gorm
$ go get github.com/spf13/cobra
```
### Run Apps

make database name same with in configs/config.toml:
setting your config database in configs/config.toml
Run Migration table and some dummy data:
```sh
$ go run application/migration/migrate.go
```
run the apps:
```sh
$ go run main.go
```


## If you have disscus ||~
Just Contact Me At:
- Email: [cakazies@gmail.com](mailto:cakazies@gmail.com)
- Instagram: [@cakazies](https://www.instagram.com/cakazies/)

License
----
MIT

