package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	movies = append(movies,
		Movie{
			ID:    "1",
			Isbn:  "438227",
			Title: "Movie 1",
			Director: &Director{
				Firstname: "John",
				Lastname:  "Doe",
			},
		})

	movies = append(movies,
		Movie{
			ID:    "2",
			Isbn:  "45455",
			Title: "Movie 2",
			Director: &Director{
				Firstname: "Stieve",
				Lastname:  "Smith",
			},
		})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("movies", createMovie).Methods("POST")
	r.HandleFunc("movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("starting server at port 8000 \n")

	log.Fatal(http.ListenAndServe("8000", r))
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for index, item := range movies {

		if item.ID == param["id"] {
			movies = append(movies[:index])
		}
	}
}
