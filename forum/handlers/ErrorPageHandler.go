package handlers

import (
	"net/http"

	"01.alem.school/git/atastemi/forum/forum/internal"
)

func (hh *HttpHandler) Handle404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	internal.RenderErrorPage(w, r)
}
