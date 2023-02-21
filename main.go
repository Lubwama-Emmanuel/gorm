package main

import (
	"fmt"
	"net/http"

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
func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint for deleting a user")
}

// func intialMigration() {
// 	gorm.Open()
// }

func HandleRequests() {
	http.HandleFunc("/users", allUsers)
	http.HandleFunc("/users/{name}", deleteUser)
	http.HandleFunc("/users/{name}/{email}", createUser)
	_ = http.ListenAndServe(":8080", nil)

}

func main() {
	fmt.Println("Started")

	HandleRequests()
}
