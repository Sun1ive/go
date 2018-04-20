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
	Name  string
	Email string
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

	json.NewEncoder(res).Encode(users)
}

func NewUser(res http.ResponseWriter, req *http.Request) {
	connect()
	defer db.Close()

	vars := mux.Vars(req)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{Name: name, Email: email})

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

	fmt.Fprintf(res, "Successfully updated")
}
