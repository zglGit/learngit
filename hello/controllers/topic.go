package controllers

import(
	"github.com/astaxie/beego"
	"hello/models"
)

//@Topic API
type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get(){
	topics ,err := models.GetAllTopics()
	if err != nil{
		beego.Error(err)
	}else{
		c.Data["Topics"] = topics
		return
	}
}
//主要是添加文章和更新文章
func (c *TopicController) Post(){
	if !checkAccount(c.Ctx) {
		//301永久性重定向和302临时性重定向
		c.Redirect("/login",302)
		return
	}
	title := c.GetString("title")
	content := c.GetString("content")
	err := models.AddTopic(title,content)
	if err!= nil {
		beego.Error(err)
	}
	c.Redirect("/topic",302)
	return
}
//自动路由匹配到
//Param
func(c *TopicController) Views(){
  c.Ctx.Input.Param("0")
}

//自动路由匹配到
func (C *TopicController) Add(){

}
