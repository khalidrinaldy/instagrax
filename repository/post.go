package repository

import (
	"database/sql"
	"fmt"
	"instagrax/structs"
)

func GetUsersAllPosts(db *sql.DB, user_id string) (posts []structs.PostToShow, err error) {
	sql := `select p.id, p.image_url, p.caption, p.user_id, count(l.post_id) as likes, count(c.post_id) as comments, p.created_at, p.updated_at 
			from post p inner join likes l on p.id = l.post_id inner join comments c on p.id = c.post_id where p.user_id=$1 group by p.id`
	rows, err := db.Query(sql, user_id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	fmt.Println("ROWS")
	fmt.Println(rows)

	for rows.Next() {
		var post structs.PostToShow
		err = rows.Scan(&post.Id, &post.ImageUrl, &post.Caption, &post.UserId, &post.Likes, &post.Comments, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			panic(err)
		}
		posts = append(posts, post)
		fmt.Println("POST")
		fmt.Println(post)
	}

	return
}

func CreatePost(db *sql.DB, post structs.Post) error {
	sql := "insert into post (image_url, caption, user_id) values ($1,$2,$3)"
	err := db.QueryRow(sql, post.ImageUrl, post.Caption, post.UserId)
	return err.Err()
}

func EditPost(db *sql.DB, post structs.Post) error {
	sql := "update post set image_url=$1, caption=$2 where id=$3"
	err := db.QueryRow(sql, post.ImageUrl, post.Caption, post.UserId)
	return err.Err()
}

func DeletePost(db *sql.DB, id string) error {
	sql := "delete from post where id=$1"
	err := db.QueryRow(sql, id)
	return err.Err()
}
