package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/kyler-swanson/truncr/db"
)

var dbInstance db.Database

func Register(router chi.Router, db db.Database) {
	dbInstance = db

	router.MethodNotAllowed(methodNotAllowedHandler)
	router.NotFound(notFoundHandler)

	router.Route("/api", func(router chi.Router) {
		router.Route("/link", link)
	})
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(405)
	render.Render(w, r, ErrMethodNotAllowed)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(400)
	render.Render(w, r, ErrNotFound)
}
