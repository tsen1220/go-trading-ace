package controllers

import (
	"go-uniswap/config"

	"github.com/gin-gonic/gin"
)

type IHomeController interface {
	SetRoutes(r *gin.Engine)
	Home(ctx *gin.Context)
}

type HomeController struct {
	Config *config.Config
}

func NewHomeController(config *config.Config) IHomeController {
	return &HomeController{
		Config: config,
	}
}

func (h *HomeController) SetRoutes(r *gin.Engine) {
	r.GET("/", h.Home)
}

func (h *HomeController) Home(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"hello": "world"})
}
