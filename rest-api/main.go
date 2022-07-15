package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Music struct {
	Id      string `json:"Id"`
	Artist  string `json:"Artist"`
	Album   string `json:"Album"`
	Title   string `json:"Title"`
	Listens int    `json:"Listens"`
}

var Songs []Music

func deleteSong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for i, music := range Songs {
		if music.Id == id {
			Songs = append(Songs[:i], Songs[i+1:]...)
		}
	}
}

func createSong(w http.ResponseWriter, r *http.Request) {
	getBody, _ := ioutil.ReadAll(r.Body)
	var music Music
	json.Unmarshal(getBody, &music)
	Songs = append(Songs, music)

	json.NewEncoder(w).Encode(music)
}

func allSongs(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Songs)
}

func singleArtist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["artist"]

	for _, song := range Songs {
		if song.Id == key {
			json.NewEncoder(w).Encode(song)
		}
	}
}

func singleSong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, song := range Songs {
		if song.Artist == key {
			json.NewEncoder(w).Encode(song)
		}
	}
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the main page")
}

func requestHandler() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", mainPage)
	router.HandleFunc("/music", allSongs).Methods("GET")
	router.HandleFunc("/music/{artist}", singleArtist).Methods("GET")
	router.HandleFunc("/music/{id}", singleSong).Methods("GET")
	router.HandleFunc("/music", createSong).Methods("POST")
	router.HandleFunc("/music/{id}", deleteSong).Methods("DELETE")

	fmt.Println("Started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	Songs = []Music{
		{Id: "1", Artist: "Juice-World", Album: "Goodbye And Good Riddance", Title: "Lucid Dreams", Listens: 1993436185},
		{Id: "2", Artist: "Juice-World", Album: "Legends Never Die", Title: "Wishing Well", Listens: 663193859},
		{Id: "3", Artist: "Juice-World", Album: "Death Race For Love", Title: "Robbery", Listens: 1006047658},
		{Id: "4", Artist: "XXXTENTACION", Album: "17", Title: "Revenge", Listens: 637256256},
		{Id: "5", Artist: "XXXTENTACION", Album: "?", Title: "Sad", Listens: 1858784818},
		{Id: "6", Artist: "XXXTENTACION", Album: "LOOK AT ME: THE ALBUM", Title: "I spoke to the devil in miami, he said everything would be fine", Listens: 12831707},
	}
	requestHandler()
}
