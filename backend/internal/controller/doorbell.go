package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"strings"
	"telegram-door-bell/internal/defines"
	"telegram-door-bell/internal/service"
	"telegram-door-bell/internal/utils/jwt"
)

type DoorBellController interface {
	Ring(ctx *gin.Context)
}

type doorBellController struct {
	doorBellSrv service.DoorBellService
	userSrv     service.UserService
}

func NewDoorBellController(doorBellSrv service.DoorBellService, userSrv service.UserService) DoorBellController {
	return &doorBellController{
		doorBellSrv: doorBellSrv,
		userSrv:     userSrv,
	}
}

func (c *doorBellController) Ring(ctx *gin.Context) {
	authorization := ctx.GetHeader("Authorization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	bearerTokenSplit := strings.Split(authorization, " ")

	if len(bearerTokenSplit) != 2 || bearerTokenSplit[0] != "Bearer" {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if !jwt.Verify(bearerTokenSplit[1], os.Getenv(defines.EnvTokenSecret)) {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	payload, err := jwt.GetPayload(bearerTokenSplit[1])
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	chatID, err := strconv.ParseInt(payload.Subject, 10, 64)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	user, err := c.userSrv.Read(chatID)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if user.Token != bearerTokenSplit[1] {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.doorBellSrv.SendMessage(chatID)

	ctx.String(http.StatusOK, "Message sent.")
}
