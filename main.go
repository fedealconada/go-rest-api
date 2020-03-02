package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/fedealconada/go-rest-api/app"
	"github.com/fedealconada/go-rest-api/controllers"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/signup", auth.CreateAccount).Methods("POST")
	router.HandleFunc("/api/login", auth.Authenticate).Methods("POST")
	router.HandleFunc("/api/contacts", contacts.GetContactsFor).Methods("GET")
	router.HandleFunc("/api/contacts", contacts.CreateContact).Methods("POST")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}