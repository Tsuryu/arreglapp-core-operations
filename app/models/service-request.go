package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// ServiceRequest : request a new service transaction
type ServiceRequest struct {
	_ID         primitive.ObjectID `bson:"_id,omitempty"`
	ID          string
	Username    string
	Type        string
	Title       string
	Description string
	Location    Location
}
