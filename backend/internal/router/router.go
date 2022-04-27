package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	tg "gopkg.in/tucnak/telebot.v2"
	"os"
	"telegram-door-bell/internal/controller"
	"telegram-door-bell/internal/defines"
	"telegram-door-bell/internal/repository"
	"telegram-door-bell/internal/service"
)

func New(b *tg.Bot) *gin.Engine {
	r := gin.Default()

	mapRoutes(r, b)

	return r
}

func mapRoutes(r *gin.Engine, b *tg.Bot) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv(defines.EnvRedisURL),
		Username: os.Getenv(defines.EnvRedisUsername),
		Password: os.Getenv(defines.EnvRedisPassword),
	})

	// Repositories
	userRepo := repository.NewUsersRepository(redisClient)

	// Services
	doorBellSrv := service.NewDoorBellService(b)
	userSrv := service.NewUserService(userRepo)

	// Controllers
	doorBellCtrl := controller.NewDoorBellController(doorBellSrv, userSrv)

	r.POST(defines.EndpointRing, doorBellCtrl.Ring)
}
