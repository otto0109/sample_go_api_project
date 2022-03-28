package routes

import (
	"event_service/api/controller"
	"github.com/gin-gonic/gin"
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
