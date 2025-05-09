package utils

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/template/go-backend-gin-orm/config"
	"github.com/template/go-backend-gin-orm/dtos"
)

func GenerateJWT(userID uint, email string) (string, error) {
	env := config.NewEnv(".env", true)
	// claims := jwt.MapClaims{
	// 	"user_id": userID,
	// 	"exp":     time.Now().Add(time.Hour * time.Duration(int(env.JwtExpire))).Unix(),
	// }
	claims := dtos.UserClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(int(env.JwtExpire)))), // Token expires in 24 hours
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(env.JwtSecret))
}

func ValidateToken(tokenString string) (*dtos.UserClaims, error) {
	env := config.NewEnv(".env", true)
	token, err := jwt.ParseWithClaims(tokenString, &dtos.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(env.JwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*dtos.UserClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func GetUserClaimsFromContext(ctx *gin.Context) (*dtos.UserClaims, bool) {
	claims, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return nil, false
	}

	userClaims, ok := claims.(*dtos.UserClaims) // Cast to type dtos.UserClaims
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user data"})
		return nil, false
	}

	return userClaims, true
}
