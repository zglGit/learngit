package controllers
import (
	"github.com/astaxie/beego"
	"hello/models"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	op := c.GetString("op")
	switch op {
		case "add" :
			name := c.GetString("name")
			if len(name) ==0 {
				break
			}
			err := models.AddCategory(name)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/category",301)
			return

		case "del":
			id := c.GetString("id")
			if len(id) ==0 {
				break
			}
			err := models.DelCategory(id)
			if err != nil{
				beego.Error(err)
			}
			c.Redirect("/category",301)
			return
	}

	c.Data["Categories"] , _ = models.GetAllCategories()
	c.TplName = "category.html"
}


