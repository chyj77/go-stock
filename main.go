package main

import (
	"fmt"
	_ "go-stock/routers"

	models "go-stock/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mysql"
)

func init() {
	fmt.Println("test go mysql")
	url := beego.AppConfig.String("mysql::url")
	fmt.Println("test go url = ", url)
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册默认数据库
	orm.RegisterDataBase("default", "mysql", url)
	//注册model
	orm.RegisterModel(new(models.Ztsj))
	//自动建表
	// orm.RunSyncdb("default", false, true)
	orm.Debug = true
}

func main() {
	fmt.Println("test go run")
	beego.Run()
}
