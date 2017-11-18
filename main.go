package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myblog/models"
	_ "myblog/routers"

	"net/http"
	"html/template"
)

func init() {
	models.RegisterDB()
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 600
}
func main() {
	//beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	orm.Debug = true
	beego.ErrorHandler("404",page_not_found)
	beego.Run()
}

func page_not_found(rw http.ResponseWriter,r *http.Request){
	t,_:=template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath+"/404.html")
	data := make(map[string] interface{})
	data["content"]="page not found"
	t.Execute(rw,data)

}