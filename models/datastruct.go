package models

import (
	"github.com/astaxie/beego/validation"
)

type User struct {
	Id int
	// Topic    *Topic `orm:rel(one)`
	Name     string `form:"username" orm:"unique;pk"`
	Passwd   string `form:"password" orm:"passwd"`
	RePasswd string `form:"repassword" orm:"-"`
	Xsrf     string `form:"_xsrf" orm:"-"`
}

func (u *User) Valid(v *validation.Validation) {
	v.Required(u.Name, "shaalx")
	if u.Passwd != u.RePasswd {
		v.SetError("passwd", "repassword does not equal.")
	}
}

func (u *User) Check() int {
	return CheckUser(u)
}

func (u *User) TabelUnique() [][]string {
	return [][]string{
		[]string{"Name"},
	}
}

type Topic struct {
	Id      int    `orm:"id;pk" form:"-"`
	Userid  int    `orm:"userid" form:"-"`
	Title   string `orm:"title" form:"title"`
	Content string `orm:"content;null" form:"content"`
}

func (t *Topic) Publish() error {
	return PublishTopic(t)
}

type Remark struct {
	Id      int    `orm:"id;pk" form:"-"`
	Topicid int    `orm:"topicid" form:"topicid"`
	Userid  int    `orm:"userid" form:"-"`
	Content string `orm:"content" form:"content"`
}

func (r *Remark) Publish() error {
	return PublishRemark(r)
}
