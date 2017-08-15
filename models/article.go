package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Article struct {
	Id      int    `orm:"column(id);auto"`
	Title   string `orm:"column(title);size(255)"`
	Content string `orm:"column(content);size(2000)"`
	Del     int8   `orm:"column(del)"`
	// Ctime   int64  `orm:"column(ctime);type(timestamp);auto_now_add" description:"添加时间"`
	Ctime time.Time
}

func (this *Article) TableName() string {
	return TableName("article")
}

func CountArticle(condArr map[string]int) int64 {
	qs := orm.NewOrm().QueryTable("a_article")
	cond := orm.NewCondition()
	cond = cond.And("del", condArr["del"])

	num, _ := qs.SetCond(cond).Count()
	// num, _ := qs.Count()
	return num
}

func ListArticle(condArr map[string]int, page int, offset int) (num int64, err error, art []Article) {
	qs := orm.NewOrm().QueryTable("a_article")
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

	var article []Article
	// qs = qs.OrderBy("-knowid")
	num, errs := qs.Limit(offset, start).All(&article)
	return num, errs, article
}

func GetArticle(id int) (Article, error) {
	art := Article{Id: id}
	err := orm.NewOrm().Read(&art)

	return art, err
}

func AddArticle(addArt Article) (int64, error) {
	art := new(Article)

	art.Title = addArt.Title
	art.Content = addArt.Content

	id, err := orm.NewOrm().Insert(art)
	if err != nil {
		return 0, err
	}
	return id, nil

	// article := Article{Title: title, Content: content}
	// id, err := orm.NewOrm().Insert(&article)
	// if err != nil {
	// 	return 0, err
	// }
	// return id, nil
}

func UpdArticle(id int, updArt Article) error {
	art := Article{Id: id}

	art.Title = updArt.Title
	art.Content = updArt.Content
	_, err := orm.NewOrm().Update(&art, "Title", "Content")

	return err
}

func DelArticle(del int8, id int) error {
	art := Article{Id: id, Del: del}

	_, err := orm.NewOrm().Update(&art, "Del")

	return err
}
