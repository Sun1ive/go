package main

import (
	"fmt"
	"net/http"
)

func GetUsers(res http.ResponseWriter, req *http.Request)  {
	fmt.Fprintf(res, "All users endpoint hit")
}

func NewUser(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "User endpoint hit")
}

func DeleteUser(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "User delete endpoint hit")
}

func UpdateUser(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Update endpoint hit")
}