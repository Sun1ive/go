package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

func InitialMigration() {
	db, err = gorm.Open("sqlite3", "test.sqlite")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	defer db.Close()

	db.AutoMigrate(&User{})
}

func setHeaders(res http.ResponseWriter) {
	res.Header().Set("Content-Type", "application/json")
}

func connect() {
	db, err = gorm.Open("sqlite3", "test.sqlite")
	if err != nil {
		panic("Failed to connect to database")
	}
}

func GetUsers(res http.ResponseWriter, req *http.Request) {
	connect()
	defer db.Close()

	var users []User
	db.Find(&users)

	setHeaders(res)
	json.NewEncoder(res).Encode(users)
}

func NewUser(res http.ResponseWriter, req *http.Request) {
	connect()
	defer db.Close()

	vars := mux.Vars(req)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{Name: name, Email: email})

	setHeaders(res)
	fmt.Fprintf(res, "New user successfully created!")
}

func DeleteUser(res http.ResponseWriter, req *http.Request) {
	connect()
	defer db.Close()

	vars := mux.Vars(req)
	name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	setHeaders(res)
	fmt.Fprintf(res, "User successfully Deleted")
}

func UpdateUser(res http.ResponseWriter, req *http.Request) {
	connect()
	defer db.Close()

	vars := mux.Vars(req)
	name := vars["name"]
	email := vars["email"]

	var user User
	db.Where("name = ?", name).Find(&user)

	user.Email = email

	db.Save(&user)

	setHeaders(res)
	fmt.Fprintf(res, "Successfully updated")
}
