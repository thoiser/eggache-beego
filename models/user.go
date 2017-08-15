package models

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int    `orm:"column(id);auto"`
	Username string `orm:"column(username);size(60)" description:"登录用户名"`
	Password string `orm:"column(password);size(60)" description:"密码"`
}

func (this *User) TableName() string {
	return TableName("user")
}

func LoginUser(username, password string) (int, error) {
	var user User
	err := orm.NewOrm().QueryTable(TableName("user")).Filter("username", username).Filter("password", password).One(&user, "id")

	if err != nil {
		return 0, err
	}
	return user.Id, nil

}

//md5方法
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}
