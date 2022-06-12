package model

import (
	"github.com/emvi/null"
)
type Board struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	IsActive bool `json:"is_active"`
	CreatedAt string `json:"created_at"`
	DeletedAt null.String `json:"deleted_at"`
}

