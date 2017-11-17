package routers

import (
	"myblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/controller",&controllers.UEditorController{},"*:GetAndPost")
	beego.Router("/upcontext", &controllers.UpContextControllers{})
	beego.Router("/view", &controllers.ViewController{})

}
