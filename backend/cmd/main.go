package main

import (
	_ "github.com/joho/godotenv/autoload"
	"log"
	"telegram-door-bell/internal/bot"
	"telegram-door-bell/internal/router"
)

func main() {
	// Telegram bot
	b := bot.New()
	go b.Start()

	// REST server
	r := router.New(b)
	if err := r.Run(); err != nil {
		log.Fatalln(err)
	}
}
