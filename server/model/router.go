package model

type Router struct {
	ID        string `json:"id" xorm:"pk"`
	Title     string `json:"title" xorm:"varchar(64)"`
	Icon      string `json:"icon" xorm:"varchar(64)"`
	Order     int    `json:"order" xorm:"int"`
	Path      string `json:"path" xorm:"varchar(64)"`
	Exact     bool   `json:"exact" xorm:"TINYINT(1)"`
	Component string `json:"component" xorm:"varchar(64)"`
}

// interface start
func (r *Router) Update() (e error) {
	_, e = engine.ID(r.ID).Update(r)
	return
}

func (r *Router) GetID() string {
	return r.ID
}

// interface end
