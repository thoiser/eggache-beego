package photo

import (
	"context"
	"eggache/controllers"
	"eggache/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"log"
	"strconv"
	"time"
)

var (
	accessKey = beego.AppConfig.String("qiniu.accessKey")
	secretKey = beego.AppConfig.String("qiniu.secretKey")
	bucket    = "thoise-go"
)

type PhotoController struct {
	controllers.BaseController
}

func (this *PhotoController) Get() {
	page, err := this.GetInt("p")
	offset, _ := beego.AppConfig.Int("pageoffset")

	if err != nil {
		page = 1
	}

	condArr := make(map[string]int)
	condArr["del"] = 0

	countPhoto := models.CountPhoto(condArr)
	paginator := pagination.SetPaginator(this.Ctx, offset, countPhoto)
	_, _, photoList := models.ListPhoto(condArr, page, offset)

	for i := 0; i < len(photoList); i++ {
		mac := qbox.NewMac(accessKey, secretKey)
		domain := "http://oue7xfhbg.bkt.clouddn.com"
		key := photoList[i].Url
		deadline := time.Now().Add(time.Second * 3600).Unix() //1小时有效期
		photoList[i].Url = storage.MakePrivateURL(mac, domain, key, deadline)
	}

	this.Data["paginator"] = paginator
	this.Data["photoList"] = photoList

	this.TplName = "photo/photo.tpl"
}

type UpPhotoController struct {
	controllers.BaseController
}

func (this *UpPhotoController) Post() {
	title := this.GetString("title")
	f, h, err := this.GetFile("file_name")

	if "" == title {
		title = "未命名"
	}

	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer f.Close()

	s := time.Now().Unix()
	key := fmt.Sprintf("%d", s)
	//生成新的文件名
	filename := key + h.Filename

	this.SaveToFile("file_name", "tmp/upload/"+filename) // 保存位置在 static/upload, 没有文件夹要先创建

	localFile := "tmp/upload/" + filename

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuabei
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	resumeUploader := storage.NewResumeUploader(&cfg)
	upToken := putPolicy.UploadToken(mac)

	ret := storage.PutRet{}

	qnErr := resumeUploader.PutFile(context.Background(), &ret, upToken, filename, localFile, nil)
	if qnErr != nil {
		fmt.Println(qnErr)
		return
	} else {

		var p models.Photo
		p.Title = title
		p.Url = filename

		_, addErr := models.AddPhoto(p)
		if addErr != nil {
			this.Data["json"] = map[string]interface{}{"code": -1, "message": "上传失败"}
			this.ServeJSON()
		} else {
			this.Redirect("/photo", 302)
		}
	}

}

type DelPhotoController struct {
	controllers.BaseController
}

func (this *DelPhotoController) Get() {
	idstr := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": -1, "message": "转换失败"}
		this.ServeJSON()
	}

	delerr := models.DelPhoto(1, id)
	if delerr != nil {
		this.Data["json"] = map[string]interface{}{"code": -1, "message": "删除失败"}
	} else {
		this.Redirect("/article", 302)
	}
}
