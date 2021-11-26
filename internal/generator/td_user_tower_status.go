package generator

type TD_UserTowerStatus struct {
	ResetData string                    `json:"rd"`
	Data      TD_UserTowerStatus_MyData `json:"dt"`
}

type TD_UserTowerStatus_MyData struct {
	SeasonID         string                                             `json:"sId"`
	QuestStageID     string                                             `json:"qsId"`
	LastQuestStageID string                                             `json:"lqsId"`
	QuestData        map[string]map[string]TD_UserTowerStatus_QuestData `json:"qd"`
	SR               []string                                           `json:"sr"`
}

type TD_UserTowerStatus_QuestData struct {
	StarCount   int `json:"sCnt"`
	MinimumTurn int `json:"mt"`
	MaxDamage   int `json:"md"`
}
