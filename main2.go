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

	"context"
    "fmt"
    "time"
 
    "video_photo_repository/configs" //add this

)



func close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc){
	 
	// CancelFunc to cancel to context
	defer cancel()

	// client provides a method to close
	// a mongoDB connection.
	defer func(){

	// client.Disconnect method also has deadline.
	// returns error if any,
	if err := client.Disconnect(ctx); err != nil{
		panic(err)
	}
	}()
}

// This is a user defined method that returns mongo.Client,
// context.Context, context.CancelFunc and error.
// mongo.Client will be used for further database operation.
// context.Context will be used set deadlines for process.
// context.CancelFunc will be used to cancel context and
// resource associated with it.

func connect(uri string)(*mongo.Client, context.Context,
				   context.CancelFunc, error) {
						
	// ctx will be used to set deadline for process, here
	// deadline will of 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(),
									30 * time.Second)

	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

// This is a user defined method that accepts
// mongo.Client and context.Context
// This method used to ping the mongoDB, return error if any.
func ping(client *mongo.Client, ctx context.Context) error{

	// mongo.Client has Ping to ping mongoDB, deadline of
	// the Ping method will be determined by cxt
	// Ping method return error if any occurred, then
	// the error can be handled.
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		fmt.Println("eroor")
		return err
	}
	fmt.Println("connected successfully")
	return nil
}


// %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%


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

func uploadVideo(writter http.ResponseWriter, request *http.Request){
	videoId := get_variable_string_from_uri("id", writter, request)
	jsonResponse(videoId, writter, request)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/upload/", upload)
	router.HandleFunc("/api/download_by_id/", downloadVideo)

	// client, ctx, cancel, err := connect("mongodb://localhost:27017")
    // if err != nil{
    //     panic(err)
    // }
     
    // // Release resource when the main
    // // function is returned.
    // defer close(client, ctx, cancel)
     
    // // Ping mongoDB with Ping method
    // ping(client, ctx)
	configs.ConnectDB()

    // quickstartDatabase := client.Database("video_repository")
	// videoCollection := quickstartDatabase.Collection("video")
    // photoCollection := quickstartDatabase.Collection("photo")
	// videoCollection, err := podcastsCollection.InsertOne(ctx, bson.D{
	// 	{"title", "The Polyglot Developer Podcast"},
	// 	{"author", "Nic Raboy"},
	// 	{"tags", bson.A{"development", "programming", "coding"}},
	// })


	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	log.Println("listening on 8081")
	http.ListenAndServe(":8081", loggedRouter)
}
