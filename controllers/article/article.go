package article

import (
	"eggache/controllers"
	"eggache/models"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

type ArticleController struct {
	controllers.BaseController
}

func (this *ArticleController) Get() {
	page, err := this.GetInt("p")
	offset, _ := beego.AppConfig.Int("pageoffset")

	if err != nil {
		page = 1
	}
	condArr := make(map[string]int)
	condArr["del"] = 0

	countArticle := models.CountArticle(condArr)
	paginator := pagination.SetPaginator(this.Ctx, offset, countArticle)
	_, _, articles := models.ListArticle(condArr, page, offset)
	this.Data["paginator"] = paginator
	this.Data["articles"] = articles
	this.TplName = "article/article_list.tpl"
}

type AddArticleController struct {
	controllers.BaseController
}

func (this *AddArticleController) Get() {
	this.TplName = "article/write_article.tpl"
}

func (this *AddArticleController) Post() {
	title := this.GetString("title")
	content := this.GetString("content")
	if "" == title {
		title = "未命名"
	}

	var art models.Article
	art.Title = title
	art.Content = content

	_, err := models.AddArticle(art)

	if err == nil {
		this.Redirect("/article", 302)
	} else {
		this.Data["json"] = map[string]interface{}{"code": -1, "message": "添加出错"}
		this.ServeJSON()
	}
}

type EditArticleController struct {
	controllers.BaseController
}

func (this *EditArticleController) Get() {
	idstr := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idstr)

	art, err := models.GetArticle(id)

	if err != nil || art.Del == 1 {
		this.Abort("404")
		//this.Data["json"] = map[string]interface{}{"code": -1, "message": "文章不存在"}
		//this.ServeJSON()
	}

	this.Data["art"] = art
	this.TplName = "article/upd_article.tpl"
}

func (this *EditArticleController) Post() {
	var m = make(map[string]interface{})
	title := this.GetString("title")
	content := this.GetString("content")
	id, err := this.GetInt("id")
	a, err := models.GetArticle(id)
	if err != nil || a.Del == 1 {
		m = map[string]interface{}{"code": -1, "message": "文章不存在"}
		this.Data["json"] = m
		this.ServeJSON()
	}
	if "" == title {
		title = "未命名"
	}

	var art models.Article
	art.Title = title
	art.Content = content

	upderr := models.UpdArticle(id, art)
	if upderr != nil {
		this.Data["json"] = map[string]interface{}{"code": -1, "message": "更改失败"}
	} else {
		this.Redirect("/article", 302)
	}

}

type DelArticleController struct {
	controllers.BaseController
}

func (this *DelArticleController) Get() {
	idstr := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": -1, "message": "转换失败"}
		this.ServeJSON()
	}

	delerr := models.DelArticle(1, id)
	if delerr != nil {
		this.Data["json"] = map[string]interface{}{"code": -1, "message": "删除失败"}
	} else {
		this.Redirect("/article", 302)
	}

}

type ShowArticleController struct {
	controllers.BaseController
}

func (this *ShowArticleController) Get() {
	idstr := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idstr)
	art, err := models.GetArticle(id)

	if err != nil || art.Del == 1 {
		this.Abort("404")
		//this.Data["json"] = map[string]interface{}{"code": -1, "message": "文章不存在"}
		//this.ServeJSON()
	}
	this.Data["art"] = art
	this.TplName = "article/article_detial.tpl"
}
