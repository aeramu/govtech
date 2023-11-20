package api

import "errors"

type GetProductListFilter struct {
	Search     string
	CategoryID int64
	SortColumn string
	SortType   string
	Page       int64
	Size       int64
}

type ReviewProductRequest struct {
	Rating  int32  `json:"rating"`
	Comment string `json:"comment"`
}

func (req ReviewProductRequest) Validate() error {
	if req.Rating < 1 || req.Rating > 5 {
		return errors.New("invalid rating range")
	}
	return nil
}

func (filter *GetProductListFilter) Validate() error {
	if filter.SortType != "" && filter.SortType != "asc" && filter.SortType != "desc" {
		return errors.New("invalid sort type")
	}
	if filter.SortColumn != "" && filter.SortColumn != "created_at" && filter.SortColumn != "rating" {
		return errors.New("invalid sort column")
	}
	if filter.Page <= 0 {
		filter.Page = 1
	}
	if filter.Size <= 0 {
		filter.Size = 10
	}
	return nil
}
