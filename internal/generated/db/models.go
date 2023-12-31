// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"time"
)

type LogLine struct {
	ID          int64     `json:"id"`
	Value       float64   `json:"value"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	TagID       int64     `json:"tagId"`
}

type Tag struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
