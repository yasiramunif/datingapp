package services

import (
	"datingapp-api/models"
	"datingapp-api/repositories"
	"github.com/astaxie/beego/logs"
	"fmt"
	"time"
)

// SubscriptionService struct
type SubscriptionService struct {
	subscriptionRepo      repositories.ISubscriptionRepository
}

// NewSubscriptionService function to create SubscriptionService
func NewSubscriptionService(ur repositories.ISubscriptionRepository) *SubscriptionService {
	return &SubscriptionService{
		subscriptionRepo: ur,
	}
}

// CreateSubcription funct to create subscription
func (svc *SubscriptionService) CreateSubcription(request *models.SubscriptionRequest, userId int64) models.Response {
	logs.Info("Start Service CreateSubcription")

	var subscription models.Subscription
	subscription.User = &models.User{Id: userId}
	subscription.Feature = &models.Feature{Id:request.FeatureId}
	subscription.PurchasedAt = time.Now().Format("2006-01-02 15:04:05")

	subscriptions, err := svc.subscriptionRepo.CreateSubscription(subscription)
	if err != nil {
		message := "failed get all subscriptions"
		logs.Error(message)
		logs.Error(err)
		return models.GetResponse(nil, 503, message)
	}

	return models.GetResponse(subscriptions, 200, "success")
}

func (svc *SubscriptionService) GetSubscriptionById(id int64) models.Response {
	logs.Info("Start Service GetSubscriptionById")

	subscription, err := svc.subscriptionRepo.GetSubscriptionById(id)
	if err!= nil {
		message := fmt.Sprintf("failed get subscription with id %d", id)
    logs.Error(message)
		logs.Error(err)
		return models.GetResponse(nil, 400, message)
	}

	return models.GetResponse(subscription, 200, "success")
}