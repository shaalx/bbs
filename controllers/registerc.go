package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/everfore/bbs/models"
	"html/template"
)

type RegisterController struct {
	MainController
}

// @router /register [get]
func (c *RegisterController) LoadRegister() {
	c.Data["xsrfdata"] = template.HTML(c.XsrfFormHtml())
	c.TplNames = "register.html"
}

// @router /register [post]
func (c *RegisterController) Register() {
	var usr models.User
	c.ParseForm(&usr)
	beego.Notice(usr)
	valid := validation.Validation{}
	usr.Valid(&valid)
	if valid.HasErrors() {
		c.Abort("401")
	}
	err := models.RegisterUser(&usr)
	if nil != err {
		beego.Debug(err)
		c.Abort("401")
	}
	c.MainController.LoginSetSession(usr.Id)
	c.MainController.Get()
}
