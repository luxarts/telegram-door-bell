package router

import (
	"github.com/gin-gonic/gin"
	tg "gopkg.in/tucnak/telebot.v2"
	"telegram-door-bell/internal/controller"
	"telegram-door-bell/internal/defines"
	"telegram-door-bell/internal/service"
)

func New(b *tg.Bot) *gin.Engine {
	r := gin.Default()

	mapRoutes(r, b)

	return r
}

func mapRoutes(r *gin.Engine, b *tg.Bot) {
	// Service
	doorBellSrv := service.NewDoorBellService(b)

	// Controllers
	doorBellCtrl := controller.NewDoorBellController(doorBellSrv)

	r.POST(defines.EndpointRing, doorBellCtrl.Ring)
}
