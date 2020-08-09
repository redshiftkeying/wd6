package model

import (
	"crypto/md5"
	"fmt"
	"time"
)

type User struct {
	ID                string    `xorm:"pk" json:"id"`
	UserName          string    `xorm:"varchar(64) unique notnull" json:"username"`
	ShowName          string    `xorm:"varchar(64)" json:"name"` // 显示名称
	PassWord          string    `xorm:"varchar(32) notnull" json:"password"`
	Phone             string    `xorm:"varchar(32) " json:"phone"` // 电话号码
	IsAdmin           bool      `xorm:"TINYINT(1)" json:"isadmin"`
	RoleGroupID       string    `xorm:"notnull" json:"role_group_id"` // 权限组ID
	Domain            string    `xorm:"varchar(255)" json:"domain"`
	NeedChagePassword bool      `xorm:"TINYINT(1)" json:"changepassword"` // 需要修改密码才可登录
	Type              string    `xorm:"varchar(18)" json:"type"`          // 保留字段
	Email             string    `xorm:"varchar(255)" json:"email"`
	Avatar            string    `xorm:"varchar(255)" json:"avatar"`
	OrgID             string    `xorm:"varchar(255) notnull" json:"org_id"` // 归属组织机构ID
	Comments          string    `xorm:"varchar(255)" json:"comments"`       // 备注
	UpdateTime        time.Time `xorm:"updated" json:"-"`
	CreateTime        time.Time `xorm:"created" json:"-"`
	DeletedTime       time.Time `xorm:"deleted" json:"-"` // 软删除
}

// function Auth basic authentication name and passwd
func (u *User) Auth(name string, passwd string) bool {
	typeName, err := u.FindByNameTypes(name)
	if err == nil && u.ID != "" {
		if u.PassWord == passwd && typeName != "" {
			return true
		}
	}
	return false
}

// function FindByNameTypes find user use name,showname,phone and so on,
// return type of field name as json tags.
func (u *User) FindByNameTypes(input string) (nameType string, err error) {
	//--------------------------------------------
	querys := map[string]string{"username": "user_name=?", "name": "show_name=?", "phone": "phone=?", "email": "email=?"}
	//--------------------------------------------
	for typeName, query := range querys {
		// TODO: goroutines
		if has, err := engine.Where(query, input).Get(u); err == nil && has {
			return typeName, nil
		} else if err != nil {
			return "", err
		}
	}
	return "", nil
}

// Encryption functon
func (u *User) Encryption(pwd string) {
	sal := "hell12owords#" // TDODO:viper function (platform_id)
	str := pwd + sal
	data := []byte(str)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	u.PassWord = md5str1
}

// CURD interface start
func (u *User) Update() (e error) {
	_, e = engine.ID(u.ID).Update(u)
	if u.NeedChagePassword == false {
		_, e = engine.Table(new(User)).ID(u.ID).Update(map[string]interface{}{"need_chage_password": 0})
	}
	if u.IsAdmin == false {
		_, e = engine.Table(new(User)).ID(u.ID).Update(map[string]interface{}{"is_admin": 0})
	}
	return
}
func (u *User) GetID() string {
	return u.ID
}

// CURD interface end
