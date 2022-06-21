package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

var users []User

func main() {
	fmt.Println("hello!")
	users = append(users, User{ID: "1", Name: "John Smith"})
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/users", getUsers).Methods("GET")
	r.HandleFunc("/api/v1/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/api/v1/users", createUser).Methods("POST")
	r.HandleFunc("/api/v1/users/{id}", updateUser).Methods("PUT")
	http.ListenAndServe(":8080", r)
}

