package entity

import "time"

type Book struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Author      string     `json:"author"`
	Description string     `json:"description,omitempty"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"-"`
}
