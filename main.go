package main

import (
	"eggache/models"
	_ "eggache/routers"
	"html/template"
	"net/http"

	"github.com/astaxie/beego/context"

	"github.com/astaxie/beego"
)

func main() {
	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	beego.ErrorHandler("404", page_not_found)
	beego.ErrorHandler("401", service_unavailable)
	models.Init()
	beego.Run()
}

var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("userLogin").(string)
	if !ok && ctx.Request.RequestURI != "/" {
		ctx.Redirect(302, "/")
	}
}

func page_not_found(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("404.tpl").ParseFiles("views/404.tpl")
	data := make(map[string]interface{})
	t.Execute(rw, data)
}

func service_unavailable(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("401.tpl").ParseFiles("views/503.tpl")
	data := make(map[string]interface{})
	t.Execute(rw, data)
}
