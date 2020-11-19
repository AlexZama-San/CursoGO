package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Route struct {
	Name string
	Method string
	Pattern string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Name(route.Name).Methods(route.Method).Path(route.Pattern).Handler(route.HandleFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"Get",
		"/",
		Index,
	},
	Route{
		"peliculaList",
		"Get",
		"/peliculas",
		PeliculaList,
	},
	Route{
		"MOVIESHOW",
		"Get",
		"/peliculas{id}",
		MovieShow,
	},
	Route{
		"MovieAdd",
		"Post",
		"/pelicula",
		MovieAdd,
	},
	Route{
		"MovieUpdate",
		"Put",
		"/pelicula/{id}",
		MovieUpdate,
	},
	Route{
		"MovieRemove",
		"Delete",
		"/peliculaQ/{id}",
		MovieRemove,
	},
}