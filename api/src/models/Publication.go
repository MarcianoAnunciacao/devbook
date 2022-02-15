package models

import "time"

type Publication struct {
	ID             uint64    `json:"id,omitempty"`
	Title          string    `json:"title,omitempty"`
	Content        string    `json:"content,omitempty"`
	AuthorID       uint64    `json:"authorId,omitempty"`
	AuthorNickName string    `json:"authorNickName,omitempty"`
	Likes          uint64    `json:"likes"`
	CreatedAt      time.Time `json:"createdAt,omitempty"`
}
