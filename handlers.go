package main

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {

	fmt.Printf("%#v -> %s (%s %s) \n", update.Message.From.ID, update.Message.From.Username, update.Message.From.FirstName, update.Message.From.LastName)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Whaaaaaaaaaaat?",
	})
}

func myStartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	userInfo := update.Message.From
	response := fmt.Sprintf("Hi %s , %s", userInfo.Username, update.Message.Text)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   response,
	})
}

func myDataHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	userInfo := update.Message.From
	userContact := update.Message.Contact
	response := fmt.Sprintf("Your info \n %#v \n %#v", userInfo, userContact)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   response,
	})
}

func letMeInHandler(ctx context.Context, b *bot.Bot, update *models.Update) {

	if userManager.IsAllowedUser(update.Message.From.ID) {
		userInfo := update.Message.From
		response := fmt.Sprintf("%s , you are welcome \n", userInfo.FirstName)

		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   response,
		})
	} else {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "You shall not pass",
		})
	}
}
