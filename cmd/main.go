package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Ibukun-tech/try"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Remeber to set up a a var that will store something once the init function runs

// Here I am going to run some init function first to do somethings
// like to send something to the database
var handle http.Handler

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer func() {
		cancel()
	}()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://ibk:secret@localhost:27017"))
	if err != nil {
		os.Exit(1)
	}
	conn := try.NewMongoClient(client)
	s := try.NewServer(conn)
	m := mux.NewRouter()
	m.HandleFunc("/", try.RunHandler(s.RegisterHandler))
	handle = m
}

// The server is invoked and started
func main() {

	serve := &http.Server{
		Addr:    ":4000",
		Handler: handle,
	}
	log.Fatal(serve.ListenAndServe())
}
