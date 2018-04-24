package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"fmt"
)

var NOTYET = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not implemented!"))
})

func RouterHandlers(router *mux.Router)  {
	router.Handle("/status", NOTYET).Methods("GET")
	router.Handle("/products", NOTYET).Methods("GET")
	router.Handle("/products/{slug}/feedback", NOTYET).Methods("POST")
}

func main()  {
	router := mux.NewRouter()

	RouterHandlers(router)

	router.Handle("/", http.FileServer(http.Dir("./views/")))

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("./static/"))))

	fmt.Println("Server is running at port 8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}