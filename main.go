package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"fmt"
	"encoding/json"
	"github.com/gorilla/handlers"
	"os"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var NOTYET = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not implemented!"))
})

type Product struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Description string `json:"description"`
}

var products = []Product{
	Product{Id: 1, Name: "Hover Shooters", Slug: "hover-shooters", Description: "Shoot your way to the top on 14 different hoverboards"},
	Product{Id: 2, Name: "Ocean Explorer", Slug: "ocean-explorer", Description: "Explore the depths of the sea in this one of a kind underwater experience"},
	Product{Id: 3, Name: "Dinosaur Park", Slug: "dinosaur-park", Description: "Go back 65 million years in the past and ride a T-Rex"},
	Product{Id: 4, Name: "Cars VR", Slug: "cars-vr", Description: "Get behind the wheel of the fastest cars in the world."},
	Product{Id: 5, Name: "Robin Hood", Slug: "robin-hood", Description: "Pick up the bow and arrow and master the art of archery"},
	Product{Id: 6, Name: "Real World VR", Slug: "real-world-vr", Description: "Explore the seven wonders of the world in VR"},
}

var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("API is up and running"))
})

var ProductsHandler = http.HandlerFunc(func(w http.ResponseWriter, r* http.Request) {
	payload, _ := json.Marshal(products)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
})

var AddFeedBackHanlder = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var product Product

	vars := mux.Vars(r)
	slug := vars["slug"]

	for _, p := range products {
		if p.Slug == slug {
			product = p
		}
	}

	w.Header().Set("Content-Type", "application/json")

	if product.Slug != "" {
		payload, _ := json.Marshal(product)
		w.Write([]byte(payload))
	} else {
		w.Write([]byte("Product Not Found"))
	}
})

var mySigninKey = []byte("secret")

type MyCustomClaims struct {
	User string
	Admin bool
}

var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter,
	r *http.Request){
	// Создаем новый токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"admin": true,
		"name": "sunlive",
		"nbf": time.Now().Add(time.Hour * 24).Unix(),
	})


	// Подписываем токен нашим секретным ключем
	tokenString, err := token.SignedString(mySigninKey)
	if err != nil {
		panic(err)
	}

	// Отдаем токен клиенту
	w.Write([]byte(tokenString))
})

func RouterHandlers(router *mux.Router)  {
	router.Handle("/status", StatusHandler).Methods("GET")
	router.Handle("/products", ProductsHandler).Methods("GET")
	router.Handle("/products/{slug}/feedback", AddFeedBackHanlder).Methods("POST")
	router.Handle("/token", GetTokenHandler).Methods("GET")
}

func main()  {
	router := mux.NewRouter()
	RouterHandlers(router)

	router.Handle("/", http.FileServer(http.Dir("./views/")))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))


	fmt.Println("Server is running at port 8081")
	log.Fatal(http.ListenAndServe(":8081", handlers.LoggingHandler(os.Stdout, router)))
}