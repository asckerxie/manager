package users

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Users struct {
	Id       int64
	Username string
	Password string
	Avatar   string
	Status   int
}

//登录
func Login(userName string, userPass string) (err error, user Users) {
	o := orm.NewOrm()
	qs := o.QueryTable("user")   //注册表名称
	cond := orm.NewCondition()

	cond = cond.And("username", userName)
	cond = cond.And("password", userPass)
	cond = cond.And("status", 1)

	qs = qs.SetCond(cond)

	var users Users
	err = qs.Limit(1).One(&users, "userid", "username", "avatar")
	fmt.Println(err)

	return err, users
}
