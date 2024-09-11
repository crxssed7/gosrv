package main

import (
	"gosrv/database"
	"gosrv/models"
	"gosrv/routes"

	"log"
	"net/http"
)

func main() {
	database.Connect()
	database.DATABASE.AutoMigrate(&models.Person{})

	router := routes.RegisterWebRoutes()

	log.Println("Server starting on :1337")

	fs := http.FileServer(http.Dir("./static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":1337", router))
}
