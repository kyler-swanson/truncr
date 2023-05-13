package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/kyler-swanson/truncr/db"
	"github.com/kyler-swanson/truncr/models"
)

var linkIdKey = "linkId"

func link(router chi.Router) {
	router.Post("/", createLink)
	router.Route("/{shortLink}", func(router chi.Router) {
		router.Get("/", getLink)
	})
}

func createLink(w http.ResponseWriter, r *http.Request) {
	link := &models.Link{}

	if err := render.Bind(r, link); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	if err := dbInstance.AddLink(link); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, link); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func getLink(w http.ResponseWriter, r *http.Request) {
	shortLink := chi.URLParam(r, "shortLink")

	link, err := dbInstance.GetLinkByShortLink(shortLink)

	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}

	if err := render.Render(w, r, &link); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
