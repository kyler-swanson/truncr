package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/kyler-swanson/truncr/db"
	"github.com/kyler-swanson/truncr/models"
)

var linkIdKey = "linkId"
var shortLinkKey = "shortLink"

func link(router chi.Router) {
	router.Post("/", createLink)
	router.Route("/{linkId}", func(router chi.Router) {
		router.Use(LinkContext)
		router.Get("/", getLink)
	})
}

func LinkContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		linkId := chi.URLParam(r, "linkId")

		if linkId == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("link ID is required")))
			return
		}

		id, err := strconv.Atoi(linkId)
		if err != nil {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid link ID")))
			return
		}

		ctx := context.WithValue(r.Context(), linkIdKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ShortLinkContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		shortLink := chi.URLParam(r, "shortLink")

		if shortLink == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("short link is required")))
			return
		}

		ctx := context.WithValue(r.Context(), shortLinkKey, shortLink)
		next.ServeHTTP(w, r.WithContext(ctx))
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
	linkId := r.Context().Value(linkIdKey).(int)

	link, err := dbInstance.GetLinkById(linkId)

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

func redirectLink(w http.ResponseWriter, r *http.Request) {
	shortLink := r.Context().Value(shortLinkKey).(string)

	link, err := dbInstance.GetLinkByShortLink(shortLink)

	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}

	http.Redirect(w, r, link.Link, http.StatusMovedPermanently)
}
