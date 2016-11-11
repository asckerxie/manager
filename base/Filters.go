package base

import (
	"github.com/astaxie/beego/context"
)

//拦截器   验证用户是否已登录
var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("userLogin").(string)
	if !ok && ctx.Request.RequestURI != "/login" {
		ctx.Redirect(302, "/login")
	}
}
