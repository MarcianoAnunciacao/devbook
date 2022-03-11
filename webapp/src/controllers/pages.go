package controllers

import (
	"net/http"
	"webapp/src/utils"
)

func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "login.html", nil)
}

func LoadCreateUserPage(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "userregistration.html", nil)
}

func LoadMainPage(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "home.html", nil)
}
