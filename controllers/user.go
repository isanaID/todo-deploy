package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
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

	regexEmail := regexp.MustCompile(`^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`)

	if !regexEmail.MatchString(user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Email is not valid",
		})
		return
	}

	if len(user.Password) < 5 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Password must be at least 5 characters",
		})
		return
	}

	if err != nil {
		panic(err)
	}

	err = repository.CheckEmail(database.DbConnection, user.Email)

	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Email already exists",
		})
		return
	}

	user.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	user.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	
	err = repository.RegisterUser(database.DbConnection, user)

	if err != nil {
		panic(err)
	}

	type response struct {
		Name string `json:"name"`
		Email string `json:"email"`
		Created_at string `json:"created_at"`
		Updated_at string `json:"updated_at"`
	}

	var res response
	res.Name = user.Name
	res.Email = user.Email
	res.Created_at = user.CreatedAt
	res.Updated_at = user.UpdatedAt

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "User registered successfully",
		"data": res,
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

func IdUser(c *gin.Context) int {

	token := c.Request.Header.Get("Authorization")
	token = token[7:]
	claims, err := DecryptJWT(token)
	if err != nil {
		fmt.Println(err)
	}
	id := int(claims["user_id"].(float64))
	return id
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
		c.JSON(400, gin.H{
			"message": "wrong email/password",
		})
		return
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
		"data": gin.H{
			"token": token,
		},
	})
}