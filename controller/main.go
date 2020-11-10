package controller

import (
	"github.com/snail007/gmc"
	"mygmcadmin/model"
)

type Main struct {
	gmc.Controller
}

func (this *Main) Index() {
	u,err:=model.User.GetByID("1")
	this.StopE(err, func() {
		this.Write(err)
	})
	this.View.Set("user",u)
	this.View.Layout("index").Render("main/index")
}

func (this *Main) Main() {
	this.View.Layout("page").Render("main/main")
}