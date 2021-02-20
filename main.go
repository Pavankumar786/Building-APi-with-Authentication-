package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/login", Login)
	r.HandleFunc("/user", Home)
	r.HandleFunc("/refresh", Refresh)
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":9079", r))
}
func main() {
	InitialMigration()
	initializeRouter()

}