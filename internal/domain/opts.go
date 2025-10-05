package domain

import "github.com/google/uuid"

type FilterOpts struct {
	UserId      uuid.UUID
	ServiceName string
	PageToken   uuid.UUID
	PageSize    int
}
