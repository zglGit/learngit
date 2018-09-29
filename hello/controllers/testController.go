package controllers

import (
	"github.com/astaxie/beego"
	"hello/models"
	"github.com/astaxie/beego/orm"
)

type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {
	//user , err := models.QueryUserByUserId(1)
	//if err != nil {
	//	beego.Info(err)
	//}
	//c.Data["User"] = user
	//c.TplName = "index.html"

}

func (c *TestController) QueryUserByUserId(){
	user , err := models.QueryUserByUserId(1)
	if err != nil {
		beego.Info(err)
	}
	c.Data["User"] = user
	c.TplName = "index.html"
}

func (c *TestController) QueryProfileByUser(){
	profile , err := models.QueryProfileByUser(1)
	if err != nil{
		beego.Info(err)
	}
	c.Data["Profile"] = profile
	c.TplName = "index.html"
}


func (c *TestController) QueryUserByPostTitle(){
	user , err := models.QueryUserByPostTitle("666")
	if err != nil{
		beego.Info(err)
	}
	c.Data["User"] = user
	c.TplName = "index.html"
}


func (c *TestController) QueryPostbyUserId(){
	posts, err := models.QueryPostbyUserId(1)
	if err != nil{
		beego.Info(err)
	}
	c.Data["Posts"] = posts
	c.TplName = "index.html"
}


func (c *TestController) QueryAll(){
	user , err := models.QueryUserByUserId(1)
	if err != nil {
		beego.Info(err)
	}
	c.Data["User"] = user

	profile , err := models.QueryProfileByUser(1)
	if err != nil{
		beego.Info(err)
	}
	c.Data["Profile"] = profile

	user2 , err := models.QueryUserByPostTitle("666")
	if err != nil{
		beego.Info(err)
	}
	c.Data["User2"] = user2

	posts, err := models.QueryPostbyUserId(1)
	if err != nil{
		beego.Info(err)
	}
	c.Data["Posts"] = posts
	c.TplName = "index.html"
}

//User的增刪該查
func (c *TestController) TestMethods(){
	//
	//user := models.User{Id:2,Name:"lc",ProfileId:22}
	//n ,err := orm.NewOrm().Insert(&user)
	//if err != nil && n>0{
	//	beego.Info(err)
	//}
	//
	//user := models.User{Id :2}
	//n,err := orm.NewOrm().Delete(&user)
	//if n>0 && err!= nil{
	//	beego.Info(err)
	//}

	//user:= models.User{Id:2,Name:"wangyang",ProfileId:22}
	//n,err := orm.NewOrm().Update(&user)
	//if err != nil&&n > 0{
	//	beego.Info(err)
	//}
	user:= models.User{}
	err := orm.NewOrm().QueryTable("user").Filter("Id",2).One(&user)
	if err != nil{
		beego.Info(err)
	}
	c.Data["User"] = user
	c.TplName = "showInfo.html"


}

