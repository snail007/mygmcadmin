package controller

type Main struct {
	Admin
}

func (this *Main) Index() {
	this.View.Layout("index").Render("main/index")
}

func (this *Main) Main() {
	this.View.Layout("page").Render("main/main")
}

func (this *Main) Other() {
	this.View.Layout("list").Render("other/code")
}