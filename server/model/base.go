package model

import (
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
	"xorm.io/xorm/caches"
)

var engine *xorm.Engine

func init() {
	var err error

	// TODO: get with env viper
	driverName := "sqlite3"
	dataSourceName := "./db/dev.db"
	engine, err = xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		log.Println("init database engine error:", err)
	}
	// TODO: i18n setting
	local := "Asia/Shanghai"
	engine.TZLocation, err = time.LoadLocation(local)

	// cache
	cacheItermMax := 1000 // TODO: viper
	cacher := caches.NewLRUCacher(caches.NewMemoryStore(), cacheItermMax)
	engine.SetDefaultCacher(cacher)

	// TODO: uncache seting
	// Use a function
	// Code like :
	// engine.MapCacher(&user,nil)
	// the nil value

	// show the sql debug
	engine.ShowSQL(true) // TODO: viper dev true other false

	// call migrate
	Migrate(engine)
}

// migrate data schema
func Migrate(x *xorm.Engine) error {
	// some code
	// TODO: feat

	// temp demo sheet
	err := x.Sync(new(User), new(Router))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// iferface CURD
// By xorm bool bug modelStruct has a self update method
// CURD object is a pointer
type CURD interface {
	Update() error
	GetID() string
}

func Create(m CURD) (e error) {
	_, e = engine.Insert(m)
	return
}
func Delete(m CURD) (e error) {
	_, e = engine.ID(m.GetID()).Delete(m)
	return
}
func Update(m CURD) (e error) {
	e = m.Update()
	return
}
func Get(m CURD) (e error) {
	_, e = engine.Get(m)
	return
}
func GetAll(ms interface{}) (e error) {
	e = engine.Find(ms)
	return
}
