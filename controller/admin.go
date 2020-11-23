package controller

import (
	"github.com/snail007/gmc"
)

type Admin struct {
	Base
	User gmc.Mss
}

func (this *Admin) Before() {
	err := this.SessionStart()
	if err != nil {
		this.Stop(err)
	}
	if u, ok := this._IsLogin(); !ok {
		this.Ctx.Redirect("/")
		return
	} else {
		this.User = u
	}
	this.View.Set("admin",this.User)
}

func (this *Admin) _IsLogin() (user gmc.Mss, isLogin bool) {
	u := this.Session.Get("admin")
	if u != nil && u.(gmc.Mss)["username"] != "" {
		return u.(gmc.Mss), true
	}
	return nil, false
}
