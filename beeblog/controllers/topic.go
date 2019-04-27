package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego"
	"path"
	"strings"
)
type TopicController struct {
	beego.Controller
}
func (this *TopicController)Get(){
	this.Data["IsTopic"] =true
	this.TplName="topic.html"
	this.Data["IsLogin"]=checkAccount(this.Ctx)
	topics,err:=models.GetAllTopics("","",false)
	if err != nil {
		this.Abort("401")
	}else {
		this.Data["Topics"]=topics
	}
}
func (this *TopicController)Post(){
	if checkAccount(this.Ctx)!=true{
		this.Redirect("/login",302)
		return
	}
	//解析表单
	label:=this.Input().Get("label")
	tid:=this.Input().Get("tid")
	title:=this.Input().Get("title")
	content:=this.Input().Get("content")
	category:=this.Input().Get("category")
	//存文件
	//获取附件
	_,fh,err:=this.GetFile("attachment")
	/*if err!=nil {
		this.Abort("this is uploadfile error")
	}*/
	var attachment string
	if fh!=nil {
		//保存附件
		attachment=fh.Filename
		err=this.SaveToFile("attachment",path.Join("attachment",attachment))
		if err!=nil{
			this.Abort(err.Error()+"11111111111111111"+attachment)
		}

	}
	if len(tid)==0{
		err=models.AddTopic(title,category,label,content,attachment)
	}else {
		err=models.ModifyTopic(tid,title,category,label,content,attachment)
	}
	if err != nil {
		this.Abort("405")
	}
	this.Redirect("/topic",302)

}
func (this *TopicController)Add(){
	this.Data["IsLogin"]=checkAccount(this.Ctx)
	this.TplName="topic_add.html"

}
func (this *TopicController)View(){
	this.TplName="topic_view.html"
	topic,err:=models.GetTopic(this.Ctx.Input.Param("0"))
	if err!=nil {
		this.Abort("401")
		this.Redirect("/",302)
		return
	}

	this.Data["Topic"]=topic
	this.Data["Labels"]=strings.Split(topic.Labels," ")
	replies,err:=models.GetALLReplies(this.Ctx.Input.Param("0"))
	if err!=nil {
		this.Abort("401")
		return
	}
	this.Data["Replies"]=replies
	this.Data["IsLogin"]=checkAccount(this.Ctx)
	this.Data["Tid"]=this.Ctx.Input.Param("0")
}
func (this *TopicController)Modify(){
     this.TplName="topic_modify.html"
     tid:=this.Input().Get("tid")
     topic,err:=models.GetTopic(tid)
	if err!=nil {
		this.Abort("401")
		this.Redirect("/",302)
		return
	}
     this.Data["Topic"]=topic
	 this.Data["Tid"]=tid
}
func (this *TopicController)Delete(){
	if !checkAccount(this.Ctx) {
		this.Redirect("/login",302)
		return
	}
	err:=models.DeleteTopic(this.Input().Get("tid"))
	if err !=nil{
		this.Abort("this is Delete error")
	}
	this.Redirect("/topic",302)
}
