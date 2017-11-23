package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/orm"
	"html/template"
	"strconv"
)

type ViewController struct{
	beego.Controller
}
func (c *ViewController) Get(){
	viewID :=c.GetString("idtext")
	if viewID ==""{
		c.Abort("404")
		fmt.Println("没有疏导数据")
	}
	var maps []orm.Params
	//insert_sql := "SELECT * FROM content WHERE id = "+viewID
	select_sql := "SELECT a.uptime,a.title,a.content,a.looks,b.classify FROM content as a INNER JOIN classify as b ON a.classify = b.id_class WHERE a.id = "+viewID
	o := orm.NewOrm()
	num ,err :=o.Raw(select_sql).Values(&maps)
	if err !=nil{
		c.Abort("404")
		fmt.Println("数据库查询错误")
	}
	if num ==1 {
		//更新浏览次数
		go UpdateLooks(viewID)

		//	c.Data["coverimmag"] =maps[0]["coverimmag"]
	//	c.Data["author"] =maps[0]["author"]
		c.Data["uptime"] =maps[0]["uptime"]
		c.Data["title"] =maps[0]["title"]
		looks,_:=strconv.Atoi(maps[0]["looks"].(string))
		c.Data["looks"] =looks
		//	c.Data["abstract"] =maps[0]["abstract"]
		c.Data["classify"] =maps[0]["classify"]
		c.Data["content"] =template.HTML(maps[0]["content"].(string))
	}else{
		c.Abort("404")
	}
	c.TplName = "content.html"
}
func UpdateLooks(id string){
	select_sql := "UPDATE content SET looks = looks+1 where id ="+ id
	o := orm.NewOrm()
	_ ,err :=o.Raw(select_sql).Exec()
	if err !=nil{
		fmt.Println("数据库查询错误")
	}

}