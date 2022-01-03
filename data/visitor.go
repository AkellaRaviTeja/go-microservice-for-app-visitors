package data

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/url"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Model for thevisitor info
type Visitor struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	SDKInfo struct { 
		SDKVersion  uint `json:"sdk_version,omitempty" bson:"sdk_version,omitempty"`
		ConfigID string `json:"config_id,omitempty" bson:"config_id,omitempty"`
		TickInterval float32 `json:"tick_interval,omitempty" bson:"tick_interval,omitempty"`
		SDKStartTime int64 `json:"sdk_start_time,omitempty" bson:"sdk_start_time,omitempty"`
	} `json:"sdk_info,omitempty" bson:"sdk_info,omitempty"`
	SessionInfo struct {
		ID string `json:"id,omitempty" bson:"id,omitempty"`
		LaunchTime int64 `json:"launch_time,omitempty" bson:"launch_time,omitempty"`
		StartupTime uint `json:"startup_time,omitempty" bson:"startup_time,omitempty"`
		IsBackground bool `json:"is_first_session,omitempty" bson:"is_first_session,omitempty"`
		IsFirstSession bool `json:"_id,omitempty" bson:"_id,omitempty"`
		Duration float64 `json:"duration,omitempty" bson:"duration,omitempty"`
		TotalDuration float64 `json:"total_duration,omitempty" bson:"total_duration,omitempty"`
		ProcessName string `json:"process_name,omitempty" bson:"process_name,omitempty"`
		LaunchType string `json:"launch_type,omitempty" bson:"launch_type,omitempty"`
	} `json:"session_info,omitempty" bson:"session_info,omitempty"`
	AppInfo struct {
		AppName string `json:"app_name,omitempty" bson:"app_name,omitempty"`
		AppID string `json:"app_id,omitempty" bson:"app_id,omitempty"`
		AppVersion string `json:"app_version,omitempty" bson:"app_version,omitempty"`
		BundleID string `json:"bundle_id,omitempty" bson:"bundle_id,omitempty"`
		AppMode string `json:"app_mode,omitempty" bson:"app_mode,omitempty"`
	} `json:"app_info,omitempty" bson:"app_info,omitempty"`
	DeviceInfo struct {
		ID string `json:"id,omitempty" bson:"id,omitempty"`
		HardwareModel string `json:"hardware_model,omitempty" bson:"hardware_model,omitempty"`
		Dimensions struct {
			Height uint `json:"height,omitempty" bson:"height,omitempty"`
			Width uint `json:"width,omitempty" bson:"width,omitempty"`
		} `json:"dimensions,omitempty" bson:"dimensions,omitempty"`
		OSVersion string `json:"os_version,omitempty" bson:"os_version,omitempty"`
		Platform string `json:"platform,omitempty" bson:"platform,omitempty"`
		DPI uint `json:"dpi,omitempty" bson:"dpi,omitempty"`
		DimensionsInPixels struct {
			Height uint `json:"height,omitempty" bson:"height,omitempty"`
			Width uint `json:"width,omitempty" bson:"width,omitempty"`
		} `json:"dimensions_in_pixels,omitempty" bson:"dimensions_in_pixels,omitempty"`
	} `json:"device_info,omitempty" bson:"device_info,omitempty"`
}

// Adds one visitor to the visitor collection in the Apxor Database
func AddOneVisitor(requestData io.ReadCloser, logger *log.Logger, ctx context.Context, visitorCollection *mongo.Collection) (*mongo.InsertOneResult,error) {
	var visitor Visitor
	// Decode the request to map it to the Visitor struct
	json.NewDecoder(requestData).Decode(&visitor)
	// Insert the visitor
	result, err := visitorCollection.InsertOne(ctx,visitor)
	if err != nil {
		logger.Printf("Failed to insert visitor %v\n",err)
		return result,err
	}
	return result,nil
}

// This is a custom imlementation to create a filter from the query params
// Since the filter can be complex with filter not sent, this custom builder creates the conditions only when required
func createFilterFromParams(params url.Values) []primitive.M{
	var AndQuery []primitive.M
	// Read the query params
	sdk_version,_ := strconv.Atoi(params.Get("sdkversion"))
	config_id := params.Get("configid")
	process_name := params.Get("processname")
	app_name := params.Get("appname")
	bundle_id := params.Get("bundleid")
	device_id := params.Get("deviceid")
	os_version := params.Get("osversion")

	// For each of the query param if it exists, append it to the filter conditions.
	if sdk_version != 0 {
		currentFilter := bson.M{"sdk_info.sdk_version":sdk_version}
		AndQuery = append(AndQuery, currentFilter)
	}
	
	if config_id != "" {
		currentFilter := bson.M{"sdk_info.config_id" : config_id}
		AndQuery = append(AndQuery, currentFilter)
	}
	
	if process_name != "" {
		currentFilter := bson.M{"session_info.process_name" : process_name}
		AndQuery = append(AndQuery, currentFilter)
	}
	
	if app_name != "" {
		currentFilter := bson.M{"app_info.app_name" : app_name}
		AndQuery = append(AndQuery, currentFilter)
	}
	
	if bundle_id != "" {
		currentFilter := bson.M{"app_info.bundle_id" : bundle_id}
		AndQuery = append(AndQuery, currentFilter)
	}
	
	if device_id != "" {
		currentFilter := bson.M{"device_info.id" : device_id}
		AndQuery = append(AndQuery, currentFilter)
	}
	
	if os_version != "" {
		currentFilter := bson.M{"device_info.os_version" : os_version}
		AndQuery = append(AndQuery, currentFilter)
	}
	return AndQuery
}

// Fetches the visitors from the visitor collection of Apxor database for a given filter
func GetVisitors(params url.Values,logger *log.Logger, ctx context.Context, visitorCollection *mongo.Collection) ([]Visitor,error) {
	var visitors []Visitor
	// Create the custom filter
	filter := createFilterFromParams(params)
	// Query the collection
	cursor, err := visitorCollection.Find(ctx,bson.M{"$and":filter})
	if err != nil {
		logger.Printf("Failed to fetch a visitor\n%v\n",err)
		return visitors,err
	}
	defer cursor.Close(ctx)

	// Query response can have multiple documents, decode each of them.
	for cursor.Next(ctx) {
		var visitor Visitor
		cursor.Decode(&visitor)
		visitors = append(visitors, visitor)
		logger.Printf("Visitors%v\n",visitor)
	}

	if err := cursor.Err(); err != nil {
		logger.Printf("Failed to fetch a visitor\n")
		return visitors,err
	}

	return visitors,nil
}