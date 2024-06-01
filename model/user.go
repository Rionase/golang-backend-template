package model

import (
	"gorm.io/gorm"
)

// THINK OF IT'S RELATION LIKE : [
// 	{
// 		ID : 1
// 		Username : "Andy",
// 		Password : "Andy123",
// 		Role: "user",
// 		Post : [
// 			{...}, {...} ,...
// 		]
// 	}, ...
// ]

type User struct {
	gorm.Model
	Username string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	Role     string `gorm:"not null"`
	Post     []Post `gorm:"foreignKey:UserId;references:ID"`
	// foreignKey IS FOR THE CHILD COLUMN ( POST ) THAT WILL BE FK, references IS FOR THE PARENT COLUMN ( USER ) THAT IS PK

}

type UserLoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserAuthBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
