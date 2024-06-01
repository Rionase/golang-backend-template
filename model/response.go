package model

import "time"

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type JwtResponse struct {
	Token     string    `json:"authorization"`
	ExpiredAt time.Time `json:"expiredAt"`
	Message   string    `json:"message"`
}
