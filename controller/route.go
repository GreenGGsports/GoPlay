package controller

import (
	"net/http"
)

func Register() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", ping())
	mux.HandleFunc("/joke", joke())
	mux.HandleFunc("/add", create())
	return mux
}
