package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/gq-leon/sport-backend/bootstrap"
	"github.com/gq-leon/sport-backend/domain"
)

type UserController struct {
	Env         *bootstrap.Env
	UserUseCase domain.UserUseCase
}

func (uc *UserController) Signup(c *gin.Context) {
	var request domain.SignupRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		domain.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if _, err := uc.UserUseCase.GetUserByEmail(c, request.Email); err == nil {
		domain.ErrorResponse(c, http.StatusConflict, fmt.Errorf("user already exists with the given email"))
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		domain.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	request.Password = string(password)

	user := &domain.User{
		Name:      request.Name,
		Email:     request.Email,
		Password:  request.Password,
		CreatedAt: time.Now(),
	}

	if err = uc.UserUseCase.Create(c, user); err != nil {
		domain.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	tokenResponse, err := uc.getUserToken(user)
	if err != nil {
		domain.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	domain.SuccessResponse(c, tokenResponse)
}

func (uc *UserController) Login(c *gin.Context) {
	var request domain.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		domain.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	user, err := uc.UserUseCase.GetUserByEmail(c, request.Email)
	if err != nil {
		domain.ErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("user not found with the given email"))
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		domain.ErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("user not found with the given email"))
		return
	}

	tokenResponse, err := uc.getUserToken(&user)
	if err != nil {
		domain.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	domain.SuccessResponse(c, tokenResponse)
}

func (uc *UserController) getUserToken(user *domain.User) (domain.TokenResponse, error) {
	accessToken, err := uc.UserUseCase.CreateAccessToken(user, uc.Env.AccessTokenSecret, uc.Env.AccessTokenExpiryHour)
	if err != nil {
		return domain.TokenResponse{}, err
	}

	refreshToken, err := uc.UserUseCase.CreateRefreshToken(user, uc.Env.RefreshTokenSecret, uc.Env.RefreshTokenExpiryHour)
	if err != nil {
		return domain.TokenResponse{}, err
	}

	return domain.TokenResponse{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}
