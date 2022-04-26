package controller

import (
	"fmt"
	tg "gopkg.in/tucnak/telebot.v2"
	"log"
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
	userSrv  service.UserService
}

func NewTelegramController(bot *tg.Bot, tokenSrv service.TokenService, userSrv service.UserService) TelegramController {
	return &telegramController{
		bot:      bot,
		tokenSrv: tokenSrv,
		userSrv:  userSrv,
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

	err := c.userSrv.Create(m.Sender.ID, token)
	if err != nil {
		log.Printf("Error: %+v\n", err)
		c.errorRespond(m.Sender)
		return
	}

	_, err = c.bot.Send(m.Sender, fmt.Sprintf(defines.MessageToken, token), tg.ModeMarkdown)
	if err != nil {
		log.Printf("Error: %+v\n", err)
		c.errorRespond(m.Sender)
		return
	}
}

func (c *telegramController) errorRespond(recipient tg.Recipient) {
	_, err := c.bot.Send(recipient, fmt.Sprintf(defines.MessageError), tg.ModeMarkdown)
	if err != nil {
		log.Printf("Error sending message: %+v\n", err)
	}
}
