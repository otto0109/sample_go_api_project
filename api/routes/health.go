package routes

import (
	"github.com/gin-gonic/gin"
)

type EventController interface {
	Ping(context *gin.Context)
}

func (r router) addEventRoutes(rg *gin.RouterGroup, controller EventController) {
	eventRoutes := rg.Group("/health")

	eventRoutes.POST("/", controller.Ping)
}
