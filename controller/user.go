package controller

import (
	"github.com/snail007/gmc"
	"mygmcadmin/model"
	"time"
)

type User struct {
	Admin
}

func (this *User) Add() {
	this.View.Layout("form").Render("user/form")
}

func (this *User) Create() {
	username := this.Ctx.POST("username")
	nickname := this.Ctx.POST("nickname")
	password := this.Ctx.POST("password")
	password1 := this.Ctx.POST("password1")
	if username == "" || password == "" || password != password1 {
		this.JsonFail("信息不完整")
	}
	dbUser, err := model.User.GetBy(gmc.M{"username": username})
	if err != nil || len(dbUser) > 0 {
		this.JsonFail("用户已经存在")
	}
	now:=time.Now().Unix()
	data := gmc.M{
		"username": username,
		"nickname":nickname,
		"password": model.EncodePassword(password),
		"create_time":now,
		"update_time":now,
	}
	_, err = model.User.Insert(data)
	this.StopE(err, func() {
		this.JsonFail(err.Error())
	}, func() {
		this.JsonSuccess("", nil, "/user/list")
	})
}

func (this *User) Edit() {
	userID:=this.Ctx.GET("id")
	if userID==""{
		this.Ctx.Redirect("/user/list")
	}
	user,err:=model.User.GetByID(userID)
	this.StopE(err, func() {
		this.JsonFail(err.Error())
	})
	this.View.Set("user",user)
	this.View.Layout("form").Render("user/form")
}

func (this *User) Save() {
	nickname := this.Ctx.POST("nickname")
	password := this.Ctx.POST("password")
	password1 := this.Ctx.POST("password1")
	userID:=this.Ctx.GET("id")
	if userID==""{
		this.Ctx.Redirect("/user/list")
	}
	if password != ""  {
		if password != password1{
			this.JsonFail("信息不完整")
		}
		password=model.EncodePassword(password)
	}
	dbUser,err:=model.User.GetByID(userID)
	if err != nil || len(dbUser) == 0 {
		this.JsonFail("用户不存在")
	}
	data := gmc.M{
		"nickname": nickname,
		"update_time":time.Now().Unix(),
	}
	if password!=""{
		data["data"]=password
	}
	_, err = model.User.UpdateByIDs([]string{userID},data)
	this.StopE(err, func() {
		this.JsonFail(err.Error())
	}, func() {
		this.JsonSuccess("", nil, "/user/list")
	})
}

func (this *User) Delete() {
	var ids []string
	this.Request.ParseForm()
	id := this.Request.Form["ids"]
	if len(id) > 0 {
		ids = append(ids, id...)
	}
	for _,v:=range ids{
		if v=="1"{
			this.JsonFail("系统账号禁止删除")
		}
	}
	_, err := model.User.UpdateByIDs(ids, gmc.M{"is_delete": 1})
	this.StopE(err, func() {
		this.JsonFail(err.Error())
	})
	this.JsonSuccess("", nil, "/user/list")
}

func (this *User) List() {
	search_field:=this.Ctx.GET("search_field")
	keyword:=this.Ctx.GET("keyword")
	where:=gmc.M{"is_delete": 0}
	rule:=map[string]bool{"username":true,"nickname":true,"user_id":true}
	if search_field!=""&&keyword!=""&&rule[search_field]{
		if search_field=="user_id"{
			where[search_field]=keyword
		}else{
			where[search_field+" like"]="%"+keyword+"%"
		}
	}
	perPage := gmc.ToInt(this.Ctx.GET("count"))
	if perPage <= 0 || perPage > 100 {
		perPage = 10
	}
	offset := perPage * (gmc.ToInt(this.Ctx.GET("page")) - 1)
	if offset < 0 {
		offset = 0
	}
	users, total, err := model.User.Page(where, offset, perPage)
	this.StopE(err, func() {
		this.WriteE(err)
	})
	this.View.Set("users", users)
	this.View.Set("paginator", this.Ctx.NewPager(perPage, int64(total)))
	this.View.Layout("list").Render("user/list")
}

