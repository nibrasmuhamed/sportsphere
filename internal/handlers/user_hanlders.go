package handlers

// TO MOVED
import (
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func ListUsers(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("List of users"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User created"))
}
