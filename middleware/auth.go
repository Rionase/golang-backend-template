package middleware

import (
	"golang-backend-template/lib/contains"
	"golang-backend-template/lib/getEnv"
	"golang-backend-template/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func WithAuth(role ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenText := c.GetHeader("Authorization")
		if tokenText == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorResponse{
				Error: "user not login",
			})
			return
		}
		tokenText = tokenText[len("Bearer "):]
		claim := &model.JWT{}
		token, err := jwt.ParseWithClaims(tokenText, claim, func(t *jwt.Token) (interface{}, error) {
			return []byte(getEnv.GetEnvVariable("JWT_SECRET_KEY")), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorResponse{
					Error: "signiture is invalid",
				})
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorResponse{
				Error: "token is not valid",
			})
			return
		}

		if !contains.Contains(claim.Role, role) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorResponse{
				Error: "role unauthorized",
			})
			return
		}
		c.Set("id", claim.ID)
		c.Set("role", claim.Role)
		c.Next()
	}
}
