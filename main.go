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



func succes_place(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"respose":"1","Status": "true", "Err": ""})
}


// func get_variable_string_from_uri(variable_name string, w http.ResponseWriter, r *http.Request) []string {
// 	variable, ok := r.URL.Query()[variable_name]
// 	if !ok || len(variable[0]) < 1 {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write([]byte(`{"message": "missed"}`))
// 	}
// 	return variable
// }

// func get_integers_for_equalize(one_name string, second_name string, w http.ResponseWriter, r *http.Request) (int, int) {

// 	one := get_variable_string_from_uri(one_name, w, r)
// 	two := get_variable_string_from_uri(second_name, w, r)
// 	first_int, err := strconv.Atoi(one[0])
// 	second_int, err := strconv.Atoi(two[0])
// 	if err != nil {
// 		fmt.Println((err))
// 	}
// 	return first_int, second_int
// }



func get(w http.ResponseWriter, r *http.Request) {
	// first, second := get_integers_for_equalize(one_predicate, two_predicate, w, r)
	// result := first + second
	succes_place(w, r)
}

func upload(w http.ResponseWriter, r *http.Request) {
	// first, second := get_integers_for_equalize(one_predicate, two_predicate, w, r)
	// result := first + second
	succes_place(w, r)
}


func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/upload/", upload)
	router.HandleFunc("/api/download/", download)
	router.HandleFunc("/api/stream_video/", download)
	router.HandleFunc("/api/video_list/", download)
	router.HandleFunc("/api/photo_list/", download)

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	log.Println("listening on 8081")
	http.ListenAndServe(":8081", loggedRouter)
}
