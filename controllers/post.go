package controllers

import (
	"github.com/gin-gonic/gin"
	"instagrax/database"
	"instagrax/repository"
	"instagrax/structs"
	"net/http"
	"strings"
)

func GetUsersAllPosts(c *gin.Context) {
	id := c.Param("id")
	posts, err := repository.GetUsersAllPosts(database.DbConnection, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "err",
			"data":    err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "get all post by user id",
		"data":    posts,
	})
}

func CreatePost(c *gin.Context) {
	var post structs.Post

	err := c.ShouldBindJSON(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Bad request body",
			"data":    err,
		})
		return
	}

	if strings.TrimSpace(post.ImageUrl) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "mohon isi semua field",
			"data":    err,
		})
		return
	}

	userId := ExtractTokenID(c)
	post.UserId = userId
	_, err = repository.CheckId(database.DbConnection, userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": err.Error(),
			"data":    map[string]string{},
		})
		return
	}

	err = repository.CreatePost(database.DbConnection, post)
	if err != nil {
		c.JSON(http.StatusRequestTimeout, gin.H{
			"code":    http.StatusRequestTimeout,
			"message": "error in database",
			"data":    map[string]string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "buat post berhasil",
		"data":    post,
	})
}

func EditPost(c *gin.Context) {
	var post structs.Post

	err := c.ShouldBindJSON(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Bad request body",
			"data":    err,
		})
		return
	}

	if strings.TrimSpace(post.ImageUrl) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "mohon isi semua field",
			"data":    err,
		})
		return
	}

	userId := ExtractTokenID(c)
	post.UserId = userId
	_, err = repository.CheckId(database.DbConnection, userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": err.Error(),
			"data":    map[string]string{},
		})
		return
	}

	err = repository.EditPost(database.DbConnection, post)
	if err != nil {
		c.JSON(http.StatusRequestTimeout, gin.H{
			"code":    http.StatusRequestTimeout,
			"message": "error in database",
			"data":    map[string]string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "edit post berhasil",
		"data":    map[string]interface{}{},
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	userId := ExtractTokenID(c)
	_, err := repository.CheckId(database.DbConnection, userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": err.Error(),
			"data":    map[string]string{},
		})
		return
	}

	err = repository.DeletePost(database.DbConnection, id)
	if err != nil {
		c.JSON(http.StatusRequestTimeout, gin.H{
			"code":    http.StatusRequestTimeout,
			"message": "error in database",
			"data":    map[string]string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "delete post berhasil",
		"data":    map[string]interface{}{},
	})
}

func AddLike(c *gin.Context) {
	var like structs.Like
	id := c.Param("id")
	like.PostId = id

	userId := ExtractTokenID(c)
	like.UserId = userId
	_, err := repository.CheckId(database.DbConnection, userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": err.Error(),
			"data":    map[string]string{},
		})
		return
	}

	err = repository.AddLike(database.DbConnection, like)
	if err != nil {
		c.JSON(http.StatusRequestTimeout, gin.H{
			"code":    http.StatusRequestTimeout,
			"message": "error in database",
			"data":    map[string]string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "like post berhasil",
		"data":    map[string]interface{}{},
	})
}

func DeleteLike(c *gin.Context) {
	id := c.Param("id")

	userId := ExtractTokenID(c)
	_, err := repository.CheckId(database.DbConnection, userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": err.Error(),
			"data":    map[string]string{},
		})
		return
	}

	err = repository.DeleteLike(database.DbConnection, id)
	if err != nil {
		c.JSON(http.StatusRequestTimeout, gin.H{
			"code":    http.StatusRequestTimeout,
			"message": "error in database",
			"data":    map[string]string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "unlike post berhasil",
		"data":    map[string]interface{}{},
	})
}

func AddComment(c *gin.Context) {
	var comment structs.Comment
	id := c.Param("id")
	comment.PostId = id

	err := c.ShouldBindJSON(&comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Bad request body",
			"data":    err,
		})
		return
	}

	userId := ExtractTokenID(c)
	comment.UserId = userId
	_, err = repository.CheckId(database.DbConnection, userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": err.Error(),
			"data":    map[string]string{},
		})
		return
	}

	err = repository.AddComment(database.DbConnection, comment)
	if err != nil {
		c.JSON(http.StatusRequestTimeout, gin.H{
			"code":    http.StatusRequestTimeout,
			"message": "error in database",
			"data":    map[string]string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "comment post berhasil",
		"data":    map[string]interface{}{},
	})
}

func DeleteComment(c *gin.Context) {
	id := c.Param("id")

	userId := ExtractTokenID(c)
	_, err := repository.CheckId(database.DbConnection, userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": err.Error(),
			"data":    map[string]string{},
		})
		return
	}

	err = repository.DeleteComment(database.DbConnection, id)
	if err != nil {
		c.JSON(http.StatusRequestTimeout, gin.H{
			"code":    http.StatusRequestTimeout,
			"message": "error in database",
			"data":    map[string]string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "delete comment post berhasil",
		"data":    map[string]interface{}{},
	})
}