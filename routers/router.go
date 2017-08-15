package routers

import (
	"eggache/controllers/article"
	"eggache/controllers/photo"
	"eggache/controllers/user"
	"eggache/controllers/video"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &user.IndexController{})
	beego.Router("/login", &user.LoginController{})

	beego.Router("/up", &photo.UpPhotoController{})
	beego.Router("/photo", &photo.PhotoController{})
	beego.Router("/photo/del/:id", &photo.DelPhotoController{})

	beego.Router("/video", &video.VideoController{})
	beego.Router("/video/add", &video.AddVideoController{})
	beego.Router("/video/:id", &video.ShowVideoController{})
	beego.Router("/video/del/:id", &video.DelVideoController{})

	beego.Router("/article", &article.ArticleController{})
	beego.Router("/article/add", &article.AddArticleController{})
	beego.Router("/article/:id", &article.ShowArticleController{})
	beego.Router("/article/del/:id", &article.DelArticleController{})
	beego.Router("/article/edit/:id", &article.EditArticleController{})
}
