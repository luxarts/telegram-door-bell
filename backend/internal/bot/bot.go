package bot

import (
	tg "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"telegram-door-bell/internal/controller"
	"telegram-door-bell/internal/defines"
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
	// Services
	tokenSrv := service.NewTokenService()

	// Controllers
	telegramCtrl := controller.NewTelegramController(b, tokenSrv)

	// Handlers
	b.Handle(defines.CommandStart, telegramCtrl.Start)
	b.Handle(defines.CommandHelp, telegramCtrl.Help)
	b.Handle(defines.CommandToken, telegramCtrl.Token)
}
