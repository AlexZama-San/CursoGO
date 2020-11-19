package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return session
}
var collection = getSession().DB("DBAPI").C("cosas")

func responseone(w http.ResponseWriter, status int, results Movie){
	w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(results)
}
func responseAll(w http.ResponseWriter, status int, results []Movie){
	w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(results)
}

func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hola cosa GO")
}
func PeliculaList(w http.ResponseWriter, r *http.Request){

	//fmt.Fprintf(w, "listado de peliculas")
	var result []Movie
	err := collection.Find(nil).Sort("-_id").All(&result)
	if err != nil{
		log.Fatal(err)
	}else{
		fmt.Println("Resultados:",result)
	}
	responseAll(w, 200, result)
}
func MovieShow(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	movie_ID := params["id"]

	if !bson.IsObjectIdHex(movie_ID){
		
		w.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(movie_ID)

	results := Movie{}
	err := collection.FindId(oid).One(&results)
	if err != nil{
		w.WriteHeader(404)
		return
	}else{
		responseone(w, 200, results)
	}


}
func MovieAdd(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data)
	if(err != nil){
		panic(err)
	}
	defer r.Body.Close()
	
	err = collection.Insert(movie_data)

	if err != nil {
		w.WriteHeader(500)
		return	
	}
	responseone(w, 200, movie_data)
}

func MovieUpdate(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	movie_ID := params["id"]

	if !bson.IsObjectIdHex(movie_ID){
		
		w.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(movie_ID)
	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data)

	if err != nil{
		panic(err)
		w.WriteHeader(404)
		return
	}

	defer r.Body.Close()

	document :=bson.M{"_id": oid}
	change := bson.M{"$set":movie_data}
	err = collection.Update(document,change)
	if err != nil{
		panic(err)
		w.WriteHeader(404)
		return
	}
	responseone(w, 200, movie_data)


}

type Message struct {
	Status string "json.status"
	Message  string "json.message"}	

func (this *Message) setStatus(data string){
	this.Status=data
}
func (this *Message) setMessage(data string){
	this.Message=data
}

func MovieRemove(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	movie_ID := params["id"]

	if !bson.IsObjectIdHex(movie_ID){
		
		w.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(movie_ID)

	
	err := collection.RemoveId(oid)
	if err != nil{
		w.WriteHeader(404)
		return
	}

	message := new(Message)
	message.setStatus("success")
	message.setMessage("la pelicula con ID"+movie_ID+"a sido borrada")

	results := message
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)


}