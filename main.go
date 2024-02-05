package main

import (
	_ "datingapp-api/routers"
	"datingapp-api/models"
	"datingapp-api/middlewares"

	// beego "github.com/beego/beego/v2/server/web"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:@/dating_app?charset=utf8")
	orm.RegisterModel(new(models.User))
	orm.RegisterModel(new(models.Profile))
	orm.RegisterModel(new(models.Feature))
	orm.RegisterModel(new(models.Subscription))
	orm.RegisterModel(new(models.UserActivity))
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type",
			"user", "email", "action", "page", "menu", "token", "message", "merchantPubId"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "token", "expires_in", "refresh_token", "role", "username"},
		AllowCredentials: true,
	}))
	beego.InsertFilter("/*", beego.BeforeRouter, middlewares.TokenAuthMiddleware)

	beego.Run()
}
