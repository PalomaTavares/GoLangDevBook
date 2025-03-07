package models

import "time"

type Post struct {
	ID         uint64    `json:"id,omempty"`
	Title      string    `json:"title,omempty"`
	Content    string    `json:"content,omempty"`
	IDAuthor   uint64    `json:"idAuthor,omempty"`
	NickAuthor string    `json:"nickAuthor,omempty"`
	Likes      uint64    `json:"likes"`
	CreatedIn  time.Time `"json:createdIn,omempty`
}
