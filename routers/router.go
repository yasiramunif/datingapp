// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"datingapp-api/controllers"

	// beego "github.com/beego/beego/v2/server/web"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/auth",
			beego.NSInclude(
				&controllers.AuthController{},
			),
		),
		beego.NSNamespace("/features",
			beego.NSInclude(
				&controllers.FeatureController{},
			),
		),
		beego.NSNamespace("/subscriptions",
			beego.NSInclude(
				&controllers.SubscriptionController{},
			),
		),
		beego.NSNamespace("/activities",
			beego.NSInclude(
				&controllers.UserActivityController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
