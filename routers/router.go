package routers

import (
	"github.com/astaxie/beego"
	"github.com/toukii/bbs/controllers"
)

func init() {
	beego.Include(&controllers.MainController{}, &controllers.RegisterController{})
	beego.ErrorController(&controllers.ErrorController{})
	// ns := beego.NewNamespace("/",
	// 	// beego.NSRouter("/", &controllers.MainController{}),
	// 	beego.NSInclude(&controllers.MainController{}),
	// 	beego.NSInclude(&controllers.RegisterController{}),
	// )
	// beego.AddNamespace(ns)
}
