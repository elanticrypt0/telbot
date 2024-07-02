package telbot

import (
	"regexp"
	"telbot/handlers"
	"telbot/usermanager"

	"github.com/go-telegram/bot"
)

type Telbot struct {
	UserManager *usermanager.UserManager
	Handlers    handlers.Handlers
	Rules       []Rule
}

func New() *Telbot {
	return &Telbot{
		Handlers: handlers.New(),
	}
}

func (me *Telbot) SetUserManager(userManager *usermanager.UserManager) {
	me.UserManager = userManager
}

func (me *Telbot) SetHandlers(handlerIntance handlers.Handlers) {
	me.Handlers = handlerIntance
}

func (me *Telbot) AddRule(handlerType bot.HandlerType, regex string, handlerFunc bot.HandlerFunc) {

	rule := Rule{
		HandlerType: handlerType,
		Regex:       regex,
		HandlerFunc: handlerFunc,
	}

	me.Rules = append(me.Rules, rule)

}

func (me *Telbot) InitRules(b *bot.Bot) {

	for _, rule := range me.Rules {
		b.RegisterHandlerRegexp(rule.HandlerType, regexp.MustCompile(rule.Regex), rule.HandlerFunc)
	}

}
