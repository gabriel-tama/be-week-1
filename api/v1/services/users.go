// services/user_service_impl.go

package services

import (
	"errors"

	"github.com/gabriel-tama/be-week-1/models"
	"golang.org/x/crypto/bcrypt"
)




type UserService interface {
    Register(username, name, password string) (*models.User, error)
    Authenticate(username, password string) (*models.User, error)
}

type userServiceImpl struct {
    userModel *models.UserModel
}

func NewUserService(userModel *models.UserModel) UserService {
    return &userServiceImpl{userModel: userModel}
}

func (us *userServiceImpl) Register(username, name, password string) (*models.User, error) {
    user,err := us.userModel.Create(username, name, password)
    if err != nil {
        return nil, err
    }


    return user, nil
}

func (us *userServiceImpl) Authenticate(username, password string) (*models.User, error) {
    user, err := us.userModel.FindUserByUsername(username)
    if err != nil {
        return nil, err
    }
        // Compare passwords
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return nil, errors.New("invalid username/password")
    }


    return user, nil
}