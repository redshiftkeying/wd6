package model

import (
	"crypto/md5"
	"fmt"
)

// Auth function basic authentication name and passwd
func (u *User) Auth(name string, passwd string) bool {
	engine.Where("name=?", name).First(u)
	if u.Password == passwd {
		return true
	}
	return false
}

// Encryption functon
func (u *User) Encryption(pwd string) {
	sal := "lslslsklsksksgoykekclkg#" // TODO:viper function (platform_id)
	str := pwd + sal
	data := []byte(str)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	u.Password = md5str1
}
