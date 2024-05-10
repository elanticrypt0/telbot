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

	response := fmt.Sprintf("Hi %s , %s", update.Message.From.FirstName, update.Message.Text)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   response,
	})
}

func letMeInHandler(ctx context.Context, b *bot.Bot, update *models.Update) {

	if userManager.IsAllowedUser(update.Message.From.ID) {
		response := fmt.Sprintf("%s , %s", userManager.GetUser().Name, update.Message.Text)

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
