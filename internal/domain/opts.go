package domain

import "github.com/google/uuid"

type FilterOpts struct {
	UserID      uuid.UUID
	ServiceName string
	PageToken   uuid.UUID
	PageSize    int
}
