package model

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	UserId      uint   `gorm:"not null"`
}

type JoinedReadPost struct {
	ID          uint      `json:"post_id"`
	Username    string    `json:"username"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

type PostSearchFilter struct {
	Username string
}

type PostBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type PostAddParam struct {
	Title       string
	Description string
	UserId      uint
}
