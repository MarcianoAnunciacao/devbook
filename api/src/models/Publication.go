package models

import (
	"errors"
	"strings"
	"time"
)

type Publication struct {
	ID             uint64    `json:"id,omitempty"`
	Title          string    `json:"title,omitempty"`
	Content        string    `json:"content,omitempty"`
	AuthorID       uint64    `json:"authorId,omitempty"`
	AuthorNickName string    `json:"authorNickName,omitempty"`
	Likes          uint64    `json:"likes"`
	CreatedAt      time.Time `json:"createdAt,omitempty"`
}

func (publication *Publication) ValidateInputData() error {
	if err := publication.validateFields(); err != nil {
		return err
	}

	publication.formatFields()
	return nil
}

func (publication *Publication) validateFields() error {
	if publication.Title == "" {
		return errors.New("Please provide a Publication Title")
	}

	if publication.Content == "" {
		return errors.New("Please provide a Publication Content")
	}

	return nil
}

func (publication *Publication) formatFields() {
	publication.Title = strings.TrimSpace(publication.Title)
	publication.Content = strings.TrimSpace(publication.Content)
}
