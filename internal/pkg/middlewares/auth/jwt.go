package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"mall/internal/pkg/constant"
	"mall/pkg/auth/jwtauth"
)

// JWTStrategy defines jwt bearer authentication strategy.
type JWTStrategy struct {
	auth *jwtauth.JwtAuth
}

func (j *JWTStrategy) AuthFunc() gin.HandlerFunc {
	return j.JwtMiddleware()
}

func (j *JWTStrategy) JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": constant.AccountEmptyAuthHeader,
			})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": constant.AccountInvalidAuthHeader,
			})
			return
		}

		tokenString := parts[1]
		claims, err := j.auth.ParseJwtToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": constant.AccountInvalidToken,
			})
			return
		}

		if claims.HasExpired() {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": constant.AccountTokenExpired,
			})
			return
		}

		c.Set("identity", claims.Identity)
		c.Next()
	}
}
