package main

import (
    "log"
    "video_photo_repository/configs" //add this
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()

    //run database
    configs.ConnectDB()

    log.Fatal(http.ListenAndServe(":6000", router))
}