package telbot

import (
	"regexp"

	"github.com/go-telegram/bot"
)

type Telbot struct {
	UserManager *UserManager
	Rules       []Rule
}

func New() *Telbot {
	return &Telbot{}
}

func (me *Telbot) SetUserManager(userManager *UserManager) {
	me.UserManager = userManager
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
