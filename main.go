package main

import (
	"encoding/json"
	"net/http"
	"os"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// memory storage
var users []User

func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST only", http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	users = append(users, user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "User added",
	})
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/create", createUser)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	http.ListenAndServe(":"+port, nil)
}
