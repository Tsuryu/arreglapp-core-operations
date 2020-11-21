package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// ServiceRequest : request a new service transaction
type ServiceRequest struct {
	_ID             primitive.ObjectID `bson:"_id,omitempty"`
	ID              string             `json:"id,omitempty"`
	Username        string             `json:"username,omitempty"`
	Type            string             `json:"type,omitempty"`
	Title           string             `json:"title,omitempty"`
	Description     string             `json:"description,omitempty"`
	Location        Location           `json:"location,omitempty"`
	OperationType   OperationType      `json:"operation_type,omitempty"`
	UserContactInfo UserContactInfo    `json:"user_contact_info,omitempty"`
	Chats           []UserContactInfo  `json:"chats,omitempty"`
	Budget          *Budget            `json:"budget,omitempty"`
}
