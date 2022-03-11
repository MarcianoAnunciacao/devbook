package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responses"
	"webapp/src/utils"
)

func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println(response.StatusCode, err)

	utils.RenderTemplate(w, "home.html", publications)
}
