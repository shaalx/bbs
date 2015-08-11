package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:1234@tcp(localhost:3306)/bbs?charset=utf8")
	orm.RegisterModel(new(User), new(Topic), new(Remark))
	ORM = orm.NewOrm()
	orm.Debug = true
}

var ORM orm.Ormer

func RegisterUser(usr *User) error {
	_, err := ORM.Insert(usr)
	return err
}

func PublishTopic(topic *Topic) error {
	// CheckUser

	n, err := ORM.Insert(topic)
	if n <= 0 {
		return fmt.Errorf("not insert into  db.")
	}
	return err
}

func PublishRemark(remark *Remark) error {
	// CheckUser

	n, err := ORM.Insert(remark)
	if n <= 0 {
		return fmt.Errorf("not insert into  db.")
	}
	return err
}

func CheckUser(usr *User) int {
	err := ORM.QueryTable(usr).Filter("Name", usr.Name).Filter("Passwd", usr.Passwd).One(usr)
	if err != nil {
		return -1
	}
	return usr.Id
}

func AllTopics() []Topic {
	var topics []Topic
	_, err := ORM.QueryTable((*Topic)(nil)).All(&topics)
	if err != nil {
		return nil
	}
	return topics
}

func TopicById(id int) *Topic {
	var topic Topic
	if err := ORM.QueryTable((*Topic)(nil)).Filter("Id", id).One(&topic); err != nil {
		return nil
	}
	return &topic
}

func RemarksById(id int) []Remark {
	var remarks []Remark
	_, err := ORM.QueryTable((*Remark)(nil)).Filter("Topicid", id).All(&remarks)
	if err != nil {
		return nil
	}
	return remarks
}

func UserById(id int) *User {
	var usr User
	if err := ORM.QueryTable((*User)(nil)).Filter("Id", id).One(&usr); err != nil {
		fmt.Println(err, id)
		return nil
	}
	fmt.Println(usr)
	return &usr
}

func DelRemardById(id int) bool {
	n, err := ORM.QueryTable((*Remark)(nil)).Filter("Id", id).Delete()
	if n <= 0 || err != nil {
		return false
	}
	return true
}

func DelTopicById(id int) bool {
	n, err := ORM.QueryTable((*Topic)(nil)).Filter("Id", id).Delete()
	if n <= 0 || err != nil {
		return false
	}
	return true
}
