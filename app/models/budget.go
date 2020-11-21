package models

import (
	"time"
)

// Budget : request intended cost
type Budget struct {
	Amount   float64
	Date     time.Time
	Username string
}
