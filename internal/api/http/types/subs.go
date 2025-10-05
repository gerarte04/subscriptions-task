package types

import (
	"encoding/json"
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"path"
	"strconv"
	"subs-service/internal/config"
	"subs-service/internal/domain"

	"github.com/google/uuid"
)

func checkPrice(price int64, cfg config.DataConfig) bool {
	return price > 0 && price <= cfg.MaxPrice
}

func checkServiceName(name string, cfg config.DataConfig) bool {
	return len(name) > 0 && len(name) <= cfg.MaxServiceNameLength
}

func checkPageSize(size int, cfg config.DataConfig) bool {
	return size >= 0 && size <= cfg.MaxPageSize
}

// Requests ----------------------------------------------------------------------

type GetSubRequest struct {
	Id uuid.UUID
}

func CreateGetSubRequest(r *http.Request) (*GetSubRequest, error) {
	const op = "CreateGetSubRequest"

	id, err := uuid.Parse(path.Base(r.URL.Path))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &GetSubRequest{Id: id}, nil
}

type PostSubRequest struct {
	Sub domain.Sub
}

func CreatePostSubRequest(r *http.Request, cfg config.DataConfig) (*PostSubRequest, error) {
	const op = "CreatePostSubRequest"

	var req PostSubRequest

	if err := json.NewDecoder(r.Body).Decode(&req.Sub); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if !checkPrice(req.Sub.Price, cfg) {
		return nil, fmt.Errorf("%s: %w", op, ErrBadPriceValue)
	} else if !checkServiceName(req.Sub.ServiceName, cfg) {
		return nil, fmt.Errorf("%s: %w", op, ErrBadServiceNameLength)
	}

	return &req, nil
}

type PutSubRequest struct {
	Id  uuid.UUID
	Sub domain.Sub
}

func CreatePutSubRequest(r *http.Request, cfg config.DataConfig) (*PutSubRequest, error) {
	const op = "CreatePutSubRequest"

	var req PutSubRequest

	id, err := uuid.Parse(path.Base(r.URL.Path))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	req.Id = id

	if err := json.NewDecoder(r.Body).Decode(&req.Sub); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if !checkPrice(req.Sub.Price, cfg) {
		return nil, fmt.Errorf("%s: %w", op, ErrBadPriceValue)
	} else if !checkServiceName(req.Sub.ServiceName, cfg) {
		return nil, fmt.Errorf("%s: %w", op, ErrBadServiceNameLength)
	}

	return &req, nil
}

type DeleteSubRequest struct {
	Id uuid.UUID
}

func CreateDeleteSubRequest(r *http.Request) (*DeleteSubRequest, error) {
	const op = "DeleteSubRequest"

	id, err := uuid.Parse(path.Base(r.URL.Path))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &DeleteSubRequest{Id: id}, nil
}

type ListSubsRequest struct {
	Opts domain.FilterOpts
}

func CreateListSubsRequest(r *http.Request, cfg config.DataConfig) (*ListSubsRequest, error) {
	const op = "CreateListSubsRequest"

	req := ListSubsRequest{
		Opts: domain.FilterOpts{
			PageSize:  cfg.DefaultPageSize,
			PageToken: uuid.Nil,
		},
	}

	var err error

	req.Opts.UserId, err = uuid.Parse(r.URL.Query().Get("user_id"))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if serviceName := r.URL.Query().Get("service_name"); len(serviceName) != 0 {
		req.Opts.ServiceName = serviceName
	}

	if pageSize, err := strconv.Atoi(r.URL.Query().Get("page_size")); err == nil && checkPageSize(pageSize, cfg) {
		req.Opts.PageSize = pageSize
	}

	if pageToken, err := uuid.Parse(r.URL.Query().Get("page_token")); err == nil {
		req.Opts.PageToken = pageToken
	}

	return &req, nil
}

type GetSummaryRequest struct {
	Opts domain.FilterOpts
}

func CreateGetSummaryRequest(r *http.Request, cfg config.DataConfig) (*GetSummaryRequest, error) {
	const op = "CreateGetSummaryRequest"

	var req GetSummaryRequest
	var err error

	req.Opts.UserId, err = uuid.Parse(r.URL.Query().Get("user_id"))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if serviceName := r.URL.Query().Get("service_name"); len(serviceName) != 0 {
		req.Opts.ServiceName = serviceName
	}

	return &req, nil
}

// Responses ---------------------------------------------------------------------

type ListSubsResponse struct {
	Subs          []*domain.Sub `json:"subs"`
	NextPageToken string        `json:"next_page_token"`
}

func CreateListSubsResponse(subs []*domain.Sub) *ListSubsResponse {
	return &ListSubsResponse{
		Subs: subs,
		NextPageToken: subs[len(subs) - 1].Id.String(),
	}
}
