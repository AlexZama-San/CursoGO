package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"log"
)
var movies = Movies{
	Movie{"Transformers 5", 2018, "Michael Bay"},
	Movie{"Batman Inicia", 1999, "No se xD"},
	Movie{"HellBoy", 2005, "Del toro"},
}

func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hola cosa GO")
}
func PeliculaList(w http.ResponseWriter, r *http.Request){

	//fmt.Fprintf(w, "listado de peliculas")

	json.NewEncoder(w).Encode(movies)
}
func MovieShow(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	movie_ID := params["id"]

	fmt.Fprintf(w, "has cargado la pelicula numero %s", movie_ID)
}
func MovieAdd(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data)
	if(err != nil){
		panic(err)
	}
	defer r.Body.Close()
	log.Println(movie_data)
	movies = append(movies,movie_data)
}