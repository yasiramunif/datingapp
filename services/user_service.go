package services

import (
	"datingapp-api/models"
	"datingapp-api/repositories"
	"github.com/astaxie/beego/logs"
	"fmt"
)

// UserService struct
type UserService struct {
	userRepo      repositories.IUserRepository
}

// NewUserService function to create UserService
func NewUserService(ur repositories.IUserRepository) *UserService {
	return &UserService{
		userRepo: ur,
	}
}

// GetAll funct to get all users
func (svc *UserService) GetAll() models.Response {
	logs.Info("Start Service GetAll")

	users, err := svc.userRepo.GetAllUsers()
	if err != nil {
		message := "failed get all users"
		logs.Error(message)
		logs.Error(err)
		return models.GetResponse(nil, 503, message)
	}

	return models.GetResponse(users, 200, "success")
}

func (svc *UserService) GetUserByUsername(username string) models.Response {
	logs.Info("Start Service GetUserByUsername")

	user, err := svc.userRepo.GetUserByUsername(username)
	if err!= nil {
		message := fmt.Sprintf("failed get user with username %s", username)
    logs.Error(message)
		logs.Error(err)
		return models.GetResponse(nil, 400, message)
	}

	return models.GetResponse(user, 200, "success")
}