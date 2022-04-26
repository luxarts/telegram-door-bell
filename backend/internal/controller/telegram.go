package controller

import (
	"fmt"
	tg "gopkg.in/tucnak/telebot.v2"
	"os"
	"telegram-door-bell/internal/defines"
	"telegram-door-bell/internal/service"
	"time"
)

type TelegramController interface {
	Start(m *tg.Message)
	Help(m *tg.Message)
	Token(m *tg.Message)
}

type telegramController struct {
	bot      *tg.Bot
	tokenSrv service.TokenService
}

func NewTelegramController(bot *tg.Bot, tokenSrv service.TokenService) TelegramController {
	return &telegramController{
		bot:      bot,
		tokenSrv: tokenSrv,
	}
}

func (c *telegramController) Start(m *tg.Message) {
	c.bot.Send(m.Sender, fmt.Sprintf(defines.MessageStart, os.Getenv(defines.EnvBackendURL)), tg.ModeMarkdown)
}
func (c *telegramController) Help(m *tg.Message) {
	c.bot.Send(m.Sender, fmt.Sprintf(defines.MessageHelp, os.Getenv(defines.EnvBackendURL)), tg.ModeMarkdown)
}
func (c *telegramController) Token(m *tg.Message) {
	token := c.tokenSrv.Generate(m.Sender.ID, time.Now())
	c.bot.Send(m.Sender, fmt.Sprintf(defines.MessageToken, token), tg.ModeMarkdown)
}
