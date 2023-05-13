package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kyler-swanson/truncr/db"
	"github.com/kyler-swanson/truncr/handler/api"
)

var dbInstance db.Database

func CreateHandler(db db.Database) http.Handler {
	router := chi.NewRouter()
	dbInstance = db

	router.Get("/", renderApp)

	router.Route("/{shortLink}", func(router chi.Router) {
		router.Get("/", redirectLink)
	})

	api.Register(router, db)

	return router
}
