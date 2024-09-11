package main

import (
	"gosrv/database"
	"gosrv/models"
	"gosrv/routes"
	"gosrv/static"

	"log"
	"net/http"
)

func main() {
	database.Connect()
	database.DATABASE.AutoMigrate(&models.Person{})

	router := routes.RegisterWebRoutes()

	log.Println("Server starting on :1337")

	fs := http.FileServer(http.FS(static.STATIC_FILES))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":1337", router))
}
