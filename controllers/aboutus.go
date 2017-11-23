package controllers

import "github.com/astaxie/beego"

type AboutUsController struct{
	beego.Controller
}
func (this *AboutUsController) Get(){
	this.TplName="lifeTime.html"
}