package generator

type TransData struct {
	User                      TD_User                                `json:"User"`
	UserCharacterContents     []TD_UserCharacterContents             `json:"UserCharacterContents"`
	UserAchievement           map[string]TD_UserAchievement          `json:"UserAchievement"`
	UserAchievementAccumulate UserAchievementAccumulateData          `json:"UserAchievementAccumulate"`
	UserEmblem                []interface{}                          `json:"UserEmblem"`
	UserFriend                []interface{}                          `json:"UserFriend"`
	UserQuest                 map[string]map[string]TD_UserQuestInfo `json:"UserQuest"`
	UserCostume               map[string]TD_User_Costume             `json:"UserCostume"`
	UserAccessory             []interface{}                          `json:"UserAccessory"`
	UserWeapon                map[string]TD_UserWeapon               `json:"UserWeapon"`
	UserLibrary               []TD_UserLibrary
	UserBoxGacha              []interface{}                 `json:"UserBoxGacha"`
	UserTowerStatus           map[string]TD_UserTowerStatus `json:"UserTowerStatus"`
	UserShootingModeItem      map[string][]string           `json:"UserShootingModeItem"`
}
