package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"os"
	"github.com/astaxie/beego/orm"
)

type UpContextControllers struct{
	beego.Controller
}

func (c *UpContextControllers) Post(){
	title := c.GetString("title")
	author := c.GetString("author")
	context := c.GetString("context")
	optionsRadiosinline := c.GetString("optionsRadiosinline")
	abstract := c.GetString("abstract")


	fmt.Println("--------------title:",c.GetString("title"))
	fmt.Println("--------------author:",c.GetString("author"))
	fmt.Println("--------------context:",c.GetString("context"))
	fmt.Println("--------------optionsRadiosinline:",c.GetString("optionsRadiosinline"))
	fmt.Println("--------------abstract:",c.GetString("abstract"))
	//创建上传文件夹
	path := ".\\static\\img\\coverimg"
	err := os.MkdirAll(path,0777)
	if err !=nil{
		beego.Error(err)
	}
	//获取上传的文件，直接可以获取表单名称对应的文件名，不用另外提取
	_, h, err := c.GetFile("coverimg")
	fmt.Println("filename----------",h.Filename)
	if err != nil {
		beego.Error(err)
	}
	//保存上传的图片
	path1:= path+"\\"+h.Filename
	fmt.Println("path--------",path1)
	err = c.SaveToFile("coverimg",path1)
	if err !=nil {
		beego.Error(err)
	}
	insert_sql := "INSERT INTO content ( title, author,abstract,content,coverimmag,classify,uptime ) VALUES ( '"+title+"','" +author+"','"+ abstract +"','" + context + "','"+"/static/img/coverimg/"+h.Filename+"','"+optionsRadiosinline+"',now())"
	o := orm.NewOrm()
	_,err = o.Raw(insert_sql).Exec()
	if err ==nil {
		c.TplName = "edit-text.html"
	} else{
		c.Ctx.WriteString("存入数据库错误")
	}
}