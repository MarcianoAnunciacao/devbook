package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/requests"
)

type User struct {
	ID           uint64        `json:"id"`
	Name         string        `json:"name"`
	Email        string        `json:"email"`
	Nick         string        `json:"nick"`
	CreatedAt    string        `json:"createdAt"`
	Followers    []User        `json:"followers"`
	Following    []User        `json:"following"`
	Publications []Publication `json:"publications"`
}

func SearchUserWithAllInformations(userID uint64, r *http.Request) (User, error) {
	userChannel := make(chan User)
	followersChannel := make(chan []User)
	followingChannel := make(chan []User)
	publicationsChannel := make(chan []Publication)

	go SearchUserData(userChannel, userID, r)
	go SearchFollowers(followersChannel, userID, r)
	go SearchFollowing(followingChannel, userID, r)
	go SearchPublications(publicationsChannel, userID, r)

	var (
		user         User
		followers    []User
		following    []User
		publications []Publication
	)

	for i := 0; i < 4; i++ {
		select {
		case loadedUser := <-userChannel:
			if loadedUser.ID == 0 {
				return User{}, errors.New("Error searching for User")
			}

			user = loadedUser
		case loadedFollowers := <-followersChannel:
			if loadedFollowers == nil {
				return User{}, errors.New("Error searching for Followers")
			}

			followers = loadedFollowers
		case loadedFollowing := <-followingChannel:
			if loadedFollowing == nil {
				return User{}, errors.New("Error searching for Users following")
			}

			following = loadedFollowing
		case loadedPublications := <-publicationsChannel:
			if loadedPublications == nil {
				return User{}, errors.New("Error searching for Users publications")
			}

			publications = loadedPublications
		}
	}

	user.Followers = followers
	user.Following = following
	user.Publications = publications

	return user, nil
}

func SearchUserData(channel chan<- User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d", config.ApiUrl, userID)
	response, err := requests.MakeARequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- User{}
		return
	}
	defer response.Body.Close()

	var user User
	if err = json.NewDecoder(response.Body).Decode(&user); err != nil {
		channel <- User{}
		return
	}

	channel <- user
}

func SearchFollowers(channel chan<- []User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/followers", config.ApiUrl, userID)
	response, err := requests.MakeARequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var followers []User
	if err = json.NewDecoder(response.Body).Decode(&followers); err != nil {
		channel <- nil
		return
	}

	if followers == nil {
		channel <- make([]User, 0)
		return
	}

	channel <- followers
}

func SearchFollowing(channel chan<- []User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/following", config.ApiUrl, userID)
	response, err := requests.MakeARequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var following []User
	if err = json.NewDecoder(response.Body).Decode(&following); err != nil {
		channel <- nil
		return
	}

	if following == nil {
		channel <- make([]User, 0)
		return
	}

	channel <- following
}

func SearchPublications(channel chan<- []Publication, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/publications", config.ApiUrl, userID)
	response, err := requests.MakeARequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var publications []Publication
	if err = json.NewDecoder(response.Body).Decode(&publications); err != nil {
		channel <- nil
		return
	}

	if publications == nil {
		channel <- make([]Publication, 0)
		return
	}

	channel <- publications
}
