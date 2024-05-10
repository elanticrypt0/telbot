package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"telbot/telbot"

	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
)

var userManager = telbot.NewUserManager()

// Send any text message to the bot after the bot has been started

func main() {

	err := godotenv.Load() // 👈 load .env file
	if err != nil {
		log.Fatal(err)
	}

	telegram_api_key := os.Getenv("TELEGRAM_API_KEY")

	// sets admins telegram user id
	admin_user_id_str := os.Getenv("ADMIN_USER_ID")
	auID, err := strconv.Atoi(admin_user_id_str)
	if err != nil {
		log.Panic(err)
	}
	admin_user_id := int64(auID)
	// sets as admin (just to use some functions inside)
	userManager.AddUser(admin_user_id, "admin", "super", "admin")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(defaultHandler),
	}

	b, err := bot.New(telegram_api_key, opts...)
	if err != nil {
		panic(err)
	}

	telbot := telbot.New()
	telbot.SetUserManager(userManager)
	telbot.AddRule(bot.HandlerTypeMessageText, `^/start`, myStartHandler)
	telbot.AddRule(bot.HandlerTypeMessageText, `^/mydata`, myDataHandler)
	telbot.AddRule(bot.HandlerTypeMessageText, `^/letmein`, letMeInHandler)

	telbot.InitRules(b)

	botinfo, err := b.GetMe(ctx)

	if err != nil {
		fmt.Println(err)
	}

	appBanner(botinfo)

	b.Start(ctx)
}
