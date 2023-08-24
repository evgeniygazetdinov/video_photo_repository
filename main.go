package main

import (
	"encoding/json"
	// "fmt"
	"log"
	"net/http"
	"os"
	// "strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)


func jsonResponse(forJsonCondition string, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"respose":forJsonCondition,"Status": "true", "Err": ""})
}


func get_variable_string_from_uri(variable_name string, w http.ResponseWriter, r *http.Request) string {
	variable := r.URL.Query().Get(variable_name)
	if len(variable) < 0 {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "missed"}`))
	}
	return variable
}

func upload(w http.ResponseWriter, r *http.Request) {
	jsonResponse("1", w, r)
}

func downloadVideo(writter http.ResponseWriter, request *http.Request){
	videoId := get_variable_string_from_uri("id", writter, request)
	jsonResponse(videoId, writter, request)
}


func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/upload/", upload)
	router.HandleFunc("/api/download_by_id/", downloadVideo)


	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	log.Println("listening on 8081")
	http.ListenAndServe(":8081", loggedRouter)
}
