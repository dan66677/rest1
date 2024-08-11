package main

import (
	"log"
	"net/http"

	"restapi/film/internal/app"
	"restapi/film/internal/db"
)

func main() {
	db.Init()

	http.HandleFunc("/actors", app.AddActorHandler)
	http.HandleFunc("/actors/update", app.UpdateActorHandler)
	http.HandleFunc("/actors/delete", app.DeleteActorHandler)
	http.HandleFunc("/movies", app.AddMovieHandler)
	http.HandleFunc("/movies/list", app.GetMoviesHandler)
	http.HandleFunc("/actors/list", app.GetActorsHandler)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
