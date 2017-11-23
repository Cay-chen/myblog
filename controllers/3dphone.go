package controllers

import "github.com/astaxie/beego"

type PhoneController struct{
	beego.Controller
}
func (this *PhoneController) Get(){
	this.TplName="3dphoto.html"
}