package model

type UserStory struct {
	ID         string   `json:"id"`
	Priority   string   `json:"priority"`
	Risk       string   `json:"risk"`
	Points     int      `json:"points"`
	Dependency []string `json:"dependency"`
	Story      string   `json:"story"`
}

// interface start
func (s *UserStory) Update() (e error) {
	_, e = engine.ID(s.ID).Update(s)
	return
}

func (s *UserStory) GetID() string {
	return s.ID
}

// interface end
