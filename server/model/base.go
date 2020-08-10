package model

import (
	"github.com/jinzhu/gorm"
	// import sqlit3 module
	_ "github.com/mattn/go-sqlite3"
)

var engine *gorm.DB

func init() {
	var err error

	// TODO: get with env viper
	driverName := "sqlite3"
	dataSourceName := "./db/dev.db"
	engine, err = gorm.Open(driverName, dataSourceName)
	if err != nil {
		panic("failed to connect database")
	}

	// TODO: uncache seting
	// Use a function
	// Code like :
	// engine.MapCacher(&user,nil)
	// the nil value

	// show the sql debug
	engine.LogMode(true) // TODO: viper dev true other false

	// call migrate
	Migrate(engine)
}

// Migrate data schema
func Migrate(x *gorm.DB) error {
	// some code
	// TODO: feat

	// temp demo sheet
	x.AutoMigrate(&User{})
	return nil
}

// DB iferface CURD
// By xorm bool bug modelStruct has a self update method
// CURD object is a pointer
type DB interface {
	Update() error
	GetID() string
}

// func Create(m CURD) (e error) {
// 	_, e = engine.Insert(m)
// 	return
// }
// func Delete(m CURD) (e error) {
// 	_, e = engine.ID(m.GetID()).Delete(m)
// 	return
// }
// func Update(m CURD) (e error) {
// 	e = m.Update()
// 	return
// }
// func Get(m CURD) (e error) {
// 	_, e = engine.Get(m)
// 	return
// }
// func GetAll(ms interface{}) (e error) {
// 	e = engine.Find(ms)
// 	return
// }
