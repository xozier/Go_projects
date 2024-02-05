package main

import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)


type Movie struct {
	ID string `json:"id"`
	ISBN string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname`
	Surname string `json:"surname"`
}



var movies []Movie



func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for indx, m := range movies{
		if m.ID == params["id"]{
            movies= append(movies[:indx], movies[indx+1:]...)
			break
        }
	}
	json.NewEncoder(w).Encode(movies)
}


func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, m := range movies{
		if m.ID == params["id"]{
			json.NewEncoder(w).Encode(m)
			return
		}
	}
}


func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie Movie 
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa((rand.Intn(1000000)))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}


func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for indx, m := range movies{
		if m.ID == params["id"]{
			movies= append(movies[:indx], movies[indx+1:]...)
			break
		}
	}
	var movie Movie 
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa((rand.Intn(1000000)))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
	return
}

func main() {
	movies = append(movies, Movie{
		ID: "1",
	    ISBN: "1",
	    Title: "movie 1",
	    Director: &Director{
        FirstName: "john",
        Surname: "doe",
        },
    })
	movies = append(movies, Movie{
		ID: "2",
        ISBN: "2",
        Title: "movie 2",
        Director: &Director{
        FirstName: "jane",
        Surname: "doe",
        },
    })
    
	r := mux.NewRouter()

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("starting server on port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", nil))
}