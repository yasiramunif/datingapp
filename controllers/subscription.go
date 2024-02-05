package controllers

import (
	"datingapp-api/models"
	"datingapp-api/repositories"
	"datingapp-api/services"
	"datingapp-api/utils"
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
)

// Operations about Subscription
type SubscriptionController struct {
	beego.Controller
}

// @Title CreateSubscription
// @Description create subscription
// @Param	body		body 	models.SubscriptionRequest	true		"The object content"
// @Success 200 {object} models.GetResponse
// @router / [post]
func (u *SubscriptionController) CreateSubscription() {
	userContext := u.Ctx.Input.GetData("currentUser")
	currentUser, _ := utils.ConvertUserContext(userContext)
	logs.Info("Current user %v", currentUser)
	request := new(models.SubscriptionRequest)
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &request)
	resp := models.GetResponse(nil, 200, "")
	if err != nil {
		logs.Error("Error unmarshal from http request body to expected object request")
		resp = models.GetResponse(nil, 400, err.Error())
	} else {
		logs.Info("Current user id %d", currentUser.Id)
		o := orm.NewOrm()
		subscriptionRepo := repositories.NewSubscriptionRepo(o)
		service := services.NewSubscriptionService(subscriptionRepo)
		resp = service.CreateSubcription(request, currentUser.Id)
	}
	
	u.Data["json"] = resp
	u.ServeJSON()
}
