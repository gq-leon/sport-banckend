package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/gq-leon/sport-backend/domain"
	"github.com/gq-leon/sport-backend/internal/tokenutil"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		t := strings.Split(auth, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := tokenutil.IsAuthorized(authToken, secret)
			if authorized {
				userID, err := tokenutil.ExtractIDFromToken(authToken, secret)
				if err != nil {
					domain.ErrorResponse(c, http.StatusUnauthorized, err)
					c.Abort()
					return
				}
				c.Set("x-user-id", userID)
				c.Next()
				return
			}
			domain.ErrorResponse(c, http.StatusUnauthorized, err)
			c.Abort()
			return
		}
		domain.ErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("not authorized"))
		c.Abort()
	}
}
