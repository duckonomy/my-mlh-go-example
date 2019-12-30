package handlers

import (
	"net/http"
)

func New() http.Handler {
	mux := http.NewServeMux()

	// Root
	mux.Handle("/",  http.FileServer(http.Dir("web/")))

	// OauthGoogle
	mux.HandleFunc("/auth/mymlh/login", oauthLogin)
	mux.HandleFunc("/auth/mymlh/callback", oauthCallback)

	return mux
}
