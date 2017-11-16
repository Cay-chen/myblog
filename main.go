package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myblog/models"
	_ "myblog/routers"

)

func init() {
	models.RegisterDB()
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 600
}
func main() {
	//beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	orm.Debug = true
	beego.Run()
}

