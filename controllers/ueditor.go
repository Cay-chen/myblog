package controllers

import (
	"github.com/astaxie/beego"
	"os"
	"log"
	"io/ioutil"
	"fmt"
//	"encoding/json"
)

type UEditorController struct {
	beego.Controller
}
type UploadImageBC struct {
	url      string
	title    string
	original string
	state    string
}
func (c *UEditorController) GetAndPost(){
	ac := c.GetString("action")
	fmt.Println("======="+ac)
	switch ac{
	case "config":
		file,err := os.Open("conf/config.json")
		if err!=nil{
			log.Fatal(err)
			os.Exit(1)
		}
		defer file.Close()
		fd ,err :=ioutil.ReadAll(file)
		if err !=nil{
			log.Fatal(err)
			os.Exit(1)
		}
		js := string(fd)
		c.Ctx.WriteString(js)
	case "uploadimage","uploadfile", "uploadvideo":
		//创建上传文件夹
		path := ".\\static\\img\\upfile"
		err := os.MkdirAll(path,0777)
		if err !=nil{
			beego.Error(err)
		}
		//获取上传的文件，直接可以获取表单名称对应的文件名，不用另外提取
		_, h, err := c.GetFile("upfile")
		fmt.Println("filename----------",h.Filename)
		if err != nil {
			beego.Error(err)
			}
		//保存上传的图片
		path1:= path+"\\"+h.Filename
		fmt.Println("path--------",path1)
		err = c.SaveToFile("upfile",path1)
		if err !=nil {
			beego.Error(err)
		}
		/*uploadImageBC :=UploadImageBC{
			 "/static/img/upload/" +h.Filename,
			  h.Filename,
			h.Filename,
			 "SUCCESS",
			}
			//c.Data["json"]=uploadImageBC
		bc,err:=json.Marshal(uploadImageBC)
		if err !=nil{
			beego.Error(err)
		}
		fmt.Println("--------------",string(bc))
		c.Ctx.WriteString(string(bc))*/
		c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "url": "/static/img/upfile/" +h.Filename, "title": h.Filename, "original": h.Filename}

		c.ServeJSON()

	}

}