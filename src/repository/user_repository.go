package repository

import (
	"momonga_blog/database"
	"momonga_blog/repository/model"
	"time"
)

type UserRepositoryInterface interface {
	FindUserByUserID(userID string) (*model.Users, error)
	FindUserByUuid(uuid string) (*model.Users, error)
	FindUserByRefreshToken(refresh_token string) (*model.Users, error)
	SaveRefreshToken(user *model.Users, refreshToken string, exp time.Time) (*model.Users, error)
	SaveLogout(user *model.Users) error
}

type UserRepository struct {
	model model.Users
}

var _ UserRepositoryInterface = &UserRepository{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) FindUserByUserID(userID string) (*model.Users, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}

	result := db.Where("user_id = ?", userID).First(&ur.model)
	if result.Error != nil {
		return nil, result.Error
	}
	return &ur.model, nil
}

func (ur *UserRepository) FindUserByUuid(uuid string) (*model.Users, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}

	result := db.Where("uuid = ?", uuid).First(&ur.model)
	if result.Error != nil {
		return nil, result.Error
	}
	return &ur.model, nil
}

func (ur *UserRepository) FindUserByRefreshToken(refresh_token string) (*model.Users, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}

	result := db.Where("refresh_token = ?", refresh_token).First(&ur.model)
	if result.Error != nil {
		return nil, result.Error
	}
	return &ur.model, nil
}

func (ur *UserRepository) SaveRefreshToken(user *model.Users, refreshToken string, exp time.Time) (*model.Users, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}

	user.Active = true
	user.RefreshToken = &refreshToken
	user.TokenExpiry = &exp
	result := db.Save(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (ur *UserRepository) SaveLogout(user *model.Users) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	}

	user.RefreshToken = nil
	user.TokenExpiry = nil
	user.Active = false
	result := db.Save(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}