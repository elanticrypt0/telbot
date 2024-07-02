package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"telbot/app"
	"telbot/telbot"
	"telbot/usermanager"

	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
)

var userManager = usermanager.NewUserManager()

// Send any text message to the bot after the bot has been started

func main() {

	err := godotenv.Load() // 👈 load .env file
	if err != nil {
		log.Fatal(err)
	}

	// Load config
	config := app.NewConfig()

	userManager.AddUser(config.AdminUserId, "admin", "super", "admin")

	telbot := telbot.New()
	telbot.SetUserManager(userManager)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(telbot.Handlers.Default),
	}

	b, err := bot.New(config.TelegramApiKey, opts...)
	if err != nil {
		panic(err)
	}

	telbot.AddRule(bot.HandlerTypeMessageText, `^/start`, telbot.Handlers.MyStart)
	telbot.AddRule(bot.HandlerTypeMessageText, `^/args`, telbot.Handlers.GetArgs)
	telbot.AddRule(bot.HandlerTypeMessageText, `^/mydata`, telbot.Handlers.MyData)
	telbot.AddRule(bot.HandlerTypeMessageText, `^/letmein`, telbot.Handlers.LetMeIn)
	telbot.AddRule(bot.HandlerTypeMessageText, `^/ifthisis`, telbot.Handlers.IfThisIs)

	telbot.InitRules(b)

	botinfo, err := b.GetMe(ctx)

	if err != nil {
		fmt.Println(err)
	}

	app.AppBanner(botinfo)

	b.Start(ctx)
}
