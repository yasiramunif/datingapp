package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["datingapp-api/controllers:AuthController"] = append(beego.GlobalControllerRouter["datingapp-api/controllers:AuthController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["datingapp-api/controllers:AuthController"] = append(beego.GlobalControllerRouter["datingapp-api/controllers:AuthController"],
        beego.ControllerComments{
            Method: "Register",
            Router: "/register",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["datingapp-api/controllers:FeatureController"] = append(beego.GlobalControllerRouter["datingapp-api/controllers:FeatureController"],
        beego.ControllerComments{
            Method: "GetAllFeatures",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["datingapp-api/controllers:SubscriptionController"] = append(beego.GlobalControllerRouter["datingapp-api/controllers:SubscriptionController"],
        beego.ControllerComments{
            Method: "CreateSubscription",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["datingapp-api/controllers:UserActivityController"] = append(beego.GlobalControllerRouter["datingapp-api/controllers:UserActivityController"],
        beego.ControllerComments{
            Method: "CreateUserActivity",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["datingapp-api/controllers:UserActivityController"] = append(beego.GlobalControllerRouter["datingapp-api/controllers:UserActivityController"],
        beego.ControllerComments{
            Method: "ViewUser",
            Router: "/view",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["datingapp-api/controllers:UserController"] = append(beego.GlobalControllerRouter["datingapp-api/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
