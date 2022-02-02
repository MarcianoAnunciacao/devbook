package models

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

//Represents an user on the social network devbook
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	NickName  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (user *User) ValidateInputData(step string) error {
	if err := user.validateFields(step); err != nil {
		return err
	}

	user.formatFields()
	return nil
}

func (user *User) validateFields(step string) error {
	if user.Name == "" {
		return errors.New("Please provide an User Name")
	}

	if user.NickName == "" {
		return errors.New("Please provide an User Nick Name")
	}

	if user.Email == "" {
		return errors.New("Please provide an User valid E-mail")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("Please provide an User valid E-mail, ex: youremail@youremailprovider.com")
	}

	if step == "create" && user.Password == "" {
		return errors.New("Please provide an User password")
	}

	return nil
}

func (user *User) formatFields() {
	user.Name = strings.TrimSpace(user.Name)
	user.NickName = strings.TrimSpace(user.NickName)
	user.Email = strings.TrimSpace(user.Email)
}
