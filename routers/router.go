package routers

import (
	"manager/controllers"
	"manager/controllers/users"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &users.LoginController{})
	beego.Router("/main/index", &controllers.MainController{}, "get:Get")

}
