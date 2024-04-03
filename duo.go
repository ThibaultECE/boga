package main

import (
	"net/http"
	"encoding/json"
	// "errors"
	"fmt"
)

type ApiResponse struct {
	Message string `json:"message"`
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	}

func dataHandler(w http.ResponseWriter, r *http.Request) {

	// err := errors.New("an error occured")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	bouteille := ApiResponse{Message: "Voici le vestibule"}

	js,err := json.Marshal(bouteille)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func main() {

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/",fs)

	api := http.NewServeMux()
	api.HandleFunc("/api/data", dataHandler)
	http.Handle("/api/",api)

	fmt.Printf("Le back se lance")
	http.ListenAndServe("localhost:8080",nil)
}