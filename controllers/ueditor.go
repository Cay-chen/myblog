package controllers

import (
	"github.com/astaxie/beego"
	"os"
	"log"
	"io/ioutil"
	"fmt"
	"myblog/utils"
	"strconv"
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
type imageinfo struct {
	Url string `json:"url"`
}
type ListImages struct {
	State    string `json:"state"`
	List    interface{} `json:"list"`
	Start int `json:"start"`
	Total int `json:"total"`
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
	uploadImageBC :=UploadImageBC{
				 "/static/img/upfile/" +h.Filename,
				  h.Filename,
				h.Filename,
				 "SUCCESS",
				}
				c.Data["json"]=uploadImageBC
			/*bc,err:=json.Marshal(uploadImageBC)
			if err !=nil{
				beego.Error(err)
			}*/
		//	c.Ctx.WriteString(string(bc))
		//c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "url": "/static/img/upfile/" +h.Filename, "title": h.Filename, "original": h.Filename}

		c.ServeJSON()
	case "listimage":
		start,_:= strconv.Atoi(c.GetString("start"))
		total,_ :=strconv.Atoi( c.GetString("total"))
		fmt.Println(">>>>>>>>>>>>>>>start",start)
		fmt.Println(">>>>>>>>>>>>>>>total",total)

		list,_:=utils.ListDir(".\\static\\img\\upfile","","/static/img/upfile/")

		urllist:= []imageinfo{}
		for _,l :=range list{
			urllist = append(urllist, imageinfo{l})
		}
		c.Data["json"] = ListImages{"SUCCESS",urllist,start,total}
		c.ServeJSON()
	case "deleteimage":
		fmt.Println("shanchu")
		fmt.Println(c.GetString("path"))
		err := os.Remove("."+c.GetString("path"))
		if err!=nil{
			//log.Fatal(err)
			c.Ctx.WriteString("删除失败")

		}else {
			c.Ctx.WriteString("文件删除成功")
		}

	}

}