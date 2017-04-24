package users

import (
	"encoding/json"
	"fmt"
	"manager/controllers"
	. "manager/models/users"

	"github.com/astaxie/beego"
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
		l.ServeJSON()
	}

	if "" == userPass {
		//密码为空
		l.Data["json"] = map[string]interface{}{"code": 1002, "message": "密码不能为空"}
		l.ServeJSON()
	}

	isExcite := ValidataName(userName)
	if isExcite { // 存在该用户
		err, users := Login(userName, userPass)

		if err == nil {

			b, err_ := json.Marshal(users)
			if err_ == nil {
				fmt.Println("登录用户信息:" + string(b))
			}

			fmt.Println("##########################  登录成功，设置session   sessionName = ", beego.AppConfig.String("sessionName")+"_user")

			// 用户登录成功  ->  将用户信息 插入session中
			l.SetSession(beego.AppConfig.String("sessionName")+"_user", fmt.Sprintf("%d", users.Id)+"||"+users.Username+"||"+users.Avatar)

			fmt.Println("##########################  登录成功，设置session  = ", l.GetSession(beego.AppConfig.String("sessionName")+"_user"))

			l.Data["json"] = map[string]interface{}{"code": 0, "message": "贺喜你，登录成功"}
		} else {
			l.Data["json"] = map[string]interface{}{"code": 1, "message": "登录失败"}
		}
	} else {
		l.Data["json"] = map[string]interface{}{"code": 1003, "message": "用户不存在"}
	}

	l.ServeJSON()
}
