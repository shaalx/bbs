package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/shaalx/bbs/models"
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
	beego.Debug("login user:", usr, err)
	c.Data["curUser"] = &usr
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
	topic.User = c.CurUser()
	beego.Debug("publishing topic:", topic)
	perr := topic.Publish()
	beego.Debug("published err:", perr)
	if perr != nil {
		c.Abort("400")
	}
	c.Get()
}

// @router /remark/:id:int [post]
func (c *MainController) RemarkTopic() {
	var remark models.Remark
	c.ParseForm(&remark)
	remark.User = c.CurUser()
	var topicid int
	c.Ctx.Input.Bind(&topicid, ":id")
	remark.Topic = models.TopicById(topicid)
	beego.Debug("topic remarking:", remark)
	perr := remark.Publish()
	beego.Debug("remarked err:", perr)
	if perr != nil {
		c.Abort("400")
	}
	c.Get()
}

// @router /topic/:id:int
func (c *MainController) ShowTopic() {
	var idint int
	c.Ctx.Input.Bind(&idint, ":id")
	c.Data["xsrfdata"] = template.HTML(c.XsrfFormHtml())
	c.Data["topic"] = models.TopicById(idint)
	c.Data["remarks"] = models.RemarksById(idint)
	c.TplNames = "topic.html"
}

func (c *MainController) Prepare() {
	user := c.CurUser()
	c.Data["curUser"] = user
	uri := c.Ctx.Request.RequestURI
	if strings.EqualFold(uri, "/publish") || strings.Contains(uri, "/remark") || strings.Contains(uri, "/del") {
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
	if sessid == nil || strings.Contains(fmt.Sprintf("%v", sessid), "_") {
		c.Abort("401")
	}
}

func (c *MainController) LoginSetSession(usrid int) {
	sess, err := models.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil || sess == nil {
		c.Abort("401")
	}

	sess.Set("gosessionid", usrid)
	beego.Debug("set [gosessionid]----->", usrid)
	sess.SessionRelease(c.Ctx.ResponseWriter)
}

func (c *MainController) LogoutSetSession() {
	sess, err := models.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil || sess == nil {
		c.Abort("401")
	}
	sess.Set("gosessionid", "_")
	beego.Debug("set [gosessionid]-----> _")
	sess.SessionRelease(c.Ctx.ResponseWriter)
}

// @router /logout [get]
func (c *MainController) Logout() {
	c.Data["curUser"] = nil
	c.LogoutSetSession()
	c.Get()
}

func (c *MainController) CurUser() *models.User {
	sess, err := models.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil || sess == nil {
		return nil
	}

	iuserid := sess.Get("gosessionid")
	beego.Debug("get [gosessionid] <------- ", iuserid)
	if iuserid == nil {
		return nil
	}
	userid := fmt.Sprintf("%v", iuserid)
	id, err := strconv.Atoi(userid)
	if err != nil {
		return nil
	}
	if id <= 0 {
		return nil
	}
	usr := models.UserById(id)
	if nil == usr {
		return nil
	}
	beego.Debug("current user ----> ", *usr)
	return usr
}

// @router /topic/:topicid:int/delremark/:remarkid:int
func (c *MainController) DelRemark() {
	var topicid int
	var remarkid int
	c.Ctx.Input.Bind(&topicid, ":topicid")
	c.Ctx.Input.Bind(&remarkid, ":remarkid")
	remark := models.RemarkById(remarkid)
	curUser := c.CurUser()
	if remark == nil || (remark.User != nil && curUser != nil && remark.User.Id != curUser.Id) {
		c.Abort("401")
	}
	deled := models.DelRemardById(remarkid)
	beego.Debug("delremark:", deled)
	c.Get()
}

// @router /deltopic/:topicid:int
func (c *MainController) DelTopic() {
	var topicid int
	c.Ctx.Input.Bind(&topicid, ":topicid")
	topic := models.TopicById(topicid)
	curUser := c.CurUser()
	if topic == nil || (topic.User != nil && curUser != nil && topic.User.Id != curUser.Id) {
		c.Abort("401")
	}
	deled := models.DelTopicById(topicid)
	beego.Debug("deltopic:", deled)
	c.Get()
}

// @router /user [get]
func (c *MainController) User() {
	user := c.CurUser()
	if user == nil {
		c.Abort("401")
	}
	c.Data["user"] = user
	c.Data["topics"] = models.TopicsById(user.Id)
	c.Data["remarks"] = models.RemarksByUserId(user.Id)
	c.TplNames = "user.html"
}
