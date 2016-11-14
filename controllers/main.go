package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"strconv"
	"strings"
)

type BaseController struct {
	beego.Controller
	IsLogin bool
	UserUserId   int64
	UserUsername string
	UserAvatar   string
}

func (this *BaseController) Prepare() {
	userLogin := this.GetSession(beego.AppConfig.String("sessionName")+"_user")   // 获取session信息
	if userLogin == nil {
		this.IsLogin = false
	} else {
		this.IsLogin = true
		tmp := strings.Split(userLogin.(string), "||")

		userid, _ := strconv.Atoi(tmp[0])
		longid := int64(userid)
		this.Data["LoginUserid"] = longid
		this.Data["LoginUsername"] = tmp[1]
		this.Data["LoginAvatar"] = tmp[2]

		this.UserUserId = longid
		this.UserUsername = tmp[1]
		this.UserAvatar = tmp[2]

		//消息
		msgcondArr := make(map[string]string)
		msgcondArr["touserid"] = fmt.Sprintf("%d", longid)
		msgcondArr["view"] = "1"

		//countTopMessage := CountMessages(msgcondArr)
		//_, _, topMessages := ListMessages(msgcondArr, 1, 10)
		//this.Data["topMessages"] = topMessages
		//this.Data["countTopMessage"] = countTopMessage
	}
	this.Data["IsLogin"] = this.IsLogin
}

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	isLogin_ := BaseController{}.IsLogin

	if isLogin_ {
		c.TplName = "index.html"
	}else {
		c.TplName = "login.html"
	}
}


