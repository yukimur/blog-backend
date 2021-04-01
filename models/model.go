package models

import (
	// "errors"
	// "strconv"
	// "time"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
    // set default database
	orm.RegisterDataBase("default", "mysql", "ubuntu:a438552317a!@tcp(127.0.0.1:3306)/blog?charset=utf8&loc=Local")
	orm.RegisterModel(new(User),new(IPUser),new(Blog),new(Comment))
    // create table
    orm.RunSyncdb("default", false, true)
}

