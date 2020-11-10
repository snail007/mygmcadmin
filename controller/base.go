package controller

import "github.com/snail007/gmc"

type Base struct {
	gmc.Controller
}

func (this *Base) JsonSuccess(msg string, data ...interface{}) {
	this.JSON(msg,200,data...)
}

func (this *Base) JsonFail(msg string) {
	this.JSON(msg,500)
}

func (this *Base) JSON(msg string, code int, data ...interface{}) {
	var data0 interface{}
	var url1 interface{}
	if len(data) >= 1 {
		data0 = data[0]
	}
	if len(data) >= 2 {
		url1 = data[1]
	}
	d := gmc.M{
		"msg":  msg,
		"code": code,
		"data": data0,
		"url":  url1,
	}
	this.Write(d)
	this.Stop()
}
