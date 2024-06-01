package service

import (
	"errors"
	"golang-backend-template/lib/generateJwtToken"
	"golang-backend-template/model"
	"golang-backend-template/repository"
	"time"

	"gorm.io/gorm"
)

type IAuthService interface {
	RegisterAdmin(userData model.UserAuthBody, expiredAt time.Time) (string, error)
	RegisterUser(userData model.UserAuthBody, expiredAt time.Time) (string, error)
	Login(userData model.UserLoginBody, expiredAt time.Time) (string, error)
	Logout(userData model.UserAuthBody) error
}

type authService struct {
	userRepo repository.IUserRepository
}

func NewAuthService(userRepo repository.IUserRepository) *authService {
	return &authService{userRepo: userRepo}
}

func (a authService) RegisterAdmin(userData model.UserAuthBody, expiredAt time.Time) (string, error) {
	if userData.Username == "" || userData.Password == "" {
		return "", errors.New("username or password shouldn't be empty")
	}
	userData.Role = "admin"
	result, err := a.userRepo.AddUser(userData)
	if err != nil {
		return "", err
	}
	token, err := generateJwtToken.GenerateJwtToken(result.ID, result.Role, expiredAt)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a authService) RegisterUser(userData model.UserAuthBody, expiredAt time.Time) (string, error) {
	if userData.Username == "" || userData.Password == "" {
		return "", errors.New("username or password shouldn't be empty")
	}
	userData.Role = "user"
	result, err := a.userRepo.AddUser(userData)
	if err != nil {
		return "", err
	}
	token, err := generateJwtToken.GenerateJwtToken(result.ID, result.Role, expiredAt)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a authService) Login(userData model.UserLoginBody, expiredAt time.Time) (string, error) {
	if userData.Username == "" || userData.Password == "" {
		return "", errors.New("username or password shouldn't be empty")
	}
	result, err := a.userRepo.CheckUserAvail(userData.Username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", errors.New("user not registered")
	} else if err != nil {
		return "", err
	}
	if userData.Password != result.Password {
		return "", errors.New("wrong password")
	}
	token, err := generateJwtToken.GenerateJwtToken(result.ID, result.Role, expiredAt)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u authService) Logout(userData model.UserAuthBody) error {
	// CAN BE USED TO HANDLE LOGOUT LOGIC IF NEEDED IN THE FUTURE, EXP: MONITORING LOGGING WHEN LOGOUT
	return nil
}
