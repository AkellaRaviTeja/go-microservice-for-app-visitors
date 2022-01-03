package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"visitor-registry/handlers"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx context.Context 
var cancel context.CancelFunc
var visitorCollection *mongo.Collection 
var mongoClient *mongo.Client

// Service execution starts here
// Establishes connection to mongodb
// Starts a server which listens on 9191 port
func main() {

	//Create a context
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	// Logger for the standard os output
	logger := log.New(os.Stdout,"visitor-registry",log.LstdFlags)
	
	mongoClient, connectionErr := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:password@localhost:27017"))
	if connectionErr != nil {
		logger.Printf("Unable to connect to the Database %v\n",connectionErr)
		os.Exit(1)
	}
	visitorCollection := mongoClient.Database("Apxor").Collection("visitor")

	// Create a router for the visitor endpoints
	router := mux.NewRouter()
	//Injecting the Logger, context, mongo collection to the handler
	visitorHandlers := handlers.NewVisitor(logger,ctx,visitorCollection)
	// Subrouter for the GET methods
	getSubrouter := router.Methods("GET").Subrouter()
	getSubrouter.HandleFunc("/visitor",visitorHandlers.GetVisitor)
	// Subrouter for the POST methods
	postSubrouter := router.Methods("POST").Subrouter()
	postSubrouter.HandleFunc("/visitor",visitorHandlers.AddVisitor)
	
	ops := middleware.RedocOpts{SpecURL: "/swagger.yml"}
	sh := middleware.Redoc(ops,nil)
	getSubrouter.Handle("/docs",sh)
	getSubrouter.Handle("/swagger.yml",http.FileServer(http.Dir("./")))
	
	// create a new server for the visitor service
	visitorServer := http.Server{
		Addr:         ":9191",
		Handler:      router,                // set the default router
		ErrorLog:     logger,                 		// set the logger for the server
		ReadTimeout:  5 * time.Second,   			// max time to read request from the client
		WriteTimeout: 10 * time.Second,  			// max time to write response to the client
		IdleTimeout:  120 * time.Second, 			// max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		logger.Println("Started server on port 9191")

		err := visitorServer.ListenAndServe()
		if err != nil {
			logger.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	logger.Println("Got signal:", sig)

	visitorServer.Shutdown(ctx)
}