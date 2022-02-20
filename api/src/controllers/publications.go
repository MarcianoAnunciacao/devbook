package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreatePublication(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserIDFromToken(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publication models.Publication
	if err = json.Unmarshal(requestBody, &publication); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	publication.AuthorID = userID

	if err = publication.ValidateInputData(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.PublicationRepository(db)
	publication.ID, err = repository.Create(publication)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, publication)
}
func SearchPublications(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserIDFromToken(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.PublicationRepository(db)
	publications, err := repository.SearchPublications(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, publications)
}
func SearchPublication(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	publicationID, err := strconv.ParseUint(parameters["publicationID"], 10, 64)
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

	repository := repositories.PublicationRepository(db)
	publication, err := repository.SearchByID(publicationID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, publication)
}
func UpdatePublication(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserIDFromToken(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	parameters := mux.Vars(r)

	publicationID, err := strconv.ParseUint(parameters["publicationID"], 10, 64)
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

	repository := repositories.PublicationRepository(db)
	savedPublication, err := repository.SearchByID(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if savedPublication.AuthorID != userID {
		responses.Error(w, http.StatusForbidden, errors.New("You cannot update another Author publication"))
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publication models.Publication
	if err = json.Unmarshal(requestBody, &publication); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = publication.ValidateInputData(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = repository.Update(publicationID, publication); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

func DeletePublication(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserIDFromToken(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	parameters := mux.Vars(r)

	publicationID, err := strconv.ParseUint(parameters["publicationID"], 10, 64)
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

	repository := repositories.PublicationRepository(db)
	savedPublication, err := repository.SearchByID(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if savedPublication.AuthorID != userID {
		responses.Error(w, http.StatusForbidden, errors.New("You cannot delete another Author publication"))
		return
	}

	if err = repository.Delete(publicationID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func SearchPublicationsByUserID(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, err := strconv.ParseUint(parameters["userID"], 10, 64)
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

	repository := repositories.PublicationRepository(db)
	publications, err := repository.SearchPublicationsByUserID(userID)

	responses.JSON(w, http.StatusOK, publications)
}

func LikeIt(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	publicationID, err := strconv.ParseUint(parameters["id"], 10, 64)
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

	repository := repositories.PublicationRepository(db)
	if err = repository.LikeIt(publicationID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func DisLikeIt(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	publicationID, err := strconv.ParseUint(parameters["id"], 10, 64)
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

	repository := repositories.PublicationRepository(db)
	if err = repository.DisLikeIt(publicationID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
