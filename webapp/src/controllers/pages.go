package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responses"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.ReadCookie(r)

	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", http.StatusFound)
	}

	utils.RenderTemplate(w, "login.html", nil)
}

func LoadCreateUserPage(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "userregistration.html", nil)
}

func LoadMainPage(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publications", config.ApiUrl)
	response, err := requests.MakeARequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.FormatStatusCodeErrors(w, response)
		return
	}

	var publications []models.Publication
	if err = json.NewDecoder(response.Body).Decode(&publications); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Error: err.Error()})
		return
	}

	cookie, _ := cookies.ReadCookie(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.RenderTemplate(w, "home.html", struct {
		Publications []models.Publication
		UserID       uint64
	}{
		Publications: publications,
		UserID:       userID,
	})
}

func LoadPublicationEditionPage(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	publicationID, err := strconv.ParseUint(parameters["publicationId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications/%d", config.ApiUrl, publicationID)
	response, err := requests.MakeARequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.FormatStatusCodeErrors(w, response)
		return
	}

	var publication models.Publication
	if err = json.NewDecoder(response.Body).Decode(&publication); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Error: err.Error()})
	}

	utils.RenderTemplate(w, "update-publication.html", publication)
}

func LoadUserProfile(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	cookie, _ := cookies.ReadCookie(r)
	loggedUserID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	if userID == loggedUserID {
		http.Redirect(w, r, "/profile", http.StatusFound)
	}

	user, err := models.SearchUserWithAllInformations(userID, r)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}

	utils.RenderTemplate(w, "user.html", struct {
		User         models.User
		LoggedUserID uint64
	}{
		User:         user,
		LoggedUserID: loggedUserID,
	})
}

func LoadLoggedUserProfile(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.ReadCookie(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	user, err := models.SearchUserWithAllInformations(userID, r)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}

	utils.RenderTemplate(w, "profile.html", user)
}

func LoadUserProfileUpdate(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.ReadCookie(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	channel := make(chan models.User)
	go models.SearchUserData(channel, userID, r)
	user := <-channel

	if user.ID == 0 {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: "Error searching for User"})
		return
	}

	utils.RenderTemplate(w, "edit-user.html", user)
}

func LoadUpdatePasswordPage(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "update-password.html", nil)
}
