package main

import (
	"fmt"
	"manager/base"
	_ "manager/routers"

	"github.com/astaxie/beego"
)

func init() {
	//动态加载不同环境的配置文件
	var runmode = beego.AppConfig.String("runmode")
	switch runmode {
	case "dev":
		beego.LoadAppConfig("ini", "conf/dev.conf")
	case "test":
		beego.LoadAppConfig("ini", "conf/test.conf")
	case "prod":
		beego.LoadAppConfig("ini", "conf/prod.conf")
	default:
		fmt.Println("缺少配置相应的配置文件")
	}

	//加载
	base.MysqlInit()
}

func main() {
	//拦截器 通过session验证用户是否已经登录
	beego.InsertFilter("/",beego.BeforeRouter,base.FilterUser)

	//http错误重定向
	beego.ErrorHandler("404", base.Page_not_found)
	beego.ErrorHandler("401", base.Page_note_permission)

	beego.Run()
}



