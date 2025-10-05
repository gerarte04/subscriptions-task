package domain

import "github.com/google/uuid"

type Summary struct {
	UserId      uuid.UUID `json:"user_id"`
	ServiceName string    `json:"service_name,omitempty"`
	TotalPrice  int       `json:"total_price"`
}
