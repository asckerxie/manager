package users

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int64
	RoleId   int64
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

	if err != nil {
		fmt.Println("登录查询,用户登录过程中，未查询到用户信息.", err)
	}

	return err, users
}

//验证用户名是否存在
func ValidataName(userName string) bool {
	result := false

	o := orm.NewOrm()
	qs := o.QueryTable("user") //注册表名称
	cond := orm.NewCondition()

	cond = cond.And("username", userName)

	qs = qs.SetCond(cond)

	var users User

	err := qs.Limit(1).One(&users, "id", "username", "avatar", "password")

	if err != nil {
		fmt.Println("验证用户名是否存在时候,出现异常.", err)
	} else {
		if users.Id > 0 {
			result = true
		}
	}

	return result
}
