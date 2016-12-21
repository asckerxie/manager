package base

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func MysqlInit() {

	mysql_username := beego.AppConfig.String("mysqluser")
	fmt.Println("mysql_username = ", mysql_username)
	mysql_pass := beego.AppConfig.String("mysqlpass")
	//fmt.Println()
	mysql_host := beego.AppConfig.String("mysqlhost")
	//fmt.Println()
	mysql_dbName := beego.AppConfig.String("mysqldb")
	//fmt.Println()
	mysql_dbPort, err := beego.AppConfig.Int("mysqlport")
	//fmt.Println()

	if err != nil {
		mysql_dbPort = 3306
	}

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

	orm.RegisterDriver("mysql", orm.DRMySQL)

	//orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8")
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", mysql_username, mysql_pass, mysql_host, mysql_dbPort, mysql_dbName))
}
