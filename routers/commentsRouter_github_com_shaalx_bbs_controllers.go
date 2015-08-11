package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/shaalx/bbs/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/shaalx/bbs/controllers:MainController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/bbs/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/shaalx/bbs/controllers:MainController"],
		beego.ControllerComments{
			"LoadLogin",
			`/login`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/bbs/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/shaalx/bbs/controllers:MainController"],
		beego.ControllerComments{
			"Login",
			`/login`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/bbs/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/shaalx/bbs/controllers:MainController"],
		beego.ControllerComments{
			"LoadPublishTopic",
			`/publish`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/bbs/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/shaalx/bbs/controllers:MainController"],
		beego.ControllerComments{
			"PublishTopic",
			`/publish`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/bbs/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/shaalx/bbs/controllers:MainController"],
		beego.ControllerComments{
			"LoadRemarkTopic",
			`/remark/:id:int`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/bbs/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/shaalx/bbs/controllers:MainController"],
		beego.ControllerComments{
			"RemarkTopic",
			`/remark/:id:int`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/bbs/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/shaalx/bbs/controllers:MainController"],
		beego.ControllerComments{
			"ShowTopic",
			`/topic/:id:int`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/bbs/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/shaalx/bbs/controllers:MainController"],
		beego.ControllerComments{
			"Logout",
			`/logout`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/bbs/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/shaalx/bbs/controllers:MainController"],
		beego.ControllerComments{
			"RemarkDel",
			`/topic/:topicid:int/del/:remarkid:int`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/bbs/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/shaalx/bbs/controllers:MainController"],
		beego.ControllerComments{
			"TopicDel",
			`/topicdel/:id:int`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/bbs/controllers:RegisterController"] = append(beego.GlobalControllerRouter["github.com/shaalx/bbs/controllers:RegisterController"],
		beego.ControllerComments{
			"LoadRegister",
			`/register`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/bbs/controllers:RegisterController"] = append(beego.GlobalControllerRouter["github.com/shaalx/bbs/controllers:RegisterController"],
		beego.ControllerComments{
			"Register",
			`/register`,
			[]string{"post"},
			nil})

}
