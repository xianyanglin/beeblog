package controllers
import (
	"beeblog/models"
	_ "beeblog/models"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego"
)

/*type MainController struct {
	beego.Controller
}*/
type MainController struct {
	beego.Controller
}
func (this *MainController) Get() {
	/*	this.Data["Website"] = "beego.me"
		this.Data["Email"] = "astaxie@gmail.com"
		this.TplName = "index.tpl"
		this.Data["TrueCond"]=true
		this.Data["FalseCond"]=false
		type u struct {
			Name string
			Age int
			Sex string
		}
		user :=&u{
			Name:"joe",
			Age:20,
			Sex:"Male",
		}
		this.Data["User"]=user
		nums :=[]int{1,2,3,4,5,6,7,8,9,0}
		this.Data["Nums"]=nums
		this.Data["tplVar"]="hey guys"

		this.Data["Html"]="<div>hello beego</div>"*/
	topics,err:=models.GetAllTopics(this.Input().Get("cate"),this.Input().Get("label"),true)
	Categories,err:=models.GetAllCategories()
	if err != nil {
		this.Abort("401")
	}else {
		this.Data["Topics"]=topics
	}
	this.Data["Categories"]=Categories
	this.Data["IsHome"]=true
	this.TplName="home.html"
	this.Data["IsLogin"] =checkAccount(this.Ctx)


}
