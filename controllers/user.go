package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"todo/database"
	"todo/repository"
	"todo/structs"
)

func RegisterUser(c *gin.Context) {
	var user structs.User

	err := c.ShouldBindJSON(&user)

	user.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	user.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	if err != nil {
		panic(err)
	}
	
	err = repository.RegisterUser(database.DbConnection, user)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
	})
}

func DecryptJWT(token string) (map[string]interface{}, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return structs.PrivateKey, nil
	})

	if err != nil {
		return map[string]interface{}{}, err
	}

	if !parsedToken.Valid {
		return map[string]interface{}{}, err
	}
	return parsedToken.Claims.(jwt.MapClaims), nil
}

func LoginUser(c *gin.Context) {
	var userRequest structs.User

	err := c.ShouldBindJSON(&userRequest)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid input",
		})
		return
	}

	if userRequest.Email == "" || userRequest.Password == "" {
		c.JSON(400, gin.H{
			"message": "invalid input",
		})
		return
	}

	var user structs.User

	user, err = repository.LoginUser(database.DbConnection, userRequest.Email)

	if err != nil {
		panic(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password))

	if err != nil {
		c.JSON(400, gin.H{
			"message": "wrong email/password",
		})
		return
	}

	token, err := user.GenerateToken()

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user_id": user.ID,
	})
}