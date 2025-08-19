package controllers

import (
	"encoding/json"
	"io"
	"log"
	"main/src/router/config"
	"main/src/router/dtos"
	"main/src/router/models"
	"main/src/router/repository"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating new user")
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		log.Println("Error deserialization request body:", err)
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := config.Connect()
	if err != nil {
		log.Println("Error connecting to database:", err)
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewUserRepository(db)
	user.ID, err = repository.Create(user)
	if err != nil {
		log.Println("Error creating user:", err)
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	log.Printf("New user created with success")
	responses.JSON(w, http.StatusCreated, user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get users"))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get user"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update user"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user"))
}
