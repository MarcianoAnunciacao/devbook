package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/models"
	"webapp/src/responses"
)

func SignupWithEmailAndPasswordToGetToken(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("email"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/login", config.ApiUrl)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(user))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.FormatStatusCodeErrors(w, response)
	}

	var authData models.AuthData
	if err = json.NewDecoder(response.Body).Decode(&authData); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Error: err.Error()})
		return
	}

}
