package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func HandleUpdating(w http.ResponseWriter, r *http.Request) {
	path := chi.URLParam(r, "path")
	render.PlainText(w, r, "Hello from PUT" + path)
}
