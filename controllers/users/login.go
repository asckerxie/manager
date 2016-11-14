package users

import (
	. "manager/models/users"
	"manager/controllers"
	"github.com/astaxie/beego"
	"fmt"
)

type LoginController struct {
	controllers.BaseController
}

func (l *LoginController) Get() {
	l.TplName = "login.html"
}

func (l *LoginController) Post() {
	userName := l.GetString("username")
	userPass := l.GetString("userpass")

	if "" == userName {
		//用户名为空
		l.Data["json"] = map[string]interface{}{"code": 1001, "message": "用户名不能为空"}
	}

	if "" == userPass {
		//密码为空
		l.Data["json"] = map[string]interface{}{"code": 1002, "message": "密码不能为空"}
	}

	err, users := Login(userName, userPass)

	if err == nil {
		// 用户登录成功  ->  将用户信息 插入session中
		l.SetSession(beego.AppConfig.String("sessionName") + "_user", fmt.Sprintf("%d", users.Id) + "||" + users.Username + "||" + users.Avatar)

		l.Data["json"] = map[string]interface{}{"code": 0, "message": "贺喜你，登录成功"}
	} else {
		l.Data["result"] = map[string]interface{}{"code":1, "message":"登录失败"}
	}
	l.ServeJSON()
}


