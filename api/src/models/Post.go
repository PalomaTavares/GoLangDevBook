package models

import (
	"errors"
	"strings"
	"time"
)

type Post struct {
	ID         uint64    `json:"id,omempty"`
	Title      string    `json:"title,omempty"`
	Content    string    `json:"content,omempty"`
	IDAuthor   uint64    `json:"idAuthor,omempty"`
	NickAuthor string    `json:"nickAuthor,omempty"`
	Likes      uint64    `json:"likes"`
	CreatedIn  time.Time `"json:createdIn,omempty`
}

func (post *Post) Prepare() error {
	if error := post.validate(); error != nil {
		return error
	}

	post.format()
	return nil
}

func (post *Post) validate() error {
	if post.Title == "" {
		return errors.New("title is required")
	}
	if post.Content == "" {
		return errors.New("content is required")
	}
	return nil
}

func (post *Post) format() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}
