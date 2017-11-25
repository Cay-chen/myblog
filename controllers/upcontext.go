package controllers

import (
	"github.com/astaxie/beego"
	"os"
	"github.com/astaxie/beego/orm"
	"path"
	"github.com/cay-chen/golang-utils/utils"
	"errors"
	"fmt"
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
	//创建上传文件夹
	//localImagePath:= "./static/img/upfile" //Linux 系统
	localImagePath:= ".\\static\\img\\upfile" //Windows 系统
	err := os.MkdirAll(localImagePath,0777)
	if err !=nil{
		beego.Error(err)
	}
	//获取上传的文件，直接可以获取表单名称对应的文件名，不用另外提取
	_, h, err := c.GetFile("coverimg")
	if err != nil {
		beego.Error(err)
	}
	//保存上传的图片
	//saveImagePath:= localImagePath+"/"+h.Filename //Linux 系统
	saveImagePath:= localImagePath+"\\"+h.Filename //Windows 系统
	fmt.Printf(saveImagePath)
	err = c.SaveToFile("coverimg",saveImagePath)
	if err !=nil {
		fmt.Println("获取文件错误")

		beego.Error(err)
	}
	fileName:= utils.Md5File(saveImagePath)+path.Ext(saveImagePath) //获取文件MD5值保存
	if resQiNiu(saveImagePath,fileName,"myblog-icon-images") ==nil{
		deleteImage(saveImagePath) //删除文件
		insert_sql := "INSERT INTO content ( title, author,abstract,content,coverimmag,classify,looks,uptime ) VALUES ( '"+title+"','" +author+"','"+ abstract +"','" + context + "','"+"http://icon.blog.image.84jie.cn/"+fileName+"',"+optionsRadiosinline+",0,now())"
		o := orm.NewOrm()
		_,err = o.Raw(insert_sql).Exec()
		if err ==nil {
			c.TplName = "edit-text.html"
		} else{
			beego.Error(err)
			c.Ctx.WriteString("存入数据库错误")
		}

	}else {
		beego.Error(errors.New("七牛返回错误"))
		c.Ctx.WriteString("存入数据库错误")

	}

}
