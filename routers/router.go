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
	beego.Router("/editor",&controllers.EditorController{})
	beego.Router("/aboutus",&controllers.AboutUsController{})
	beego.Router("/3dphoto",&controllers.PhoneController{})
	beego.Router("/listpic",&controllers.ListpiceControllers{})
	beego.Router("/home",&controllers.HomeControllers{})

}
