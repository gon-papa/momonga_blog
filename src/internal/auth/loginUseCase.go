package auth

import (
	"context"
	"fmt"
	"momonga_blog/repository"
)

type LoginUseCaseInterface interface {
	Login(ctx context.Context, userId string, password string) (*Token, error)
	Logout(ctx context.Context) (bool, error)
}

type loginUseCase struct {
	repository repository.UserRepositoryInterface
}

var _ LoginUseCaseInterface = &loginUseCase{}

func NewLoginUseCase() LoginUseCaseInterface {
	return &loginUseCase{
		repository: repository.NewUserRepository(),
	}
}

type Token struct {
	Token string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func (luc *loginUseCase) Login(ctx context.Context, userId string, password string) (*Token, error) {
	user, err := luc.repository.FindUserByUserID(userId)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	err = ComparePassword(user.Password, password)
	if err != nil {
		return nil, fmt.Errorf("password is incorrect: %w", err)
	}

	// トークンとリフレッシュトークンを返す
	token, err := CreateAccessToken(user.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to create access token: %w", err)
	}
	refreshToken, err := CreateRefreshToken(64)
	if err != nil {
		return nil, fmt.Errorf("failed to create refresh token: %w", err)
	}
	expired := CreateRefreshTokenExpire(7)
	_, err = luc.repository.SaveRefreshToken(user, refreshToken, expired)
	if err != nil {
		return nil, fmt.Errorf("failed to save refresh token: %w", err)
	}

	return &Token{
		Token: token,
		RefreshToken: refreshToken,
	}, nil
}

func (luc *loginUseCase) Logout(ctx context.Context) (bool, error) {
	return true, nil
}