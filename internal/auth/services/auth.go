package services

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/chitano/chatapp/internal/auth"
	"github.com/chitano/chatapp/internal/user/model"
	"github.com/chitano/chatapp/internal/user/repositories"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = os.Getenv("APP_KEY")

type AuthService interface {
	Login(req *auth.LoginRequest) (*auth.LoginResponse, error)
	Register(req *auth.RegisterRequest) error
}

type authService struct {
	repo repositories.UserRepository
}

// Login implements AuthService.
func (a *authService) Login(req *auth.LoginRequest) (*auth.LoginResponse, error) {

	u, err := a.repo.GetUserByEmail(req.Email)
	if err != nil {
		return &auth.LoginResponse{}, err
	}
	// return &auth.LoginResponse{AuthUser: &auth.AuthUser{ID: u.ID}}, err

	err = auth.CheckPassword(req.Password, u.Password)

	if err != nil {
		return &auth.LoginResponse{}, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, auth.AppJWTClaims{
		ID:    strconv.Itoa(int(u.ID)),
		Email: u.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(u.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	at, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return &auth.LoginResponse{}, err
	}

	authUser := auth.AuthUser{
		ID:       u.ID,
		Email:    u.Email,
		Username: u.Username,
	}
	loginResponse := auth.LoginResponse{
		AccessToken: at,
		ExpiresAt:   time.Now().Add(24 * time.Hour).Unix(),
		AuthUser:    &authUser,
	}

	return &loginResponse, nil
}

// Register implements AuthService.
func (a *authService) Register(req *auth.RegisterRequest) error {
	existingUser, err := a.repo.GetUserByEmail(req.Email)

	if err != nil {
		return err
	}

	if existingUser.ID != 0 {
		return errors.New("the email is already exists")
	}

	hashedPassword, err := auth.HashPassword(req.Password)

	if err != nil {
		return err
	}

	user := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Username: req.Username,
	}

	//Create user if email does not exist
	if err := a.repo.CreateUser(&user); err != nil {
		return err
	}

	return nil
}

func NewAuthService(repo repositories.UserRepository) AuthService {
	return &authService{repo: repo}
}
