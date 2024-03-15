// services/user_service_impl.go

package services

import (
	"github.com/gabriel-tama/be-week-1/models"
)

type UserService interface {
	Register(username, name, password string) (*models.User, error)
	CheckUserExists(username string) (bool, error)
	GetUserExist(username string) (*models.User, error)
}

type userServiceImpl struct {
	userModel *models.UserModel
}

func NewUserService(userModel *models.UserModel) UserService {
	return &userServiceImpl{userModel: userModel}
}

func (us *userServiceImpl) Register(username, name, password string) (*models.User, error) {
	user, err := us.userModel.Create(username, name, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *userServiceImpl) CheckUserExists(username string) (bool, error) {
	exist, err := us.userModel.FindExistByUsername(username)
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (us *userServiceImpl) GetUserExist(username string) (*models.User, error) {
	user, err := us.userModel.FindUserByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}
