package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
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

//	func authMiddlewareHandler(next http.Handler) http.HandlerFunc {
//		return func(w http.ResponseWriter, r *http.Request) {
//			w.Header().Set("authorizaton", "Bearer token")
//			next.ServeHTTP(w, r)
//		}
//	}

// The server is invoked and started
func main() {
	fmt.Println("first run")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer func() {
		cancel()
	}()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://ibk:secret@localhost:27017"))
	fmt.Println(err, "why the error")
	if err != nil {
		fmt.Println(err, "if there is error ")
		log.Fatal(err)
	}
	conn := try.NewMongoClient(client)
	s := try.NewServer(conn)
	m := mux.NewRouter()
	m.HandleFunc("/add", try.RunHandler(s.RegisterHandler))
	m.HandleFunc("/getAll", try.RunHandler(s.GetAllHandler))
	handle = m
	serve := &http.Server{
		Addr:    ":4000",
		Handler: handle,
	}
	log.Fatal(serve.ListenAndServe())
}
