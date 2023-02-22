package repository

import (
	"database/sql"
	"fmt"
	"instagrax/structs"
)

func GetUsersAllPosts(db *sql.DB, userId string) (posts []structs.PostToShow, err error) {
	sqlPosts := `select * from post where user_id=$1`
	rowsPosts, err := db.Query(sqlPosts, userId)
	if err != nil {
		panic(err)
	}
	defer rowsPosts.Close()

	for rowsPosts.Next() {
		var post structs.PostToShow
		err = rowsPosts.Scan(&post.Id, &post.ImageUrl, &post.Caption, &post.UserId, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			panic(err)
		}

		sqlLikes := "select count(*) from likes where post_id=$1"
		rowLike := db.QueryRow(sqlLikes, post.Id)
		rowLike.Scan(&post.Likes)

		sqlComments := "select count(*) from comments where post_id=$1"
		rowComment := db.QueryRow(sqlComments, post.Id)
		rowComment.Scan(&post.Comments)

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
