package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	var maps []orm.Params
	//insert_sql := "SELECT * FROM content WHERE id = "+viewID
	select_sql := "SELECT a.id,a.title,a.coverimmag ,a.uptime,b.classify,a.abstract,a.author FROM content as a inner join classify as b ON a.classify = b.id_class LIMIT 0,6"
	o := orm.NewOrm()
	num ,err :=o.Raw(select_sql).Values(&maps)
	if err !=nil{
		c.Abort("404")
		fmt.Println("数据库查询错误")
	}
	if num<1{
		c.Abort("404")
	}
	c.Data["contents"] = maps
	c.TplName = "index.html"

}
