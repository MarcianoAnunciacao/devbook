package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

	if err := user.ValidateInputData("create"); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
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
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.UserRepository(db)
	users, err := repository.Search(nameOrNick)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
	}

	responses.JSON(w, http.StatusOK, users)
}

//Search an user from his ID
func SearchUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userId, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.UserRepository(db)
	user, err := repository.SearchById(userId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

//Update user with new information's
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userId, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	userIdfromToken, err := authentication.ExtractUserIDFromToken(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userId != userIdfromToken {
		responses.Error(w, http.StatusForbidden, errors.New("You are not authorized to make changes in this User"))
		return
	}

	fmt.Println(userIdfromToken)

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.ValidateInputData("update"); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.UserRepository(db)
	if err = repository.Update(userId, user); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

//Delete a user from the system
func RemoveUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userId, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	userIdfromToken, err := authentication.ExtractUserIDFromToken(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userId != userIdfromToken {
		responses.Error(w, http.StatusForbidden, errors.New("You are not authorized to delete this User"))
		return
	}

	fmt.Println(userIdfromToken)

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.UserRepository(db)
	if err = repository.Delete(userId); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func FollowAnUser(w http.ResponseWriter, r *http.Request) {
	followID, err := authentication.ExtractUserIDFromToken(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["userID"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if followID == userID {
		responses.Error(w, http.StatusForbidden, errors.New("It is not possible to follow yourself"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.UserRepository(db)
	if err = repository.Follow(userID, followID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func UnFollowAnUser(w http.ResponseWriter, r *http.Request) {
	followID, err := authentication.ExtractUserIDFromToken(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["userID"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if followID == userID {
		responses.Error(w, http.StatusForbidden, errors.New("It is not possible to unfollow yourself"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.UserRepository(db)
	if err = repository.UnFollow(userID, followID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func SearchFollowersByUserID(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.UserRepository(db)
	followes, err := repository.SearchFollowersByUserID(userID)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, followes)
}

func SearchUsersFollowedByAnUserID(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.UserRepository(db)
	users, err := repository.SearchUsersFollowedByAnUserID(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userIdfromToken, err := authentication.ExtractUserIDFromToken(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["userID"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if userID != userIdfromToken {
		responses.Error(w, http.StatusForbidden, errors.New("It is not possible to update another User passoword"))
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)

	var password models.Password

	if err = json.Unmarshal(requestBody, &password); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.UserRepository(db)
	savedPassword, err := repository.SearchPasswordByUserId(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.ValidatePassword(savedPassword, password.New); err != nil {
		responses.Error(w, http.StatusUnauthorized, errors.New("The password you entered did not match our records, Please try again"))
		return
	}

	hashedPassword, err := security.Hash(password.New)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = repository.UpdatePassword(userID, string(hashedPassword)); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
