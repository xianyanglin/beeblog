package main

import (
	"os"
	_ "os"
	"beeblog/controllers"
	"beeblog/models"
	_ "beeblog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/orm"
)
func init(){
	models.RegitsterDB()
}

func main() {
	orm.Debug=true
	orm.RunSyncdb("default",false,false)

	beego.Router("/",&controllers.MainController{})
	beego.Router("/login",&controllers.LoginController{})
	beego.Router("/category",&controllers.CategoryController{})
	beego.Router("/topic",&controllers.TopicController{})
	beego.Router("/topic/view/reply",&controllers.ReplyController{})
	beego.Router("/topic/view/reply/add",&controllers.ReplyController{},"post:Add")
	beego.Router("/topic/view/reply/delete",&controllers.ReplyController{},"Get:Delete")
	beego.Router("/attachment/:all",&controllers.AttachController{})
	beego.AutoRouter(&controllers.TopicController{})
	//创建目录
	os.Mkdir("attachment",os.ModePerm)
	//作为静态文件处理
	/*beego.SetStaticPath("/attachment","attachment")*/
	//作为单独一个控制器来处理

	beego.Run()
}

