// Package classification visitor
//
// Documentation for Visitor API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import "visitor-registry/data"

// A list of products
// swagger:response visitorsResponse
type visitorResponse struct {
	// All current products
	// in: body
	Body []data.Visitor
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponse struct {
	// Description of the error
	// in: body
	Body GenericError
}

// ID returned after registring a visitor
// swagger:response addVisitorResponse
type addVisitorResponse struct {
	// The id of the visitor
	ID int `json:"id"`
}

//swagger:parameters addVisitor
type visitorRequest struct {
	// Visitor data structure for registering one
	// in : body
	// required : true
	Body data.Visitor
}


//swagger:parameters getVisitor
type getVisitor struct {
	// in: query
	SDKVersion uint `json:"sdkversion"`
	// in: query
	ConfigID string `json:"configid"`
	// in: query
	ProcessName string `json:"processname"`
	// in: query
	AppName string `json:"appname"`
	// in: query
	BundleID string `json:"bundleid"`
	// in: query
	ID string `json:"deviceid"`
	// in: query
	OSVersion string `json:"osversion"`
} 