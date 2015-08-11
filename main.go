package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	"github.com/shaalx/bbs/models"
	_ "github.com/shaalx/bbs/routers"
)

func main() {
	beego.EnableXSRF = true
	// go TaskSessionGC()
	beego.Run()
}

func TaskSessionGC() {
	tk := toolbox.NewTask("taska", "0/10 * * * * *", func() error {
		fmt.Println("hello world")
		models.GlobalSessions.GC()
		return nil
	},
	)
	err := tk.Run()
	if err != nil {
		beego.Error(err)
	}
	toolbox.AddTask("taska", tk)
	toolbox.StartTask()
}
