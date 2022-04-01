package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responses"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"email":    r.FormValue("email"),
		"nick":     r.FormValue("nick"),
		"password": r.FormValue("password"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users", config.ApiUrl)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(user))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.FormatStatusCodeErrors(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

func LoadUsersPage(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))
	url := fmt.Sprintf("%s/users?user=%s", config.ApiUrl, nameOrNick)

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

	var users []models.User
	if err = json.NewDecoder(response.Body).Decode(&users); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Error: err.Error()})
	}

	utils.RenderTemplate(w, "users.html", users)
}

func StopFollowingUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/stop-following", config.ApiUrl, userID)
	response, err := requests.MakeARequestWithAuthentication(r, http.MethodPost, url, nil)

	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.FormatStatusCodeErrors(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/follow", config.ApiUrl, userID)
	response, err := requests.MakeARequestWithAuthentication(r, http.MethodPost, url, nil)

	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.FormatStatusCodeErrors(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

func EditUserProfile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, err := json.Marshal(map[string]string{
		"name":  r.FormValue("name"),
		"nick":  r.FormValue("nick"),
		"email": r.FormValue("email"),
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	cookie, _ := cookies.ReadCookie(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d/edit-user", config.ApiUrl, userID)

	response, err := requests.MakeARequestWithAuthentication(r, http.MethodPut, url, bytes.NewBuffer(user))

	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.FormatStatusCodeErrors(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	passwords, err := json.Marshal(map[string]string{
		"current": r.FormValue("current"),
		"new":     r.FormValue("new"),
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	cookie, _ := cookies.ReadCookie(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d/update-password", config.ApiUrl, userID)

	response, err := requests.MakeARequestWithAuthentication(r, http.MethodPost, url, bytes.NewBuffer(passwords))

	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.FormatStatusCodeErrors(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.ReadCookie(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d", config.ApiUrl, userID)

	response, err := requests.MakeARequestWithAuthentication(r, http.MethodDelete, url, nil)

	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.FormatStatusCodeErrors(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}
