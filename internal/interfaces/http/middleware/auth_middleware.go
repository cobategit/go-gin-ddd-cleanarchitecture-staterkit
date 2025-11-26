package middleware

import (
	"net/http"
	"strconv"

	"github.com/cobategit/go-gin-ddd-cleanarchitecture-staterkit/internal/infrastructure/security"
	"github.com/cobategit/go-gin-ddd-cleanarchitecture-staterkit/pkg/response"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtService *security.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		// Expect "Bearer <token>"
		const prefix = "Bearer "
		if len(tokenString) < len(prefix) || tokenString[:len(prefix)] != prefix {
			response.ToJson(c, http.StatusUnauthorized, false, "missing or invalid Authorization header", nil)
			c.Abort()
			return
		}
		tokenString = tokenString[len(prefix):]

		claims, err := jwtService.ParseToken(tokenString)
		if err != nil {
			response.ToJson(c, http.StatusUnauthorized, false, "invalid token", nil)
			c.Abort()
			return
		}

		sub, ok := claims["sub"].(float64) // jwt MapClaims stores numbers as float64
		if !ok {
			response.ToJson(c, http.StatusUnauthorized, false, "invalid token claims", nil)
			c.Abort()
			return
		}

		c.Set("userID", int64(sub))
		c.Set("userIDStr", strconv.FormatInt(int64(sub), 10))

		c.Next()
	}
}
