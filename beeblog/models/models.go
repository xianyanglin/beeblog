package models

import (
	_ "github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	_ "path"
	"strconv"
	"strings"
	"time"
)

/*const (
	_DB_NAME="data/beeblog.sql"
	_SQLITE3_DRIVER="sqlite3"
)*/

type Category struct {
	Id int64
	Title string
	Created time.Time `orm:"index"`
	Views int64 `orm:"index"`
	TopicTime time.Time `orm:"index"`
	TopicCount int64
	TopicLastUserId int64
}
type Topic struct {
	Id int64
	Uid int64
	Title string
	Category string
	Labels string
	Content string `orm:"size(5000)"`
	Attachment string
	Created time.Time `orm:"index"`
	Updated time.Time `orm:"index"`
	Views int64
	Author string
	ReplyTime time.Time `orm:"index"`
	ReplyCount int64
	RepleyLastUserId int64

}
//评论
type Comment struct {
	Id int64
	Tid int64
	Name string
	Content string `orm:size(1000)`
	Created time.Time `orm:"index"`
}
func RegitsterDB(){
/*	if !com.IsExist(_DB_NAME) {
		os.Mkdir(path.Dir(_DB_NAME),os.ModePerm)
		os.Create(_DB_NAME)
	}*/
	orm.RegisterModel(new(Category),new(Topic),new(Comment))
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default","mysql","myuser:mypassword@/beeblog",10)
}
func AddTopic(title,category,label,content,attachment string) error {
	//处理标签
	label="$"+strings.Join(strings.Split(label," "),"#$")+"#"
	//空格作为多个标签的分隔符
	//"beego orm"
	//[beego orm]
	//$beego#$orm#

	o:=orm.NewOrm()
	topic:=&Topic{
		Title:title,
		Category:category,
		Labels:label,
		Content:content,
		Attachment:attachment,
		Created:time.Now(),
		Updated:time.Now(),
	}
	_,err:=o.Insert(topic)
	if err!=nil{
		return err
	}
	//更新分类统计
	cate :=new(Category)
	qs :=o.QueryTable("category")
	err =qs.Filter("title",category).One(cate)
	if err==nil {
		cate.TopicCount++
		_,err=o.Update(cate)
	}
	return err
}
func AddCategory(name string) error{
	o:=orm.NewOrm()
	cate:= &Category{Title: name}
	qs :=o.QueryTable("Category")
	err :=qs.Filter("title",name).One(cate)
	if err ==nil {
		return err
	}
	_, err = o.Insert(cate)
	if err!=nil {
		return err
	}
	return nil
}
func AddReply(tid,nickname,content string)error{
    tidNum,err:=strconv.ParseInt(tid,10,64)
	if err!=nil {
		return err
	}

    reply :=&Comment{
    	Tid:tidNum,
    	Name:nickname,
    	Content:content,
    	Created:time.Now(),
	}
    o:=orm.NewOrm()
    _,err=o.Insert(reply)
	if err!=nil {
		return err
	}
    topic :=&Topic{Id:tidNum}
	if o.Read(topic)==nil {
		topic.ReplyTime=time.Now()
		topic.ReplyCount++
		_,err=o.Update(topic)
	}
    return err
}
func DelCategory(id string) error{
	cid,err :=strconv.ParseInt(id,10,64)
	if err!=nil {
		return err
	}
	o :=orm.NewOrm()
	cate :=&Category{Id:cid}
	_,err=o.Delete(cate)
	return err

}
func DeleteTopic(tid string) error{
	id,err :=strconv.ParseInt(tid,10,64)
	if err!=nil {
		return err
	}
	var oldCate string
	o :=orm.NewOrm()
	topic :=&Topic{Id:id}
	if o.Read(topic)==nil {
		oldCate=topic.Category
		_,err=o.Delete(topic)
		if err!=nil {
			return err
		}
	}
	if len(oldCate)>0 {
		cate :=new (Category)
		qs :=o.QueryTable("category")
		err :=qs.Filter("title",oldCate).One(cate)
		if err==nil{
			cate.TopicCount--
			_,err =o.Update(cate)
		}
	}
	return err

}
func DeleteReply(rid string) error{
	ridNum,err :=strconv.ParseInt(rid,10,64)
	if err!=nil {
		return err
	}
	o :=orm.NewOrm()
	var tidNum int64
	replie := &Comment{Id: ridNum}
	if o.Read(replie)==nil {
		tidNum =replie.Tid
		_,err=o.Delete(replie)
		if err!=nil {
			return err
		}
	}
	replies:=make([]*Comment,0)
	qs:=o.QueryTable("comment")
	_,err=qs.Filter("tid",tidNum).OrderBy("-created").All(&replies)
	if err!=nil {
		return err
	}
	topic :=&Topic{Id:tidNum}
	if o.Read(topic)==nil {
		topic.ReplyTime=replies[0].Created
		topic.ReplyCount=int64(len(replies))
		_,err=o.Update(topic)
	}
	return err
}
func GetALLReplies(tid string)(replies []*Comment,err error ){
	tidNum,err :=strconv.ParseInt(tid,10,64)
	if err!=nil {
		return nil,err
	}
	replies =make([]*Comment,0)
	o:=orm.NewOrm()
	qs :=o.QueryTable("comment")
	_,err=qs.Filter("tid",tidNum).All(&replies)
	return  replies,err

}
func GetAllTopics(cate string,label string,isDesc bool)([]*Topic,error){
	o:=orm.NewOrm()
	topics :=make([]*Topic,0)
	qs :=o.QueryTable("topic")
   var err error
	if isDesc {
		if len(cate)>0{
			qs=qs.Filter("category",cate)
		}
		if len(label)>0{
			qs=qs.Filter("labels","$"+label+"#")
		}
		_,err=qs.OrderBy("-created").All(&topics)
	}else {
		_,err =qs.All(&topics)
	}

	return  topics,err
}
func GetAllCategories()([]*Category,error){
	o :=orm.NewOrm()
	cates :=make([]*Category,0)
	qs :=o.QueryTable("category")
	_,err :=qs.All(&cates)
	return cates,err
}
func GetTopic(tid string)(*Topic ,error){
	tidNum,err:=strconv.ParseInt(tid,10,64)
	if err !=nil{
		return nil,err
	}
	o:=orm.NewOrm()
	topic :=new(Topic)
	qs:=o.QueryTable("topic")
	err =qs.Filter("id",tidNum).One(topic)
	if err!=nil {
		return nil,err
	}
	topic.Views++
	_, err=o.Update(topic)

	topic.Labels=strings.Replace(strings.Replace(topic.Labels,"#"," ",-1),"$","",-1)
	return topic,err
}
func ModifyTopic(tid,title,category,label,content,attachment string)error{
	//处理标签
	label="$"+strings.Join(strings.Split(label," "),"#$")+"#"
	tidNum,err:=strconv.ParseInt(tid,10,64)
	if err!=nil {
		return err
	}
	var oldCate,oldAttch string
	o:=orm.NewOrm()
	topic :=&Topic{Id:tidNum}
	if o.Read(topic)==nil{
		oldCate=topic.Category
		oldAttch=topic.Attachment
		topic.Category=category
		topic.Labels=label
		topic.Title=title
		topic.Content=content
		topic.Attachment=attachment
		topic.Updated=time.Now()
		_,err=o.Update(topic)
		if err!=nil {
			return err
		}
	}
	//更新分类统计
	if len(oldCate)>0 {
		cate :=new (Category)
		qs :=o.QueryTable("category")
		err :=qs.Filter("title",oldCate).One(cate)
		if err==nil{
			cate.TopicCount--
			_,err =o.Update(cate)
		}
	}
	//删除旧的附件
	if len(oldAttch)>0 {
		os.Remove(path.Join("attachment",oldAttch))
	}

	cate :=new(Category)
	qs:=o.QueryTable("category")
	err=qs.Filter("title",category).One(cate)
	if err==nil{
		cate.TopicCount++
		_,err =o.Update(cate)
	}
     return nil
}