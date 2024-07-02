package handlers

import (
	"strings"
	"telbot/usermanager"
)

type Handlers struct {
	UserManager *usermanager.UserManager
}

func New() Handlers {
	return Handlers{}
}

func (me *Handlers) SetUserManager(usermanagerInstace *usermanager.UserManager) {
	me.UserManager = usermanagerInstace
}

func (me *Handlers) ArgsGetter(text string) (int, []string) {
	args := strings.Split(text, " ")
	return len(args) - 1, args[1:]

}
