package controllers

import (
	"github.com/astaxie/beego"
	"github.com/cay/utils"
	"os"
	"log"
	"io/ioutil"
	"fmt"
	"path"
	"strconv"
	//"path/filepath"
	"time"
	"math/rand"
)
var r = rand.New(rand.NewSource(time.Now().UnixNano()))
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
		path1 := ".\\static\\img\\upfile"
		err := os.MkdirAll(path1,0777)
		if err !=nil{
			beego.Error(err)
		}
		//获取上传的文件，直接可以获取表单名称对应的文件名，不用另外提取
		_, h, err := c.GetFile("upfile")
		if err != nil {
			beego.Error(err)
			}
		//保存上传的图片
		path2:= path1+"\\"+h.Filename
		fmt.Println("path--------",path1)
		err = c.SaveToFile("upfile",path2)
		fmt.Println("filename----------",path.Ext(path2))
		if err !=nil {
			beego.Error(err)
		}
		nameFile := utils.Md5File(path2)+path.Ext(path2)
		resQiNiu(path2,nameFile,"myblog-content-images")
		delectImage(path2)
		uploadImageBC :=UploadImageBC{
				 "/" +nameFile,
				  h.Filename,
				h.Filename,
				 "SUCCESS",
				}
				c.Data["json"]=uploadImageBC
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
/*func resQiNiu(localFile,fileName string){
	ak := "U1Z0InLIRWomOHQ3RHBmeBLBeT2LsBsCaTZbLcRC"
	sk :="aXK10xN-DgKVOZNuk4yzTKTjtbL7WtrhtSKyOMWX"
	bucket:="myblog-content-images"
	key := fileName

	mac := qbox.NewMac(ak,sk)

	putPolicy := storage.PutPolicy{
		Scope:bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	putPolicy.Expires=7200
	upToken :=putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := MyPutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "myblog content image",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("'''''''''''''''''''",ret.Bucket, ret.Key, ret.Fsize, ret.Hash, ret.Name)
}*/
// 自定义返回值结构体
type MyPutRet struct {
	Key    string
	Hash   string
	Fsize  int
	Bucket string
	Name   string
}
//删除文件
func delectImage(file string) {
	err := os.Remove(file)               //删除文件test.txt
	if err != nil {
		//如果删除失败则输出 file remove Error!
		fmt.Println("file remove Error!")
		//输出错误详细信息
		fmt.Printf("%s", err)
	} else {
		//如果删除成功则输出 file remove OK!
		fmt.Print("file remove OK!")
	}

}
func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return r.Int63n(max-min) + min
}