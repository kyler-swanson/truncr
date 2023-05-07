package db

import (
	"database/sql"

	"github.com/kyler-swanson/truncr/models"
	"github.com/kyler-swanson/truncr/utils"
)

func (db Database) AddLink(link *models.Link) error {
	var id int
	var createdAt string

	var shortLink = utils.GenerateShortURL()
	link.ShortLink = shortLink

	query := `INSERT INTO links (link, short_link) VALUES ($1, $2) RETURNING id, created_at`
	err := db.Conn.QueryRow(query, link.Link, link.ShortLink).Scan(&id, &createdAt)

	if err != nil {
		return err
	}

	link.ID = id
	link.CreatedAt = createdAt

	return nil
}

func (db Database) GetLinkById(linkId int) (models.Link, error) {
	link := models.Link{}

	query := `SELECT * FROM links WHERE id = $1;`
	row := db.Conn.QueryRow(query, linkId)

	switch err := row.Scan(&link.ID, &link.Link, &link.ShortLink, &link.CreatedAt); err {
	case sql.ErrNoRows:
		return link, ErrNoMatch
	default:
		return link, err
	}
}

func (db Database) GetLinkByShortLink(shortLink string) (models.Link, error) {
	link := models.Link{}

	query := `SELECT * FROM links WHERE short_link = $1;`
	row := db.Conn.QueryRow(query, shortLink)

	switch err := row.Scan(&link.ID, &link.Link, &link.ShortLink, &link.CreatedAt); err {
	case sql.ErrNoRows:
		return link, ErrNoMatch
	default:
		return link, err
	}
}
