// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"time"
)

type Hobby struct {
	ID        int64        `json:"id"`
	UserID    int64        `json:"user_id"`
	Name      string       `json:"name"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type User struct {
	ID             int64          	`json:"id"`
	Username       string         	`json:"username"`
	Name           string         	`json:"name"`
	Email          string         	`json:"email"`
	Password       string         	`json:"password"`
	ProfilePicture *string			`json:"profile_picture"`
	CreatedAt      time.Time      	`json:"created_at"`
	UpdatedAt      *time.Time   	`json:"updated_at"`
}
