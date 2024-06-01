package generateJwtToken

import (
	"fmt"
	"golang-backend-template/lib/getEnv"
	"golang-backend-template/model"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJwtToken(id uint, role string, expiredAt time.Time) (string, error) {
	claim := model.JWT{
		ID:   id,
		Role: role,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expiredAt.Unix(),
			Subject:   "Golang Template Backend",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	stringToken, err := token.SignedString([]byte(getEnv.GetEnvVariable("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Bearer %s", stringToken), nil
}
