package domain

import (
	"encoding/json"
	"fmt"
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

type SubJsonBody struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`

	ServiceName string `json:"service_name"`
	Price       int    `json:"price"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

const (
	TimeLayout = "01-2006"
)

func (s *Sub) UnmarshalJSON(b []byte) error {
	const op = "Sub.UnmarshalJSON"

	var req SubJsonBody
	if err := json.Unmarshal(b, &req); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	s.ServiceName = req.ServiceName
	s.Price = int64(req.Price)

	var err error

	s.UserId, err = uuid.Parse(req.UserId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	s.StartDate, err = time.Parse(TimeLayout, req.StartDate)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	s.EndDate, err = time.Parse(TimeLayout, req.EndDate)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Sub) MarshalJSON() ([]byte, error) {
	const op = "Sub.MarshalJSON"

	req := SubJsonBody{
		Id:          s.Id.String(),
		UserId:      s.UserId.String(),
		ServiceName: s.ServiceName,
		Price:       int(s.Price),
		StartDate:   s.StartDate.Format(TimeLayout),
		EndDate:     s.EndDate.Format(TimeLayout),
	}

	data, err := json.Marshal(&req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return data, nil
}
