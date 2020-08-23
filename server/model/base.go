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

// CURD iferface CURD
// By xorm bool bug modelStruct has a self update method
// CURD object is a pointer
type CURD interface {
	Update() error
	GetName() string
}

func Create(m CURD) (e error) {
	engine.Create(m)
	return
}
func Delete(m CURD) (count int, e error) {
	engine.Model(m).Count(&count).Delete(m)
	return
}
func Update(m CURD) (e error) {
	e = m.Update()
	return
}
func Get(m CURD) (e error) {
	engine.First(m)
	return
}
func GetByName(m CURD) (e error) {
	engine.Where("name = ?", m.GetName()).First(m)
	return
}
func GetAll(ms interface{}) (e error) {
	engine.Find(ms)
	return
}

func Preload(m CURD, preload string) (e error) {
	engine.Preload(preload).Find(m)
	return
}
