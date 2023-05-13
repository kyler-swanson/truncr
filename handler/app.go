package handler

import (
	"net/http"

	"github.com/go-chi/render"
)

func renderApp(w http.ResponseWriter, r *http.Request) {
	render.HTML(w, r, "app")
}
