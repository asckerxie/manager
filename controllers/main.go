package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	//"fmt"
	//"strconv"
	//"strings"
	"strconv"
	"strings"
)

type BaseController struct {
	beego.Controller
	IsLogin      bool
	UserUserId   int64
	UserUsername string
	UserAvatar   string
}

func (this *BaseController) Prepare() {

	fmt.Println("项目 prepare 方法中，sessionName = ", beego.AppConfig.String("sessionName")+"_user")

	userLogin := this.GetSession(beego.AppConfig.String("sessionName") + "_user") // 获取session信息

	fmt.Println("项目 prepare 方法中，userLogin_session = ", userLogin)

	if userLogin == nil {
		this.IsLogin = false
	} else {
		this.IsLogin = true
		tmp := strings.Split(userLogin.(string), "||")

		userid, _ := strconv.Atoi(tmp[0])
		longid := int64(userid)
		this.Data["LoginUserid"] = longid   //用户id
		this.Data["LoginUsername"] = tmp[1] //用户登录名称
		this.Data["LoginAvatar"] = tmp[2]   //用户头像

		//this.UserUserId = longid
		//this.UserUsername = tmp[1]
		//this.UserAvatar = tmp[2]
		//
		////消息
		//msgcondArr := make(map[string]string)
		//msgcondArr["touserid"] = fmt.Sprintf("%d", longid)
		//msgcondArr["view"] = "1"

		//countTopMessage := CountMessages(msgcondArr)
		//_, _, topMessages := ListMessages(msgcondArr, 1, 10)
		//this.Data["topMessages"] = topMessages
		//this.Data["countTopMessage"] = countTopMessage
	}

	fmt.Println("####$$$$$$$$$$$$************", this.IsLogin)
	this.Data["IsLogin"] = this.IsLogin
}

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	isLogin_ := c.BaseController.IsLogin

	fmt.Println("登录,判断是否已经登录. isLogin = ", isLogin_)

	if isLogin_ {
		//
		//uid := BaseController{}.UserUserId
		//userName := BaseController{}.UserUsername
		//avatar := BaseController{}.UserAvatar

		//查询用户角色信息   首先从session中获取  若没有 则通过sql获取 并将信息存在session中
		c.GetSession(beego.AppConfig.String("sessionName") + "_role")

		//查询用户菜单信息

		c.TplName = "main/main.html"
	} else {
		c.TplName = "login.html"
	}
}
