package app

import (
	"restapi/film/internal/db"
)

func GetActorByID(id int) (*Actor, error) {
	var actor Actor
	err := db.DB.QueryRow("SELECT id, name, gender, birth_date FROM actors WHERE id = $1", id).Scan(&actor.ID, &actor.Name, &actor.Gender, &actor.BirthDate)
	if err != nil {
		return nil, err
	}
	return &actor, nil
}

// Аналогично добавьте другие сервисы

func GetMovieByID(id int) (*Movie, error) {
	var movie Movie
	err := db.DB.QueryRow("SELECT id, title, description, release_date, rating FROM movies WHERE id = $1", id).Scan(&movie.ID, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.Rating)
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

func GetMoviesByActor(actorID int) ([]Movie, error) {
	rows, err := db.DB.Query("SELECT m.id, m.title, m.description, m.release_date, m.rating FROM movies m JOIN movie_actors ma ON m.id = ma.movie_id WHERE ma.actor_id = $1", actorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []Movie
	for rows.Next() {
		var movie Movie
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.Rating)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func GetActorsByMovie(movieID int) ([]Actor, error) {
	rows, err := db.DB.Query("SELECT a.id, a.name, a.gender, a.birth_date FROM actors a JOIN movie_actors ma ON a.id = ma.actor_id WHERE ma.movie_id = $1", movieID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var actors []Actor
	for rows.Next() {
		var actor Actor
		err := rows.Scan(&actor.ID, &actor.Name, &actor.Gender, &actor.BirthDate)
		if err != nil {
			return nil, err
		}
		actors = append(actors, actor)
	}
	return actors, nil
}

func SearchMoviesByTitle(titleFragment string) ([]Movie, error) {
	rows, err := db.DB.Query("SELECT id, title, description, release_date, rating FROM movies WHERE title ILIKE $1", "%"+titleFragment+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []Movie
	for rows.Next() {
		var movie Movie
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.Rating)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func SearchMoviesByActorName(actorNameFragment string) ([]Movie, error) {
	rows, err := db.DB.Query(`
        SELECT m.id, m.title, m.description, m.release_date, m.rating
        FROM movies m
        JOIN movie_actors ma ON m.id = ma.movie_id
        JOIN actors a ON ma.actor_id = a.id
        WHERE a.name ILIKE $1`, "%"+actorNameFragment+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []Movie
	for rows.Next() {
		var movie Movie
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.Rating)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}
