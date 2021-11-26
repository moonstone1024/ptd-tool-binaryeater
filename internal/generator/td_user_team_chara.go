package generator

type TD_User_Team_Chara struct {
	// MCrCs is costume id
	MCrCs string `json:"mCrCs"`
	// MCrId is character id
	MCrId string `json:"mCrId"`
	// MCrWp is weapon id
	MCrWp string `json:"mCrWp"`
	// SCs1 is sub costume 1
	SCs1 string `json:"sCs1,omitempty"`
	// SCs1 is sub costume 2
	SCs2 string `json:"sCs2,omitempty"`
	// SCs1 is sub costume 3
	SCs3 string `json:"sCs3,omitempty"`
}
