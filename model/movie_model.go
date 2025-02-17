package model

import "time"

type Movie struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Score       float32   `json:"score"`
	Picture     []byte    `json:"picture"`
	ReleaseDate time.Time `json:"release_date"`
	Synopsis    string    `json:"synopsis"`
	Publisher   string    `json:"publisher"` //TODO Change this later
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}