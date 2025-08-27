package controllers

import (
	"encoding/json"
	"io"
	"log"
	"main/src/config"
	dto "main/src/dtos"
	"main/src/models"
	"main/src/repository"
	"main/src/responses"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating new user")
	var userDto dto.UserDto
	var user models.User

	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = json.Unmarshal(bodyRequest, &userDto); err != nil {
		log.Println("Error deserialization request body:", err)
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	err = userDto.Prepare()
	if err != nil {
		log.Println("Error in body dto:", err)
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := config.Connect()
	if err != nil {
		log.Println("Error connecting to database:", err)
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	user.ConvertFromDto(userDto)

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
	nameOrSlug := strings.ToLower(r.URL.Query().Get("user"))

	db, err := config.Connect()
	if err != nil {
		log.Println("Error connecting to database:", err)
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewUserRepository(db)
	users, err := repository.GetUsers(nameOrSlug)
	if err != nil {
		log.Println("Error to find users:", err)
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	log.Printf("Users found with success")
	responses.JSON(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil || !(userId > 0) {
		log.Println("Invalid param:", err)
		responses.Error(w, http.StatusInternalServerError, err)
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
	user, err := repository.GetUser(userId)
	if err != nil {
		log.Println("Error to find user:", err)
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	log.Printf("Users found with success")
	responses.JSON(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update user"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user"))
}
