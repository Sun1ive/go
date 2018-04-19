package main

import (
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

// Book Struct // Model // interface ? // class ?
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *author `json:"author"`
}

// Author struct
type author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// init books var as slice book struct
var books []Book

// set Header func

func setHeaders(res http.ResponseWriter) {
	res.Header().Set("Content-Type", "application/json")
}

// get all books
func getBooks(res http.ResponseWriter, req *http.Request) {
	// res.Header().Set("Content-Type", "application/json")
	setHeaders(res)
	json.NewEncoder(res).Encode(books)
}

// get book
func getBook(res http.ResponseWriter, req *http.Request) {
	setHeaders(res)
	params := mux.Vars(req) //get Params
	// Loop through and find id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(res).Encode(item)
			return
		}
	}
	json.NewEncoder(res).Encode(&Book{})
}

// createBook
func createBook(res http.ResponseWriter, req *http.Request) {
	setHeaders(res)
	var book Book
	_ = json.NewDecoder(req.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(res).Encode(book)
}

// edit book
func updateBook(res http.ResponseWriter, req *http.Request) {

}

// delete book
func deleteBook(res http.ResponseWriter, req *http.Request) {

}

func main() {
	// init router
	r := mux.NewRouter()

	// mock data // todo: implement database
	books = append(books, Book{ID: "1", Isbn: "455432", Title: "Book One", Author: &author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "661342", Title: "Book Two", Author: &author{Firstname: "Doe", Lastname: "John"}})
	books = append(books, Book{ID: "3", Isbn: "548763", Title: "Book Three", Author: &author{Firstname: "Steve", Lastname: "Smith"}})

	// create our route handlers endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))

}
