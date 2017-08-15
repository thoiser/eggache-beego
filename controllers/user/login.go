package user

import (
	"eggache/controllers"
	"eggache/models"
)

type IndexController struct {
	controllers.BaseController
}

func (this *IndexController) Get() {
	this.TplName = "index.tpl"
}

type LoginController struct {
	controllers.BaseController
}

func (this *LoginController) Get() {
	check := this.BaseController.IsLogin
	if check {
		this.Redirect("/photo", 302)
	} else {
		this.TplName = "login/login.tpl"
	}
}

func (this *LoginController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	password = models.Substr(models.Md5(password), 8, 16)

	_, err := models.LoginUser(username, password)

	if err == nil {
		this.SetSession("userLogin", "shit")
		this.Redirect("/photo", 302)
	} else {
		this.Redirect("/login", 302)
	}
	this.ServeJSON()
}
