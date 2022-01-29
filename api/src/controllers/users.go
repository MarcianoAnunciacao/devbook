package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Create an user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var user models.User

	if error = json.Unmarshal(requestBody, &user); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.UserRepository(db)
	userId, error := repository.Create(user)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusCreated, userId)
	w.Write([]byte(fmt.Sprintf("User has been created with ID -> %d", userId)))
}

//Search all users
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching Users"))
}

//Search an user from his ID
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching User"))
}

//Update user with new informations
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating User"))
}

//Delete a user from the system
func RemoveUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Removing User"))
}
