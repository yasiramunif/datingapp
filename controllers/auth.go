package controllers

import (
  "datingapp-api/models"
	"datingapp-api/repositories"
  "datingapp-api/services"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
	"encoding/json"
)

// Operations about Authentication
type AuthController struct {
	beego.Controller
}

// @Title Register
// @Description register user
// @Param	body		body 	models.RegisterRequest	true		"The object content"
// @Success 200 {object} models.AuthResponse
// @router /register [post]
func (c *AuthController) Register() {
	request := new(models.RegisterRequest)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &request)
	resp := models.GetResponse(nil, 200, "")
	if err != nil {
		logs.Error("Error unmarshal from http request body to expected object request")
		resp = models.GetResponse(nil, 400, err.Error())
	} else {
		o := orm.NewOrm()
		userRepo := repositories.NewUserRepo(o)
		service := services.NewAuthService(userRepo)
		resp = service.Register(request)
	}
	
	c.Data["json"] = resp
	c.ServeJSON()
}

// @Title Login
// @Description user login
// @Param	body		body 	models.LoginRequest	true		"The object content"
// @Success 200 {object} models.AuthResponse
// @router /login [post]
func (c *AuthController) Login() {
	request := new(models.LoginRequest)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &request)
	resp := models.GetResponse(nil, 200, "")
	if err != nil {
		logs.Error("Error unmarshal from http request body to expected object request")
		resp = models.GetResponse(nil, 400, err.Error())
	} else {
		o := orm.NewOrm()
		userRepo := repositories.NewUserRepo(o)
		service := services.NewAuthService(userRepo)
		resp = service.Login(request)
	}
	
	c.Data["json"] = resp
	c.ServeJSON()
}