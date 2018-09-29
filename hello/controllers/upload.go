package controllers

import (
	"github.com/astaxie/beego"
)

type UploadController struct {
	beego.Controller
}

func (c *UploadController) Get() {
	c.TplName = "upload.html"
}

func (c *UploadController) Post(){
	file,head,err := c.GetFile("file")
	if err != nil{
		beego.Info(err)
		return
	}
	defer file.Close()
	fileName := head.Filename
	err = c.SaveToFile("file","static/"+fileName)
	if err != nil{
		c.Ctx.WriteString("上传失败！")
	}else{
		c.Ctx.WriteString("上传成功！")
	}
}
