package controllers

import (
	"fmt"
	"log"
	"luwa-server/utils"
	"net/http"
	"strings"

	"luwa-server/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("password")
	if email == "" || password == "" {
		http.Error(w, "Invalid email or password", http.StatusBadRequest)
		return
	}
	// Create a new user
	newUser := models.User{
		Email:    email,
		Password: password,
	}
	err := newUser.Create()
	if err != nil {
		http.Error(w, "Error creating a new user", http.StatusInternalServerError)
		return
	}

	log.Println("New user created successfully")
	http.Redirect(w, r, fmt.Sprintf("/user/%s", email), http.StatusSeeOther)
}

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		http.Error(w, "Invalid URL path", http.StatusBadRequest)
		return
	}
	log.Println("pathParts:", pathParts)

	userEmail := pathParts[2]
	log.Println("userEmail:", userEmail)

	m := models.User{}
	user, err := m.GetUserByEmail(userEmail)
	if err != nil {
		http.Error(w, "Error getting user", http.StatusInternalServerError)
		return
	}
	log.Println("User:", user)

	data := struct {
		User models.User
	}{
		User: *user,
	}

	utils.Render(w, "user.page.html", data)
	return
}
