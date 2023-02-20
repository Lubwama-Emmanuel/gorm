package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gollira/mux"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name  string
	Email string
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint for all users")
}
func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint for creating users")
}
func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint for updating a user")
}
func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint for deleting a user")
}

func HandleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/users", allUsers).Method("GET")
	r.HandleFunc("/users/{name}", deleteUser).Method("DELETE")
	r.HandleFunc("/users/{name}/{email}", createUser).Method("POST")
	r.HandleFunc("/users/{name}/{email}", updateUser).Method("PUT")
	log.Fatal(http.ListenAndServe(":8080", r))

}

func main() {

}
