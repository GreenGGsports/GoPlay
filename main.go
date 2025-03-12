package main

import (
	"GoPlay/controller"
	"GoPlay/models"
	"log"
	"net/http"
)

func main() {
	mux := controller.Register()
	db, err := models.Connect()
	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	models.CreateTodo()
	defer db.Close()
	http.ListenAndServe("localhost:3000", mux)
}
