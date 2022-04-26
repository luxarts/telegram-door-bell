package service

import tg "gopkg.in/tucnak/telebot.v2"

type DoorBellService interface {
	SendMessage(chatID int64)
}

type doorBellService struct {
	bot *tg.Bot
}

func NewDoorBellService(bot *tg.Bot) DoorBellService {
	return &doorBellService{bot: bot}
}

func (s *doorBellService) SendMessage(chatID int64) {
	s.bot.Send(tg.ChatID(chatID), "Timbre!", tg.ModeMarkdown)
}
