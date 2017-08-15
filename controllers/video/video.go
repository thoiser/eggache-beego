package video

import (
	"eggache/controllers"
	"eggache/models"
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"strconv"
	"time"
)

var (
	accessKey = beego.AppConfig.String("qiniu.accessKey")
	secretKey = beego.AppConfig.String("qiniu.secretKey")
	bucket    = "thoise-go"
)

type VideoController struct {
	controllers.BaseController
}

func (this *VideoController) Get() {
	page, err := this.GetInt("p")
	offset, _ := beego.AppConfig.Int("pageoffset")

	if err != nil {
		page = 1
	}

	condArr := make(map[string]int)
	condArr["del"] = 0

	countVideo := models.CountVideo(condArr)
	paginator := pagination.SetPaginator(this.Ctx, offset, countVideo)
	_, _, videoList := models.ListVideo(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["videoList"] = videoList

	this.TplName = "video/video_list.tpl"
}

type ShowVideoController struct {
	controllers.BaseController
}

func (this *ShowVideoController) Get() {
	idstr := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idstr)

	video, err := models.GetVideo(id)

	if err != nil || video.Del == 1 {
		this.Data["json"] = map[string]interface{}{"code": -1, "message": "视频不存在"}
		this.ServeJSON()
	}

	mac := qbox.NewMac(accessKey, secretKey)
	domain := "http://oue7xfhbg.bkt.clouddn.com"
	key := video.Url
	deadline := time.Now().Add(time.Second * 3600).Unix() //1小时有效期
	video.Url = storage.MakePrivateURL(mac, domain, key, deadline)

	this.Data["video"] = video
	this.TplName = "video/video.tpl"
}

type AddVideoController struct {
	controllers.BaseController
}

func (this *AddVideoController) Post() {
	title := this.GetString("title")
	url := this.GetString("url")

	if "" == title {
		title = "未命名"
	}

	var v models.Video
	v.Title = title
	v.Url = url

	_, addErr := models.AddVideo(v)
	if addErr != nil {
		this.Data["json"] = map[string]interface{}{"code": -1, "message": "上传失败"}
		this.ServeJSON()
	} else {
		this.Redirect("/video", 302)
	}
}

type DelVideoController struct {
	controllers.BaseController
}

func (this *DelVideoController) Post() {
	idstr := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": -1, "message": "转换失败"}
		this.ServeJSON()
	}

	delerr := models.DelVideo(1, id)
	if delerr != nil {
		this.Data["json"] = map[string]interface{}{"code": -1, "message": "删除失败"}
	} else {
		this.Redirect("/video", 302)
	}
}
