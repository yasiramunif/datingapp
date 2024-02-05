package middlewares

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/juusechec/jwt-beego"
	"datingapp-api/models"
	"datingapp-api/services"
	"datingapp-api/repositories"
	"strings"
	"net/http"
)

// TokenAuthMiddleware middleware to validate access token
func TokenAuthMiddleware(c *context.Context) {
	currentRoute := c.Request.URL.Path
	logs.Info("Current Route %s", currentRoute)

	excludeTokenValidationRoutes := "/v1/auth/login,/v1/auth/register"
	excludedRoutes := strings.Split(excludeTokenValidationRoutes, ",")
	if shouldExcludeRoute(currentRoute, excludedRoutes) {
		logs.Info("shouldExcludeRoute %s", currentRoute)
		return
	}

	authorization := c.Input.Header("Authorization")
	if authorization == "" {
    logs.Warn("Missing authorization header")
		c.Output.SetStatus(http.StatusUnauthorized)
		c.Output.JSON(models.GetResponse(nil, http.StatusUnauthorized, "Missing authorization header"), true, true)
    return
  }
	tokenString := authorization[len("Bearer "):]
	valid, issData := validateToken(tokenString)
	logs.Info("iss %v", issData)
	if !valid {
		logs.Warn("token invalid")
		c.Output.SetStatus(http.StatusUnauthorized)
		c.Output.JSON(models.GetResponse(nil, http.StatusUnauthorized, "token invalid"), true, true)
		return
	}

	c.Input.SetData("currentUser", issData)
}

// shouldExcludeRoute function to validate exclude route
func shouldExcludeRoute(currentRoute string, excludedRoutes []string) bool {
	for _, route := range excludedRoutes {
		if currentRoute == route {
			return true
		}
	}
	return false
}

// validateToken function to validate access token
func validateToken(token string) (bool, interface{}) {
	et := jwtbeego.EasyToken{}
	valid, iss, _ := et.ValidateToken(token)

	o := orm.NewOrm()
	userRepo := repositories.NewUserRepo(o)
	service := services.NewUserService(userRepo)
	resp := service.GetUserByUsername(iss)
	
	return valid, resp.Content
}
