package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/redshiftkeying/wd6/server/controller"
)

// gin router instance
var Router *gin.Engine

// Setting default router
func init() {
	Router = gin.Default()
}

// TODO: viper get env
func SetRouters(env string) {
	switch env {
	case "development":
		SetStaticRouter()
		SetRestRouter()
		SetGraphqlRouter()
	case "production":
		SetStaticRouter()
		SetRestRouter()
		SetGraphqlRouter()
	case "test":
		SetRestRouter()
		SetGraphqlRouter()
	default:
		SetStaticRouter()
		SetRestRouter()
		SetGraphqlRouter()
	}
}

func SetGraphqlRouter() {
	Router.POST("/graphql", controller.GraphqlHandler())
}

// reat api path user auth or more
func SetRestRouter() {
	Router.POST("/login", controller.AuthenticationHandler())
}

// Static Client
func SetStaticRouter() {
	Router.Static("/", "./dist")
}
