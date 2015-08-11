package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/everfore/bbs/models"
	"html/template"
	"strconv"
	"strings"
)

type MainController struct {
	beego.Controller
}

// @router / [get]
func (c *MainController) Get() {
	c.Data["topics"] = models.AllTopics()
	c.TplNames = "index.html"
}

// @router /login [get]
func (c *MainController) LoadLogin() {
	c.Data["xsrfdata"] = template.HTML(c.XsrfFormHtml())
	c.TplNames = "login.html"
}

// @router /login [post]
func (c *MainController) Login() {
	var usr models.User
	err := c.ParseForm(&usr)
	beego.Debug(usr, err)
	if err != nil {
		c.Abort("403")
	}
	uid := usr.Check()
	if uid <= 0 {
		c.Abort("401")
	}
	c.LoginSetSession(uid)
	c.Get()
}

// @router /publish [get]
func (c *MainController) LoadPublishTopic() {
	c.Data["xsrfdata"] = template.HTML(c.XsrfFormHtml())
	c.TplNames = "publish.html"
}

// @router /publish [post]
func (c *MainController) PublishTopic() {
	var topic models.Topic
	c.ParseForm(&topic)
	topic.Userid = c.CurUser().Id
	beego.Debug(topic)
	perr := topic.Publish()
	beego.Debug(perr)
	if perr != nil {
		c.Abort("400")
	}
	c.Get()
}

// @router /remark/:id:int [get]
func (c *MainController) LoadRemarkTopic() {
	id := c.Ctx.Input.Param(":id")
	c.Data["topicid"] = id
	c.Data["xsrfdata"] = template.HTML(c.XsrfFormHtml())
	c.TplNames = "remark.html"
}

// @router /remark/:id:int [post]
func (c *MainController) RemarkTopic() {
	var remark models.Remark
	c.ParseForm(&remark)
	remark.Userid = c.CurUser().Id
	beego.Debug(remark)
	perr := remark.Publish()
	beego.Debug(perr)
	if perr != nil {
		c.Abort("400")
	}
	c.Get()
}

// @router /topic/:id:int
func (c *MainController) ShowTopic() {
	var idint int
	c.Ctx.Input.Bind(&idint, ":id")
	beego.Debug(idint)
	c.Data["topic"] = models.TopicById(idint)
	c.Data["remarks"] = models.RemarksById(idint)
	c.TplNames = "show_topic.html"
}

func (c *MainController) Prepare() {
	uri := c.Ctx.Request.RequestURI
	beego.Notice(uri)
	if strings.EqualFold(uri, "/publish") || strings.Contains(uri, "/remark") {
		c.CheckLogin()
	}
}

func (c *MainController) CheckLogin() {
	sess, err := models.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil || sess == nil {
		c.Abort("401")
	}
	sessid := sess.Get("gosessionid")
	beego.Debug(sessid)
	if sessid == nil || !strings.EqualFold(fmt.Sprintf("%v", sessid), "LOGIN_USER") {
		c.Abort("401")
	}
}

func (c *MainController) LoginSetSession(usrid int) {
	sess, err := models.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil || sess == nil {
		c.Abort("401")
	}

	sess.Set("gosessionid", "LOGIN_USER")
	sess.Set("current_user", usrid)
	sess.SessionRelease(c.Ctx.ResponseWriter)
}

func (c *MainController) LogoutSetSession() {
	sess, err := models.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil || sess == nil {
		c.Abort("401")
	}
	sess.Set("gosessionid", "_")
	sess.SessionRelease(c.Ctx.ResponseWriter)
}

// @router /logout [get]
func (c *MainController) Logout() {
	c.LogoutSetSession()
	c.Get()
}

func (c *MainController) CurUser() *models.User {
	sess, err := models.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil || sess == nil {
		c.Abort("401")
	}

	iuserid := sess.Get("current_user")
	if iuserid == nil {
		iuserid = "1"
	}
	userid := fmt.Sprintf("%v", iuserid)
	id, err := strconv.Atoi(userid)
	if err != nil {
		id = 1
	}
	if id <= 0 {
		id = 1
	}
	usr := models.UserById(id)
	if nil == usr {
		c.Abort("500")
	}
	beego.Notice(usr)
	return usr
}

// @router /topic/:topicid:int/del/:remarkid:int
func (c *MainController) RemarkDel() {
	var topicid int
	var remarkid int
	c.Ctx.Input.Bind(&topicid, ":topicid")
	c.Ctx.Input.Bind(&remarkid, ":remarkid")
	deled := models.DelRemardById(remarkid)
	beego.Debug(deled)
	// c.ShowTopic()
	// c.Redirect(fmt.Sprintf("/topic/%d", topicid), 200)
	c.Get()
}

// @router /topicdel/:id:int
func (c *MainController) TopicDel() {
	var topicid int
	c.Ctx.Input.Bind(&topicid, ":id")
	deled := models.DelTopicById(topicid)
	beego.Debug(deled)
	// c.Redirect("/", 200)
	c.Get()
}
