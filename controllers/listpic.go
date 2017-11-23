package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math"
	"strconv"
	"html/template"
//	"github.com/astaxie/beego/cache"
)

type ListpiceControllers struct{
	beego.Controller
}
func(this *ListpiceControllers) Get(){
	classify :=this.GetString("classify")
	search :=this.GetString("keyboard")
	if classify==""&&search!="" {
		SearchContent(this)
	}else {
		ClassifyContent(this)
	}

	//fmt.Println(maps)
	this.TplName="newslistpic.html"
}
func (this *ListpiceControllers) Post(){
	SearchContent(this)
	this.TplName="newslistpic.html"
}
func SearchContent(this *ListpiceControllers){
	page := this.GetString("page")
	pages,_ :=strconv.Atoi(page)
	keyboard := this.GetString("keyboard")
	start := strconv.Itoa((pages-1)*6)
	select_mysql:="SELECT a.id,a.title,a.abstract,a.uptime,a.coverimmag,b.classify FROM content as a INNER JOIN classify as b ON a.classify = b.id_class WHERE a.title like '%"+keyboard+"%' LIMIT "+start+",6"
	count_mysql := "SELECT COUNT(*) AS count FROM content as a INNER JOIN classify as b ON a.classify =b.id_class WHERE a.title like '%"+keyboard+"%'"
	var counts []orm.Params
	o:=orm.NewOrm()
	var maps []orm.Params
	_,err:=o.Raw(select_mysql).Values(&maps)
	if err!=nil{
		this.Abort("404")
	}
	_,err =o.Raw(count_mysql).Values(&counts)
	if err!=nil{
		this.Abort("404")
	}
	total_coutn := counts[0]["count"]
	this.Data["total_count"] = total_coutn
	total_num ,_:= strconv.Atoi(total_coutn.(string))
	total_page :=math.Ceil(float64(float64(total_num)/float64(6)))
	this.Data["contents"]=maps
	this.Data["total_page"]=total_page
	this.Data["now_page"]=page
	down_page:=strconv.Itoa(pages+1)
	up_page:=strconv.Itoa(pages-1)
	this.Data["contents"]=maps
	var up_herf string
	var down_herf string
	var start_herf string
	var end_herf string

	if pages ==1 {
		up_herf="<span>上一页 </span>"
	}else{
		up_herf="<a href='/listpic?page="+up_page+"&keyboard="+keyboard+"'>上一页 </a>"
	}
	if pages ==int(total_page) {
		down_herf="<span> 下一页 </span>"
	}else{
		down_herf="<a href='/listpic?page="+down_page+"&keyboard="+keyboard+"'>下一页 </a>"
	}
	start_herf ="<a href='/listpic?page=1&keyboard="+keyboard+"'>首页 </a>"
	end_herf ="<a href='/listpic?page="+strconv.Itoa(int(total_page))+"&keyboard="+keyboard+"'>尾页</a>"
	this.Data["up_down"]=template.HTML(start_herf+up_herf+down_herf+end_herf)
}
func ClassifyContent(this *ListpiceControllers){
	classify := this.GetString("classify")
	getpage := this.GetString("page")
	page,_ := strconv.Atoi(getpage)
	start := strconv.Itoa((page-1)*6)
	var select_mysql string
	var count_mysql string
	var up_herf string
	var down_herf string
	var start_herf string
	var end_herf string
	if classify ==""{
		select_mysql="SELECT a.id,a.title,a.abstract,a.uptime,a.coverimmag,b.classify FROM content as a INNER JOIN classify as b ON a.classify = b.id_class LIMIT "+start+",6"
		count_mysql = "SELECT COUNT(*) AS count FROM content as a INNER JOIN classify as b ON a.classify =b.id_class"
		this.Data["isClassify"] = false
	}else{
		select_mysql="SELECT a.id,a.title,a.abstract,a.uptime,a.coverimmag,b.classify FROM content as a INNER JOIN classify as b ON a.classify = b.id_class WHERE a.classify = "+classify+" LIMIT "+start+",6"
		count_mysql="SELECT COUNT(*) as count FROM content as a INNER JOIN classify as b ON a.classify = b.id_class WHERE a.classify = "+classify
		this.Data["isClassify"] = true
	}
	o := orm.NewOrm()
	var maps []orm.Params
	var counts []orm.Params
	_,err:=o.Raw(select_mysql).Values(&maps)
	if err!=nil{
		this.Abort("404")
	}
	_,err=o.Raw(count_mysql).Values(&counts)
	if err!=nil{
		this.Abort("404")
	}
	total_coutn := counts[0]["count"]
	this.Data["total_count"] = total_coutn
	total_num ,_:= strconv.Atoi(total_coutn.(string))
	total_page :=math.Ceil(float64(float64(total_num)/float64(6)))
	this.Data["contents"]=maps
	this.Data["total_page"]=total_page
	this.Data["now_page"]=page
	down_page:=strconv.Itoa(page+1)
	up_page:=strconv.Itoa(page-1)
	this.Data["classify"]=classify
	if page ==1 {
		up_herf="<span>上一页 </span>"
	}else{
		up_herf="<a href='/listpic?page="+up_page+"&classify="+classify+"'>上一页 </a>"
	}
	if page ==int(total_page) {
		down_herf="<span> 下一页 </span>"
	}else{
		down_herf="<a href='/listpic?page="+down_page+"&classify="+classify+"'>下一页 </a>"
	}
	start_herf ="<a href='/listpic?page=1&classify="+classify+"'>首页 </a>"
	end_herf ="<a href='/listpic?page="+strconv.Itoa(int(total_page))+"&classify="+classify+"'>尾页</a>"
	this.Data["up_down"]=template.HTML(start_herf+up_herf+down_herf+end_herf)
}