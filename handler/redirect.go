package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/kyler-swanson/truncr/db"
	"github.com/kyler-swanson/truncr/handler/api"
)

func redirectLink(w http.ResponseWriter, r *http.Request) {
	shortLink := chi.URLParam(r, "shortLink")

	link, err := dbInstance.GetLinkByShortLink(shortLink)

	if err != nil {
		if err == db.ErrNoMatch {
			renderApp(w, r)
		} else {
			render.Render(w, r, api.ErrorRenderer(err))
		}
		return
	}

	http.Redirect(w, r, link.Link, http.StatusMovedPermanently)
}
