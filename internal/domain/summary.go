package domain

import "github.com/google/uuid"

type Summary struct {
	UserID      uuid.UUID `json:"user_id"`
	ServiceName string    `json:"service_name,omitempty"`
	TotalPrice  int       `json:"total_price"`
}
