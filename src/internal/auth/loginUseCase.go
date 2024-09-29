package auth

import (
	"context"
	"fmt"
	"momonga_blog/api"
	"momonga_blog/repository"
	"time"
)

type LoginUseCaseInterface interface {
	Login(ctx context.Context, userId string, password string) (*Token, error)
	Logout(ctx context.Context) error
	RefreshToken(ctx context.Context, refresh_token string) (*Token, error)
	HandleBearerAuth(ctx context.Context, operationName string, t api.BearerAuth) (context.Context, error)
}

type loginUseCase struct {
	repository repository.UserRepositoryInterface
}

var _ LoginUseCaseInterface = &loginUseCase{}
var _ api.SecurityHandler = &loginUseCase{}

func NewLoginUseCase() LoginUseCaseInterface {
	return &loginUseCase{
		repository: repository.NewUserRepository(),
	}
}

type Token struct {
	Token string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type contextKey string
const AuthUuid contextKey = "auth_uuid"

func (luc *loginUseCase) HandleBearerAuth(ctx context.Context, operationName string, t api.BearerAuth) (context.Context, error) {
	uuid, err := AuthAccessToken(t.Token)
	if err != nil {
		return ctx, fmt.Errorf("failed to auth access token: %w", err)
	}

	user, err := luc.repository.FindUserByUuid(uuid)
	if err != nil {
		return ctx, fmt.Errorf("failed to find user by uuid: %w", err)
	}

	if user == nil {
		return ctx, fmt.Errorf("user not found")
	}

	return context.WithValue(ctx, AuthUuid, uuid), nil
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
	token, err := CreateAccessToken(user.UUID)
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

func (luc *loginUseCase) Logout(ctx context.Context) error {
	uuid := ctx.Value(AuthUuid)
	if uuid == nil {
		return fmt.Errorf("login user not found")
	}

	user, err := luc.repository.FindUserByUuid(uuid.(string))
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	if !user.Active {
		return fmt.Errorf("user is not active")
	}

	err = luc.repository.SaveLogout(user)
	if err != nil {
		return fmt.Errorf("failed to save logout: %w", err)
	}
	return nil
}

func (luc *loginUseCase) RefreshToken(ctx context.Context, refresh_token string) (*Token, error) {
	user, err := luc.repository.FindUserByRefreshToken(refresh_token)
	if err != nil {
		return nil, fmt.Errorf("failed to find user by refresh token: %w", err)
	}

	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	if !user.Active {
		return nil, fmt.Errorf("user is not active")
	}

	// リフレッシュトークンの有効期限をチェック
	if user.TokenExpiry.Before(time.Now()) {
		return nil, fmt.Errorf("refresh token is expired")
	}

	// トークンとリフレッシュトークンを返す
	token, err := CreateAccessToken(user.UUID)
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