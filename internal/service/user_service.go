package service

import (
	"context"
	"errors"
	"log"

	"github.com/lipaysamart/go-jwt-exerices/internal/model"
	"github.com/lipaysamart/go-jwt-exerices/internal/repository"
	"github.com/lipaysamart/go-jwt-exerices/pkg/jtoken"
	"github.com/lipaysamart/go-jwt-exerices/pkg/utils"
	"github.com/lipaysamart/go-jwt-exerices/pkg/validation"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	Register(ctx context.Context, req *model.UserRegisterReq) error
	Login(ctx context.Context, req *model.UserLoginReq) (*model.User, string, string, error)
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	RefreshToken(ctx context.Context, userID string) (string, error)
}

type UserService struct {
	userRepo repository.IUserRepo
	validate validation.IValidation
}

func NewUserService(repo repository.IUserRepo, val validation.IValidation) *UserService {
	return &UserService{
		userRepo: repo,
		validate: val,
	}
}

func (s *UserService) Register(ctx context.Context, req *model.UserRegisterReq) error {
	if err := s.validate.ValidateStruct(req); err != nil {
		log.Println("failed to validate request body")
		return err
	}

	var user model.User

	utils.Copy(&user, req)

	if err := s.userRepo.Create(ctx, &user); err != nil {
		return err
	}

	return nil
}

func (s *UserService) Login(ctx context.Context, req *model.UserLoginReq) (*model.User, string, string, error) {
	if err := s.validate.ValidateStruct(req); err != nil {
		log.Println("failed to validate request body")
		return nil, "", "", err
	}

	user, err := s.userRepo.FindUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, "", "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, "", "", errors.New("wrong password")
	}

	tokenData := map[string]interface{}{
		"id": user.ID,
	}

	accessToken := jtoken.GenerateAccessToken(tokenData)
	refreshToken := jtoken.RefreshToken(tokenData)

	return user, accessToken, refreshToken, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	user, err := s.userRepo.FindUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) RefreshToken(ctx context.Context, userID string) (string, error) {
	user, err := s.userRepo.FindUserByID(ctx, userID)
	if err != nil {
		return "", err
	}

	tokenData := map[string]interface{}{
		"id": user.ID,
	}
	accessToken := jtoken.GenerateAccessToken(tokenData)
	return accessToken, nil
}
