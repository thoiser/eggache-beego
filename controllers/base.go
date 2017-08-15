package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	IsLogin bool
}

func (this *BaseController) Prepare() {
	userLogin := this.GetSession("userLogin")
	if userLogin == nil {
		this.IsLogin = false
	} else {
		this.IsLogin = true
	}
}
