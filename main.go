package main

import (
	// "encoding/json"
	"log"
	"net/http"
	// "math/rand"
	// "strconv"
	"github.com/gorilla/mux"
)
// Book Struct // Model // interface ? // class ?
type Book struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	firstName string `json:"firstname"`
	lastName string `json:"lastname"`
}

// get all books
func getBooks(res http.ResponseWriter, req *http.Request)  {
	
}
// get book
func getBook(res http.ResponseWriter, req *http.Request)  {
	
}
// createBook
func createBook(res http.ResponseWriter, req *http.Request)  {
	
}
// edit book
func updateBook(res http.ResponseWriter, req *http.Request)  {
	
}
// delete book
func deleteBook(res http.ResponseWriter, req *http.Request)  {
	
}

func main() {
	// init router
	r := mux.NewRouter()

	// create our route handlers endpoints

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	
	log.Fatal(http.ListenAndServe(":8000", r))

}
