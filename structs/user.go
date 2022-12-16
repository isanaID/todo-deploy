package structs

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64  `json:"ID"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewUser(name, email, password string) *User {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return &User{
		Name:      name,
		Email:     email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
}

var PrivateKey = []byte("secret")

func (u *User) GenerateToken() (string, error) {
	claims := jwt.MapClaims{
		"user_id": u.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iss":     "isana",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(PrivateKey)
}