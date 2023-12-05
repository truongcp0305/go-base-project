package router

import (
	"go-project/library"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenS := c.GetHeader("Authorization")
		if tokenS == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		token, err := jwt.Parse(tokenS, func(token *jwt.Token) (interface{}, error) {
			return library.GetJwtKey(), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Set("user", token.Claims.(jwt.MapClaims)["user"])
		c.Next()
	}
}
