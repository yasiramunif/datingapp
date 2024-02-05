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

// Operations about UserActivity
type UserActivityController struct {
	beego.Controller
}

// @Title ViewUser
// @Description view user
// @Success 200 {object} models.Profile
// @router /view [get]
func (u *UserActivityController) ViewUser() {
	// useractivitys := models.GetAllUserActivitys()
	useractivityContext := u.Ctx.Input.GetData("currentUser")
	currentUser, _ := utils.ConvertUserContext(useractivityContext)
	logs.Info("Current useractivity %v", currentUser)
	o := orm.NewOrm()
	useractivityRepo := repositories.NewUserActivityRepo(o)
	subscriptionRepo := repositories.NewSubscriptionRepo(o)
	service := services.NewUserActivityService(useractivityRepo, subscriptionRepo)
	resp := service.ViewUser(currentUser.Id)
	u.Data["json"] = resp
	u.ServeJSON()
}

// @Title CreateUserActivity
// @Description create user activity
// @Param	body		body 	models.UserActivityRequest	true		"The object content"
// @Success 200 {object} models.Profile
// @router / [post]
func (u *UserActivityController) CreateUserActivity() {
	// useractivitys := models.GetAllUserActivitys()
	useractivityContext := u.Ctx.Input.GetData("currentUser")
	currentUser, _ := utils.ConvertUserContext(useractivityContext)
	logs.Info("Current useractivity %v", currentUser)
	request := new(models.UserActivityRequest)
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &request)
	resp := models.GetResponse(nil, 200, "")
	if err != nil {
		logs.Error("Error unmarshal from http request body to expected object request")
		resp = models.GetResponse(nil, 400, err.Error())
	} else {
		o := orm.NewOrm()
		useractivityRepo := repositories.NewUserActivityRepo(o)
		service := services.NewCreateUserActivityService(useractivityRepo)
		resp = service.CreateUserActivity(request, currentUser.Id)
	}
	
	u.Data["json"] = resp
	u.ServeJSON()
}