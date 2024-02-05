package utils

import (
	"datingapp-api/models"
	"encoding/json"
)

func ConvertUserContext(user interface{}) (models.User, error) {
	b, err := json.Marshal(user)

	if err != nil {
		return models.User{}, err
	}

	var userInfo models.User

	err = json.Unmarshal(b, &userInfo)

	if err != nil {
		return models.User{}, err
	}

	return userInfo, nil
}