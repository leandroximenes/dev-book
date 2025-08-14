package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"main/src/router/config"
	"main/src/router/models"
	"main/src/router/repository"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating new user")
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		log.Fatal(err)
	}

	db, err := config.Connect()
	if err != nil {
		log.Fatal(err)
	}

	repository := repository.NewUserRepository(db)
	id, err := repository.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("New user created with success")
	w.Write([]byte(fmt.Sprintf("Id inserted: %d", id)))
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
