package controllers

import "github.com/astaxie/beego"

type HomeControllers struct{
	beego.Controller
}
func (this *HomeControllers) Get(){
	this.TplName="home.html"
}