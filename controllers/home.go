package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	fmt.Println("Aaaa")
	c.TplName = "edit-text.html"

}
