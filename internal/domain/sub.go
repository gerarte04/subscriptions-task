package domain

import (
	"time"

	"github.com/google/uuid"
)

type Sub struct {
	Id     uuid.UUID `json:"id" swaggerignore:"true"`
	UserId uuid.UUID `json:"user_id"`

	ServiceName string    `json:"service_name"`
	Price       int64     `json:"price"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}
