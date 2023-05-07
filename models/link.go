package models

import (
	"fmt"
	"net/http"
)

type Link struct {
	ID        int    `json:"id"`
	Link      string `json:"link"`
	ShortLink string `json:"short_link"`
	CreatedAt string `json:"created_at"`
}

func (i *Link) Bind(r *http.Request) error {
	if i.Link == "" {
		return fmt.Errorf("link is a required field")
	}
	return nil
}

func (*Link) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
