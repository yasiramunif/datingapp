package services

import (
	"datingapp-api/models"
	"datingapp-api/repositories"
	"github.com/astaxie/beego/logs"
	"fmt"
)

// FeatureService struct
type FeatureService struct {
	featureRepo      repositories.IFeatureRepository
}

// NewFeatureService function to create FeatureService
func NewFeatureService(ur repositories.IFeatureRepository) *FeatureService {
	return &FeatureService{
		featureRepo: ur,
	}
}

// GetAll funct to get all features
func (svc *FeatureService) GetAll() models.Response {
	logs.Info("Start Service GetAll")

	features, err := svc.featureRepo.GetAllFeatures()
	if err != nil {
		message := "failed get all features"
		logs.Error(message)
		logs.Error(err)
		return models.GetResponse(nil, 503, message)
	}

	return models.GetResponse(features, 200, "success")
}

func (svc *FeatureService) GetFeatureById(id int64) models.Response {
	logs.Info("Start Service GetFeatureById")

	feature, err := svc.featureRepo.GetFeatureById(id)
	if err!= nil {
		message := fmt.Sprintf("failed get feature with id %d", id)
    logs.Error(message)
		logs.Error(err)
		return models.GetResponse(nil, 400, message)
	}

	return models.GetResponse(feature, 200, "success")
}