package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type Photo struct {
	Id    int    `orm:"column(id);auto"`
	Title string `orm:"column(title);size(255)"`
	Url   string `orm:"column(url);size(600);null"`
	Del   int8   `orm:"column(del)"`
	Ctime time.Time
}

func (this *Photo) TableName() string {
	return TableName("photo")
}

func CountPhoto(condArr map[string]int) int64 {
	qs := orm.NewOrm().QueryTable("a_photo")
	cond := orm.NewCondition()
	cond = cond.And("del", condArr["del"])

	num, _ := qs.SetCond(cond).Count()
	return num
}

func ListPhoto(condArr map[string]int, page int, offset int) (num int64, err error, p []Photo) {
	qs := orm.NewOrm().QueryTable("a_photo")
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

	var photo []Photo
	// qs = qs.OrderBy("-knowid")
	num, errs := qs.Limit(offset, start).All(&photo)
	return num, errs, photo
}

func AddPhoto(addP Photo) (int64, error) {
	p := new(Photo)

	p.Title = addP.Title
	p.Url = addP.Url

	id, err := orm.NewOrm().Insert(p)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func DelPhoto(del int8, id int) error {
	p := Photo{Id: id, Del: del}

	_, err := orm.NewOrm().Update(&p, "Del")

	return err
}
