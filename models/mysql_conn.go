package models

import (
"github.com/astaxie/beego/orm"
_ "github.com/go-sql-driver/mysql"
)

const (
	_OB_MYSQL_CONN = "root:0510016@tcp(127.0.0.1:3306)/myblog?charset=utf8"
	_MYSQL_DDRIVER = "mysql"
)
func RegisterDB() {
	orm.RegisterDataBase("default", _MYSQL_DDRIVER, _OB_MYSQL_CONN, 50)
}
