package main

import "net/http"

func (m *Respository) index(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home.html.template")
}
