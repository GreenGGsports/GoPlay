package controller

import (
	"GoPlay/models"
	"GoPlay/views"
	"encoding/json"
	"fmt"
	"net/http"
)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data := views.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			fmt.Println(data)
			if err := models.CreateTodo(); err != nil {
				w.Write([]byte("Some error"))
				return
			}
			fmt.Println("entry added succesfully!")

		} else if r.Method == http.MethodGet {
			name := r.URL.Query().Get("name")
			data, err := models.ReadbyName(name)
			if err != nil {
				w.Write([]byte(err.Error()))
			}
			fmt.Println(data)
			json.NewEncoder(w).Encode(data)

		}
	}
}
