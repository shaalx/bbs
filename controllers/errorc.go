package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Errorlogin() {
	c.Data["xsrfdata"] = template.HTML(c.XsrfFormHtml())
	c.TplNames = "login.html"
}
