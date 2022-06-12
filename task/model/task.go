package model

import (
	"github.com/emvi/null"
)
type Task struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Brief string `json:"brief"`
	BoardId int64 `json:"board_id"`
	Status string `json:"status"`
	ParentId string `json:"parent_id"`
	IsActive bool `json:"is_active"`
	CreatedAt string `json:"created_at"`
	DeletedAt null.String `json:"deleted_at"`
}

