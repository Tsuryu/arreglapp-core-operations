package models

// UserContactInfo : chat between client and provider
type UserContactInfo struct {
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Username  string `json:"username,omitempty"`
	Confirmed bool   `json:"confirmed"`
}
