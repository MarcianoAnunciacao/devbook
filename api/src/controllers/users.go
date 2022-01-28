package controllers

import "net/http"

//Create an user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user"))
}

//Search all users
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user"))
}

//Search an user from his ID
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user"))
}

//Update user with new informations
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user"))
}

//Delete a user from the system
func RemoveUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user"))
}
