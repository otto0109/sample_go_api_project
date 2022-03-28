package controller

import "github.com/gin-gonic/gin"

type healthController struct {
}

func NewHealthController() *healthController {
	return &healthController{}
}

func (h healthController) Ping(c *gin.Context) {
	c.JSON(200, "Pong")
}
