package controller

import (
	"github.com/snail007/gmc"
)

type Login struct {
	gmc.Controller
}

func (this *Login) Auth() {
	this.View.
		Set("title", "It works!").
		Render("welcome")
}

func (this *Login) Index() {
	this.View.Layout("login").Render("login/login")
}
