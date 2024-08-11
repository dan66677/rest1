package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"restapi/film/internal/db"
)

func AddActorHandler(w http.ResponseWriter, r *http.Request) {
	var actor Actor
	err := json.NewDecoder(r.Body).Decode(&actor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var id int
	err = db.DB.QueryRow("INSERT INTO actors (name, gender, birth_date) VALUES ($1, $2, $3) RETURNING id", actor.Name, actor.Gender, actor.BirthDate).Scan(&id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	actor.ID = id
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(actor)
}

// Аналогично добавьте обработчики для других запросов
func UpdateActorHandler(w http.ResponseWriter, r *http.Request) {
	var actor Actor
	err := json.NewDecoder(r.Body).Decode(&actor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.DB.Exec("UPDATE actors SET name=$1, gender=$2, birth_date=$3 WHERE id=$4", actor.Name, actor.Gender, actor.BirthDate, actor.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(actor)
}

func DeleteActorHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, err = db.DB.Exec("DELETE FROM actors WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func AddMovieHandler(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var id int
	err = db.DB.QueryRow("INSERT INTO movies (title, description, release_date, rating) VALUES ($1, $2, $3, $4) RETURNING id", movie.Title, movie.Description, movie.ReleaseDate, movie.Rating).Scan(&id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	movie.ID = id
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(movie)
}

func GetMoviesHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, title, description, release_date, rating FROM movies ORDER BY rating DESC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var movies []Movie
	for rows.Next() {
		var movie Movie
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.Rating)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		movies = append(movies, movie)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movies)
}

func GetActorsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, name, gender, birth_date FROM actors")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var actors []Actor
	for rows.Next() {
		var actor Actor
		err := rows.Scan(&actor.ID, &actor.Name, &actor.Gender, &actor.BirthDate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		actors = append(actors, actor)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(actors)
}
