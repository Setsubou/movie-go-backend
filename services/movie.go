package movie_service

import "time"

type Movie struct {
	Id          int `json:"id"`
	Title       string `json:"title"`
	Score       float32 `json:"score"`
	Picture     []byte `json:"picture"`
	ReleaseDate time.Time `json:"release_date"`
	Synopsis    string `json:"synopsis"`
	Publisher   string `json:"publisher"` //TODO Change this later
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func GetById(id int) *Movie {
	return &Movie{
		Id:          id,
		Title:       "The Matrix",
		Score:       8.0,
		Picture:     []byte{},
		ReleaseDate: time.Now(),
		Synopsis:    "A computer hacker learns from mysterious rebels about the true nature of his reality and his role in the war against its controllers.",
		Publisher:   "Lana Wachowski, Lilly Wachowski",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}