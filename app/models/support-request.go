package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SupportRequest : request a new service transaction
type SupportRequest struct {
	_ID         primitive.ObjectID `bson:"_id,omitempty"`
	Username    string             `json:"username,omitempty"`
	Description string             `json:"description,omitempty"`
	Date        time.Time
}
