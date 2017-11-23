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
	var classify []orm.Params
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
	classify_sql:="SELECT count(case when classify = 101 then 1 end) as man, count(case when classify = 102 then 1 end) as android, count(case when classify = 103 then 1 end) as golang, count(case when classify = 104 then 1 end) as php, count(case when classify = 105 then 1 end) as mysql, count(case when classify = 106 then 1 end) as java, count(case when classify = 107 then 1 end) as html, count(case when classify = 108 then 1 end) as js , count(case when classify = 109 then 1 end) as jq FROM content"
	_,err=o.Raw(classify_sql).Values(&classify)
	c.Data["classify"] =classify[0]
	c.Data["contents"] = maps
	c.TplName = "index.html"

}
