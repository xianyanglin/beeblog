package controllers
import (
	"beeblog/models"
	_ "beeblog/models"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego"
)

type ReplyController struct {
	beego.Controller
}

func (this *ReplyController) Add(){
	tid:=this.Input().Get("tid")
	err:=models.AddReply(tid,this.Input().Get("nickname"),this.Input().Get("content"))
	if err!=nil {
		this.Abort("this is reply error")
	}
	this.Redirect("/topic/view/"+tid,302)
}
func (this *ReplyController) Delete(){
	if !checkAccount(this.Ctx) {
		return
	}
	tid:=this.Input().Get("tid")
	err:=models.DeleteReply(this.Input().Get("rid"))
	if err!=nil {
		this.Abort("this is reply error")
	}
	this.Redirect("/topic/view/"+tid,302)
}
