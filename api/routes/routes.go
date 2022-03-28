package routes

import (
	"github.com/gin-gonic/gin"
	"sample_go_api_project/api/controller"
)

type router struct {
	router *gin.Engine
}

func NewRouter() router {
	r := router{
		router: gin.Default(),
	}

	v1 := r.router.Group("/v1")

	r.addEventRoutes(v1, controller.NewHealthController())

	return r
}

func (r router) Run(addr ...string) error {

	return r.router.Run(addr...)
}
