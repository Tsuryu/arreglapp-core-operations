package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// OperationType : operation type definition
type OperationType struct {
	_ID         primitive.ObjectID `bson:"_id,omitempty"`
	ID          string
	Name        string
	Description string
	Active      bool
	IconCode    int64  `bson:"icon_code,omitempty" json:"icon_code,omitempty"`
	IconFamily  string `bson:"icon_family,omitempty" json:"icon_family,omitempty"`
}
