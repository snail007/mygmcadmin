package controller

import (
	"github.com/snail007/gmc"
)

type Admin struct {
	Base
	Admin gmc.Mss
}

func (this *Admin) Before__() {
	err := this.SessionStart()
	if err != nil {
		this.Stop(err)
	}
	if u, ok := this.IsLogin__(); !ok {
		this.Ctx.Redirect("/")
		return
	} else {
		this.Admin = u
	}
	this.View.Set("admin",this.Admin)
}

func (this *Admin) IsLogin__() (user gmc.Mss, isLogin bool) {
	u := this.Session.Get("admin")
	if u != nil && u.(gmc.Mss)["username"] != "" {
		return u.(gmc.Mss), true
	}
	return nil, false
}
