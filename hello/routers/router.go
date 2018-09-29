//@APIVersion 1.0.0
//@Title
package routers

import (
	"hello/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login", &controllers.LoginController{})
    beego.Router("/category", &controllers.CategoryController{})
	beego.AutoRouter(&controllers.TopicController{})  //自动路由
    beego.Router("/topic", &controllers.CategoryController{})
	beego.Router("/testUser",&controllers.TestController{},"get:QueryUserByUserId")
	//beego.AutoRouter(&controllers.TestController{})
	beego.Router("/testProfile",&controllers.TestController{},"get:QueryProfileByUser")
	beego.Router("/testPost",&controllers.TestController{},"get:QueryUserByPostTitle")
	beego.Router("/testPosts",&controllers.TestController{},"get:QueryPostbyUserId")
	beego.Router("/testAll",&controllers.TestController{},"get:QueryAll")
	beego.Router("/testAllMethods",&controllers.TestController{},"get:TestMethods")

	beego.Router("/upload",&controllers.UploadController{})

}
