package bot

import (
	"github.com/go-redis/redis/v8"
	tg "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"telegram-door-bell/internal/controller"
	"telegram-door-bell/internal/defines"
	"telegram-door-bell/internal/repository"
	"telegram-door-bell/internal/service"
	"time"
)

func New() *tg.Bot {
	var err error

	b, err := tg.NewBot(tg.Settings{
		Token: os.Getenv(defines.EnvTelegramToken),
		Poller: &tg.LongPoller{
			Timeout: 30 * time.Second,
		},
		Verbose: true,
	})
	if err != nil {
		log.Fatalln(err)
	}

	mapCommands(b)

	return b
}

func mapCommands(b *tg.Bot) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv(defines.EnvRedisURL),
		Password: os.Getenv(defines.EnvRedisPassword),
	})

	// Repositories
	userRepo := repository.NewUsersRepository(redisClient)

	// Services
	tokenSrv := service.NewTokenService()
	userSrv := service.NewUserService(userRepo)

	// Controllers
	telegramCtrl := controller.NewTelegramController(b, tokenSrv, userSrv)

	// Handlers
	b.Handle(defines.CommandStart, telegramCtrl.Start)
	b.Handle(defines.CommandHelp, telegramCtrl.Help)
	b.Handle(defines.CommandToken, telegramCtrl.Token)
}
