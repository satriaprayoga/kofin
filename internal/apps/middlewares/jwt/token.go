package token

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/satriaprayoga/kofin/internal/config"
	"github.com/satriaprayoga/kofin/internal/store"
)

type JwtSetting struct {
	expired   int
	secretKey string
}

var jwtSetting *JwtSetting

func NewConfigAuth(cnf config.Config) {
	jwtSetting = &JwtSetting{
		expired:   cnf.JWTExpired,
		secretKey: cnf.App.JwtSecret,
	}
}

func GetJwtSetting() *JwtSetting {
	return jwtSetting
}

func GenerateToken(user store.KUser) (string, error) {
	tokenExpired := jwtSetting.expired
	secretKey := jwtSetting.secretKey

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.UserID,
		"username": user.Username,
		"role":     user.RoleID,
		"iat":      time.Now().Unix(),
		"eat":      time.Now().Add(time.Hour * time.Duration(tokenExpired)).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func ValidateToken(c *gin.Context) error {
	token, err := getToken(c)
	if err != nil {
		return err
	}
	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("invalid token provided")
}

func getToken(c *gin.Context) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtSetting.secretKey), nil
	})

	return token, err

}

func getTokenFromRequest(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}
