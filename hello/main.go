package main

import (
	_ "hello/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"


)
func init(){
	//注册默认数据库
	orm.RegisterDataBase("default", "mysql", "root:Ld123456@tcp(116.62.101.115:3306)/beeblog?charset=utf8",30, 30)
	//注册模型
	orm.Debug = true
	}

func main() {
	beego.Run()
	
}

