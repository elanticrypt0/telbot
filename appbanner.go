package main

import (
	"fmt"

	"github.com/go-telegram/bot/models"
)

func appBanner(botinfo *models.User) {

	botInfo := fmt.Sprintf("BotInfo : %#v", botinfo)

	fmt.Println("")
	fmt.Println(":: TELBOT ::")
	fmt.Println("::   ON   ::")
	fmt.Println(botInfo)
	fmt.Println("")
}
