package main

import (
	"net/http"

	"github.com/Light2Dark/go-starter/internal/templates"
)

func (app application) homeHandler(w http.ResponseWriter, r *http.Request) {
	templates.Home().Render(r.Context(), w)
}
