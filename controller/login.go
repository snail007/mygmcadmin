package controller

import (
	"github.com/snail007/gmc"
	"mygmcadmin/model"
)

type Login struct {
	Base
}

func (this *Login) Auth() {
	err:=this.SessionStart()
	if err!=nil{
		this.Stop(err)
	}
	u := this.Session.Get("admin")
	if u != nil && u.(gmc.Mss)["username"] != "" {
		this.JsonSuccess("","","/main/index")
	}
	username := this.Ctx.POST("username")
	password := this.Ctx.POST("password")
	if password == "" || username == "" {
		this.JsonFail("信息不完整")
	}
	dbUser, err := model.User.GetBy(gmc.M{"username": username})
	if err != nil || len(dbUser) == 0 {
		this.JsonFail("用户名或密码错误")
	}
	if dbUser["password"] != model.EncodePassword(password) {
		this.JsonFail("用户名或密码错误")
	}
	delete(dbUser, "password")
	this.Session.Set("admin", dbUser)
	this.JsonSuccess("","","/main/index")
}

func (this *Login) Index() {
	this.View.Layout("login").Render("login/login")
}

func (this *Login) Logout() {
	this.SessionStart()
	this.SessionDestroy()
	this.Ctx.Redirect("/")
}