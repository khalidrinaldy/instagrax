package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"instagrax/database"
	"instagrax/helper"
	"instagrax/repository"
	"instagrax/structs"
	"net/http"
	"os"
	"strings"
)

func GenerateToken(userId string) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = userId
	claims["authorized"] = true
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(os.Getenv("TOKEN_SECRET"))
}

func ValidateToken(c *gin.Context) (*jwt.Token, error) {
	tokenString, err := ExtractToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": err.Error(),
			"data":    map[string]string{},
		})
		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractToken(c *gin.Context) (string, error) {
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1], nil
	}
	return "", errors.New("invalid token")
}

func ExtractTokenID(c *gin.Context) string {
	token, err := ValidateToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": err.Error(),
			"data":    map[string]string{},
		})
		return err.Error()
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		id := claims["id"].(string)
		return id
	}
	return ""
}

func Login(c *gin.Context) {
	request := struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Bad request body",
			"data":    err,
		})
		return
	}

	if !helper.IsEmailValid(request.Email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "invalid email format",
			"data":    err,
		})
		return
	}

	if strings.TrimSpace(request.Email) == "" || strings.TrimSpace(request.Password) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "email or password tidak boleh kosong",
			"data":    err,
		})
		return
	}

	if len(strings.TrimSpace(request.Password)) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "password minimal 8 karakter",
			"data":    err,
		})
		return
	}

	user, err := repository.CheckEmail(database.DbConnection, request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": err.Error(),
			"data":    map[string]string{},
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "password salah",
			"data":    map[string]string{},
		})
		return
	}

	token, err := GenerateToken(user.Id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "gagal generate token",
			"data":    map[string]string{},
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "login berhasil",
		"result": map[string]string{
			"token": token,
		},
	})
}

func Register(c *gin.Context) {
	var user structs.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Bad request body",
			"data":    err,
		})
		return
	}

	if !helper.IsEmailValid(user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "invalid email format",
			"data":    err,
		})
		return
	}

	if strings.TrimSpace(user.Email) == "" || strings.TrimSpace(user.Password) == "" || strings.TrimSpace(user.Username) == "" || strings.TrimSpace(user.Name) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "mohon isi semua field",
			"data":    err,
		})
		return
	}

	if len(strings.TrimSpace(user.Password)) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "password minimal 8 karakter",
			"data":    err,
		})
		return
	}

	_, err = repository.CheckEmail(database.DbConnection, user.Email)
	if err == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "email sudah terdaftar",
			"data":    map[string]string{},
		})
		return
	}

	generatedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(generatedPassword)

	err = repository.Register(database.DbConnection, user)
	if err != nil {
		c.JSON(http.StatusRequestTimeout, gin.H{
			"code":    http.StatusRequestTimeout,
			"message": "error in database",
			"data":    map[string]string{},
		})
		return
	}

	registeredUser, _ := repository.CheckEmail(database.DbConnection, user.Email)
	generatedToken, err := GenerateToken(registeredUser.Id)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "register berhasil",
		"data": map[string]interface{}{
			"user":  registeredUser,
			"token": generatedToken,
		},
	})
}
