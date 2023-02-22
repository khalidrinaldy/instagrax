package structs

import "time"

type Comment struct {
	Id        string    `json:"id"`
	Text      string    `json:"text"`
	PostId    string    `json:"post_id"`
	UserId    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
