package generator

type TD_UserQuestInfo struct {
	ClearMissions       []string `json:"ms"`
	TotalClearCount     int      `json:"tcc"`
	TotalChallengeCount int      `json:"chl"`
	Lcd                 string   `json:"lcd"`
}
