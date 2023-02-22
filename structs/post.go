package structs

import "time"

type Post struct {
	Id        string    `json:"id"`
	ImageUrl  string    `json:"image_url"`
	Caption   string    `json:"caption"`
	UserId    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostToShow struct {
	Id        string    `json:"id"`
	ImageUrl  string    `json:"image_url"`
	Caption   string    `json:"caption"`
	UserId    string    `json:"user_id"`
	Likes     int       `json:"likes"`
	Comments  int       `json:"comments"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
