package domain

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Sub struct {
	ID     uuid.UUID `json:"id" swaggerignore:"true"`
	UserID uuid.UUID `json:"user_id"`

	ServiceName string    `json:"service_name"`
	Price       int64     `json:"price"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}

type SubJSONBody struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`

	ServiceName string `json:"service_name"`
	Price       int    `json:"price"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date,omitempty"`
}

const (
	TimeLayout = "01-2006"
)

func (s *Sub) UnmarshalJSON(b []byte) error {
	const op = "Sub.UnmarshalJSON"

	var req SubJSONBody
	if err := json.Unmarshal(b, &req); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	s.ServiceName = req.ServiceName
	s.Price = int64(req.Price)

	var err error

	s.UserID, err = uuid.Parse(req.UserID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	s.StartDate, err = time.Parse(TimeLayout, req.StartDate)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	endDate, err := time.Parse(TimeLayout, req.EndDate)
	if len(req.EndDate) != 0 && err != nil {
		return fmt.Errorf("%s: %w", op, err)
	} else if err == nil {
		s.EndDate = endDate
	}

	return nil
}

func (s *Sub) MarshalJSON() ([]byte, error) {
	const op = "Sub.MarshalJSON"

	req := SubJSONBody{
		ID:          s.ID.String(),
		UserID:      s.UserID.String(),
		ServiceName: s.ServiceName,
		Price:       int(s.Price),
		StartDate:   s.StartDate.Format(TimeLayout),
	}

	if !s.EndDate.IsZero() {
		req.EndDate = s.EndDate.Format(TimeLayout)
	}

	data, err := json.Marshal(&req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return data, nil
}
