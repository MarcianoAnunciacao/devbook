package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
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
	savedUser, err := repository.SearchByEmail(user.Email)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
	}

	if err = security.ValidatePassword(savedUser.Password, user.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	token, err := authentication.CreateToken(savedUser.ID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
	}

	userID := strconv.FormatUint(savedUser.ID, 10)

	responses.JSON(w, http.StatusOK, models.AuthData{ID: userID, Token: token})
}
