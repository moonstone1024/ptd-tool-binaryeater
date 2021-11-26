package generator

type TD_UserAchievement struct {
	AchievementDate     string `json:"aDt"`
	Category            int    `json:"cat"`
	ResetDayOfWeek      int    `json:"rdw"`
	ConditionId         int    `json:"cnd"`
	GainReward          int    `json:"gRw"`
	AchievementID       string `json:"aId"`
	MasterName          string `json:"mn"`
	MasterAchievementId string `json:"maId"`
}
