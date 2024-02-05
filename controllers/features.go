package controllers

import (
	// "datingapp-api/models"
	"datingapp-api/repositories"
	"datingapp-api/services"
	"datingapp-api/utils"
	// "encoding/json"

	// beego "github.com/beego/beego/v2/server/web"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
)

// Operations about Feature
type FeatureController struct {
	beego.Controller
}

// @Title GetAllFeatures
// @Description get all features
// @Success 200 {object} models.Feature
// @router / [get]
func (u *FeatureController) GetAllFeatures() {
	userContext := u.Ctx.Input.GetData("currentUser")
	currentUser, _ := utils.ConvertUserContext(userContext)
	logs.Info("Current user %v", currentUser)
	o := orm.NewOrm()
	featureRepo := repositories.NewFeatureRepo(o)
	service := services.NewFeatureService(featureRepo)
	resp := service.GetAll()
	u.Data["json"] = resp
	u.ServeJSON()
}
