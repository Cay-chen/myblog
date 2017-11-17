package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/orm"
	"html/template"
)

type ViewController struct{
	beego.Controller
}
func (c *ViewController) Get(){
	viewID :=c.GetString("idtext")
	if viewID ==""{
		fmt.Println("dadada")
	}
	var maps []orm.Params
	insert_sql := "SELECT * FROM content WHERE id = "+viewID
	o := orm.NewOrm()
	num ,err :=o.Raw(insert_sql).Values(&maps)
	if err !=nil{
		fmt.Println("数据库查询错误")
	}
	if num ==1 {
		c.Data["coverimmag"] =maps[0]["coverimmag"]
		c.Data["author"] =maps[0]["author"]
		c.Data["uptime"] =maps[0]["uptime"]
		c.Data["title"] =maps[0]["title"]
		c.Data["abstract"] =maps[0]["abstract"]
		c.Data["classify"] =maps[0]["classify"]
		c.Data["content"] =template.HTML(maps[0]["content"].(string))
	}
	fmt.Println(maps)
	c.TplName = "about.html"
}