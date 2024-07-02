package handlers

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (me *Handlers) Default(ctx context.Context, b *bot.Bot, update *models.Update) {

	fmt.Printf("%#v -> %s (%s %s) \n", update.Message.From.ID, update.Message.From.Username, update.Message.From.FirstName, update.Message.From.LastName)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Whaaaaaaaaaaat?",
	})
}

func (me *Handlers) MyStart(ctx context.Context, b *bot.Bot, update *models.Update) {
	userInfo := update.Message.From
	response := fmt.Sprintf("Hi %s , %s", userInfo.Username, update.Message.Text)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   response,
	})
}

func (me *Handlers) MyData(ctx context.Context, b *bot.Bot, update *models.Update) {
	userInfo := update.Message.From
	userContact := update.Message.Contact
	response := fmt.Sprintf("Your info \n %#v \n %#v", userInfo, userContact)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   response,
	})
}

func (me *Handlers) LetMeIn(ctx context.Context, b *bot.Bot, update *models.Update) {

	if me.UserManager.IsAllowedUser(update.Message.From.ID) {
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

func (me *Handlers) IfThisIs(ctx context.Context, b *bot.Bot, update *models.Update) {

	num_str := strings.Split(update.Message.Text, " ")
	fmt.Printf("%#v", len(num_str))
	num, err := strconv.Atoi(num_str[1])
	if err != nil {
		response := "The argument must be a number"
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   response,
		})
	}
	response := fmt.Sprintf("This %d is %d", num, num+2)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   response,
	})
}

func (me *Handlers) GetArgs(ctx context.Context, b *bot.Bot, update *models.Update) {

	argsQty, args := me.ArgsGetter(update.Message.Text)

	response := fmt.Sprintf("Args (%d): %#v", argsQty, args)
	fmt.Println(response)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   response,
	})
}
