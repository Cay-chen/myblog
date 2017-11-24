package controllers

import (
	"github.com/astaxie/beego"
	"github.com/cay-chen/golang-utils/utils"
	"os"
	"io/ioutil"
	"fmt"
	"path"
	"errors"
)
type UEditorController struct {
	beego.Controller
}
type UploadImageBC struct {
	Url      string `json:"url"`
	Title    string `json:"title"`
	Original string `json:"original"`
	State    string `json:"state"`
}
type ListImages struct {
	State    string `json:"state"`
	List    interface{} `json:"list"`
	Start int `json:"start"`
	Total int `json:"total"`
}
func (c *UEditorController) GetAndPost(){
	ac := c.GetString("action")
	if ac=="" {
		fmt.Println("Error:action 为空")
		c.Abort("404")
	}
	switch ac{
	case "config":
		file,err := os.Open("conf/config.json")
		if err!=nil{
			beego.Error(err)
			os.Exit(1)
		}
		defer file.Close()
		fd ,err :=ioutil.ReadAll(file)
		if err !=nil{
			beego.Error(err)
			os.Exit(1)
		}
		js := string(fd)
		c.Ctx.WriteString(js)
	case "uploadimage","uploadfile", "uploadvideo":
		//创建上传文件夹
		//localImagePath:= "./static/img/upfile" //Linux 系统
		localImagePath:= ".\\static\\img\\upfile" //Windows 系统
		err := os.MkdirAll(localImagePath,0777)
		if err !=nil{
			beego.Error(err)
		}
		//获取上传的文件，直接可以获取表单名称对应的文件名，不用另外提取
		_, h, err := c.GetFile("upfile")
		if err != nil {
			beego.Error(err)
			}
		//保存上传的图片
		//saveImagePath:= localImagePath+"/"+h.Filename //Linux 系统
		saveImagePath:= localImagePath+"\\"+h.Filename //Windows 系统
		err = c.SaveToFile("upfile",saveImagePath)
		if err !=nil {
			beego.Error(err)
		}
		fileName:= utils.Md5File(saveImagePath)+path.Ext(saveImagePath) //获取文件MD5值保存
		if resQiNiu(saveImagePath,fileName,"myblog-content-images") ==nil{
			deleteImage(saveImagePath) //删除文件
			uploadImageBC :=UploadImageBC{
					 "/" +fileName,
					  h.Filename,
					h.Filename,
					 "SUCCESS",
					}
					c.Data["json"]=uploadImageBC
					c.ServeJSON()
				}else {
					beego.Error(errors.New("七牛返回错误"))
		}
	default:
		fmt.Println("Error:没有找到对应的action")
	}

}

//删除文件
func deleteImage(file string) {
	err := os.Remove(file)
	if err != nil {
		//如果删除失败则输出 file remove Error!
		fmt.Println("file remove Error!")
		//输出错误详细信息
		fmt.Printf("%s", err)
	}
}
