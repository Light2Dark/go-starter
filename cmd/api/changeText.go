package main

import (
	"net/http"

	"github.com/Light2Dark/go-starter/internal/components"
)

func (app application) changeTextHandler(w http.ResponseWriter, r *http.Request) {
	components.NewText().Render(r.Context(), w)
}
