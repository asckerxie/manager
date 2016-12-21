package users

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int64
	Username string
	Password string
	Avatar   string
	Status   int
}

func init() {
	orm.RegisterModel(new(User))
}

//登录
func Login(userName string, userPass string) (err error, user User) {
	o := orm.NewOrm()
	qs := o.QueryTable("user") //注册表名称
	cond := orm.NewCondition()

	cond = cond.And("username", userName)
	cond = cond.And("password", userPass)
	cond = cond.And("status", 1)

	qs = qs.SetCond(cond)

	var users User
	err = qs.Limit(1).One(&users, "id", "username", "avatar", "password")
	fmt.Println(err)
	fmt.Println(users.Username)

	return err, users
}
