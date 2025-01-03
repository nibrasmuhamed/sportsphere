package handlers

// TO MOVED
import (
	"net/http"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List of users"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User created"))
}
