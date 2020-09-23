package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type User struct {
	Id   int64
	Name string
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(192.168.17.115:3306)/golong?charset=utf8", 30, 30)

	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)
}

func main() {
	AddUser()
}

func AddUser() *User {
	o := orm.NewOrm()
	user := User{Name: "deng"}

	id, err := o.Insert(&user)

	fmt.Printf("ID:%d ,ERR:%v\n", id, err)
	user.Id = id
	return &user
}
