package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"visitor-registry/data"

	"go.mongodb.org/mongo-driver/mongo"
)

// Visitors handler struct
type Visitors struct{
	logger *log.Logger
	ctx context.Context 
	visitorCollection *mongo.Collection 
}

// Inject the resources
func NewVisitor(logger *log.Logger, ctx context.Context, visitorCollection *mongo.Collection) *Visitors{
	return &Visitors{logger,ctx,visitorCollection}
}

// swagger:route GET /visitor GetAVisitor getVisitor
// Return a list of visitors from the database
// responses:
//	200: visitorsResponse
//  500: errorResponse

// Handler for fetching a visitor. Called for a GET method on /visitor path
func (v *Visitors) GetVisitor(responseWriter http.ResponseWriter, request *http.Request) {
	v.logger.Println("Get a visitors entry")
	responseWriter.Header().Add("content-type","application/json")
	queryParams := request.URL.Query()
	visitors,err := data.GetVisitors(queryParams,v.logger,v.ctx,v.visitorCollection)
	
	if err == nil {
		responseWriter.WriteHeader(http.StatusOK)
		json.NewEncoder(responseWriter).Encode(visitors)
	}else {
		responseWriter.WriteHeader(http.StatusInternalServerError)
	}
}

// swagger:route POST /visitor CreateAVisitor addVisitor
// Return a list of visitors from the database
// responses:
//	200: addVisitorResponse
//  500: errorResponse

// Handler to add a visitor. Called for a POST method on /visitor path
func (v *Visitors) AddVisitor(responseWriter http.ResponseWriter, request *http.Request) {
	v.logger.Println("Create a visitor entry")
	responseWriter.Header().Add("content-type","application/json")
	result,err := data.AddOneVisitor(request.Body, v.logger, v.ctx, v.visitorCollection)
	if err == nil {
		json.NewEncoder(responseWriter).Encode(result)
		responseWriter.WriteHeader(http.StatusCreated)
	}else{
		responseWriter.WriteHeader(http.StatusInternalServerError)
	}
}