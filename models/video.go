package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type Video struct {
	Id    int    `orm:"column(id);auto"`
	Url   string `orm:"column(url);size(600);null"`
	Title string `orm:"column(title);size(255)"`
	Del   int8   `orm:"column(del)"`
	Ctime time.Time
}

func (this *Video) TableName() string {
	return TableName("video")
}

func CountVideo(condArr map[string]int) int64 {
	qs := orm.NewOrm().QueryTable("a_video")
	cond := orm.NewCondition()
	cond = cond.And("del", condArr["del"])

	num, _ := qs.SetCond(cond).Count()
	return num
}

func ListVideo(condArr map[string]int, page int, offset int) (num int64, err error, v []Video) {
	qs := orm.NewOrm().QueryTable("a_video")
	cond := orm.NewCondition()
	cond = cond.And("del", condArr["del"])

	qs = qs.SetCond(cond)

	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset

	var video []Video
	// qs = qs.OrderBy("-knowid")
	num, errs := qs.Limit(offset, start).All(&video)
	return num, errs, video
}

func AddVideo(addV Video) (int64, error) {
	v := new(Video)

	v.Title = addV.Title
	v.Url = addV.Url

	id, err := orm.NewOrm().Insert(v)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func DelVideo(del int8, id int) error {
	art := Photo{Id: id, Del: del}

	_, err := orm.NewOrm().Update(&art, "Del")

	return err
}

func GetVideo(id int) (Video, error) {
	video := Video{Id: id}
	err := orm.NewOrm().Read(&video)

	return video, err
}
