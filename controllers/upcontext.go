package controllers

import (
	"github.com/astaxie/beego"
	"os"
	"github.com/astaxie/beego/orm"
	"fmt"
	"path"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"context"
	"github.com/cay/utils"
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
	path1 := ".\\static\\img\\coverimg"
	err := os.MkdirAll(path1,0777)
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
	path2:= path1+"\\"+h.Filename
	err = c.SaveToFile("coverimg",path2)
	if err !=nil {
		beego.Error(err)
	}
	nameFile := utils.Md5File(path2)+path.Ext(path2)
	resQiNiu(path2,nameFile,"myblog-icon-images")
	delectImage(path2)
	insert_sql := "INSERT INTO content ( title, author,abstract,content,coverimmag,classify,looks,uptime ) VALUES ( '"+title+"','" +author+"','"+ abstract +"','" + context + "','"+"http://icon.blog.image.84jie.cn/"+nameFile+"',"+optionsRadiosinline+",0,now())"
	o := orm.NewOrm()
	_,err = o.Raw(insert_sql).Exec()
	if err ==nil {
		c.TplName = "edit-text.html"
	} else{
		beego.Error(err)
		c.Ctx.WriteString("存入数据库错误")
	}
}
func resQiNiu(localFile,fileName,bucket1 string){
	ak := "U1Z0InLIRWomOHQ3RHBmeBLBeT2LsBsCaTZbLcRC"
	sk :="aXK10xN-DgKVOZNuk4yzTKTjtbL7WtrhtSKyOMWX"
	bucket:=bucket1
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
}