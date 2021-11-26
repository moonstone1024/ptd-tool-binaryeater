package generator

type TD_User_CharacterStatus struct {
	Biorhythm int `json:"bio"`
	// ExpA is yellow
	ExpA int `json:"ExA"`
	// ExpB is magenta
	ExpB int `json:"ExB"`
	// ExpC is cyan
	ExpC int `json:"ExC"`
	// ExpD is purple
	ExpD      int           `json:"ExD"`
	IdolClass string        `json:"cls"`
	LevelA    int           `json:"LvA"`
	LevelB    int           `json:"LvB"`
	LevelC    int           `json:"LvC"`
	LevelD    int           `json:"LvD"`
	Cmp       []interface{} `json:"cmp"`
	MaxLv     int           `json:"mLv"`
}
