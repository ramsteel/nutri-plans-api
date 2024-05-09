package dto

import "time"

type PaginationRequest struct {
	Limit int        `json:"limit" validate:"required,gt=0"`
	Page  int        `json:"offset" validate:"required,gt=0"`
	From  *time.Time `json:"from"`
	To    *time.Time `json:"to"`
}
