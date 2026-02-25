package model

import "time"

// Wallet represents a user's wallet with balance and timestamps for
// creation and updates.
type Wallet struct {
	Mobile    string    `json:"mobile"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
