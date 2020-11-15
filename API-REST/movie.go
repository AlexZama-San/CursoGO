package main

type Movie struct {
	Name string `json:"nombre"`
	Year int	`json:"año"`
	Director string `json:"director"`
}

type Movies []Movie