package structs

import "time"

type Like struct {
	Id        string    `json:"id"`
	PostId    string    `json:"post_id"`
	UserId    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
