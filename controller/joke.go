package controller

import (
	"GoPlay/views"
	"encoding/json"
	"net/http"
)

func joke() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := views.Response{
				Code: http.StatusOK,
				Body: "egy két há kurva anyád",
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		}
	}
}
