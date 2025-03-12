package controller

import (
	"GoPlay/views"
	"encoding/json"
	"net/http"
)

func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := views.Response{
				Code: http.StatusOK,
				Body: "Pong",
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		}
	}

}
