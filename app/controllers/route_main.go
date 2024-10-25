package controllers

import (
	"net/http"
)

// ハンドラの役割(=HttpServe)
func top(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "Hello!", "layout", "public_navbar", "top")
}
