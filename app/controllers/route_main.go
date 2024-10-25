package controllers

import (
	"net/http"
)

// ハンドラの役割(=HttpServe)
func top(w http.ResponseWriter, r *http.Request) {
	// Cookie認証がなければ表示
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, "Hello!", "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/todos", http.StatusFound)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	// Cookie認証があれば表示
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "index")
	}
}
