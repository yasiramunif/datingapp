package services

import (
	"datingapp-api/models"
	"datingapp-api/repositories"
	"github.com/astaxie/beego/logs"
)

// UserActivityService struct
type UserActivityService struct {
	useractivityRepo      repositories.IUserActivityRepository
	subscriptionRepo      repositories.ISubscriptionRepository
}

// NewUserActivityService function to create UserActivityService
func NewUserActivityService(ur repositories.IUserActivityRepository, sr repositories.ISubscriptionRepository) *UserActivityService {
	return &UserActivityService{
		useractivityRepo: ur,
		subscriptionRepo: sr,
	}
}

// ViewUser funct to view user
func (svc *UserActivityService) ViewUser(userId int64) models.Response {
	logs.Info("Start Service ViewUser")

	subscription, err := svc.subscriptionRepo.GetSubscriptionByUserId(userId)
	if err != nil {
		message := "failed get subscription by user"
		logs.Error(message)
		logs.Error(err)
		return models.GetResponse(nil, 503, message)
	}

	if (subscription.Feature.Name != "Unlimited Swipe") {
		countView, err := svc.useractivityRepo.CountViewUserInADay(userId)
		if err != nil {
			message := "failed get count view user in a day"
			logs.Error(message)
			logs.Error(err)
			return models.GetResponse(nil, 503, message)
		}

		if countView > 9 {
			return models.GetResponse(nil, 402, "You reached the limit")
		}
	}

	otherUser, err := svc.useractivityRepo.ViewUser(userId)
	if err != nil {
		message := "failed view user"
		logs.Error(message)
		logs.Error(err)
		return models.GetResponse(nil, 503, message)
	}

	return models.GetResponse(otherUser, 200, "success")
}

// CreateUserActivityService struct
type CreateUserActivityService struct {
	useractivityRepo      repositories.IUserActivityRepository
}

// NewCreateUserActivityService function to create CreateUserActivityService
func NewCreateUserActivityService(ur repositories.IUserActivityRepository) *CreateUserActivityService {
	return &CreateUserActivityService{
		useractivityRepo: ur,
	}
}

func (svc *CreateUserActivityService) CreateUserActivity(request *models.UserActivityRequest, userId int64) models.Response {
	logs.Info("Start Service CreateUserActivity")

	var userActivity models.UserActivity
	userActivity.OtherUser = &models.User{Id: request.OtherUserId}
	userActivity.User = &models.User{Id: userId}
	userActivity.ActivityType = request.ActivityType
	_, err := svc.useractivityRepo.CreateUserActivity(userActivity)
	if err != nil {
		message := "failed create user activity"
		logs.Error(message)
		logs.Error(err)
		return models.GetResponse(nil, 503, message)
	}

	return models.GetResponse(nil, 200, "success")
}