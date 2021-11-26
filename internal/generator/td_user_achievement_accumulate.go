package generator

type TD_UserAchievementAccumulate struct {
	AccumulateId                        string                                   `json:"acmId"`
	AchievementCategoryCount            map[string]int                           `json:"aCTC"`
	BattleAchievementAccumulate         map[string]int                           `json:"bAAc"`
	BeginnerTrophyClear                 map[string]int                           `json:"btCl"`
	BoxGachaInformationCount            map[string]int                           `json:"bgIC"`
	Category                            string                                   `json:"cat"`
	CharacterStatus                     map[string]TD_AchievementCharacterStatus `json:"csts"`
	ConsumeStaminaTotal                 int                                      `json:"cSTT"`
	CostumeSellCountTotal               int                                      `json:"cSCT"`
	EventBattleConditionCount           map[string]int                           `json:"eBCC"`
	EventQuestClearCount                map[string]map[string]int                `json:"eQCC"`
	GainExpResourceTotal                map[string]int                           `json:"gERT"`
	GainFriendPointTotal                int                                      `json:"gFPT"`
	GainGameCoinTotal                   int                                      `json:"gGCT"`
	GainItemRarityCount                 map[string]map[int]int                   `json:"gIRC"`
	GainMaxStatusWeaponCount            map[int]map[int]int                      `json:"gMSWC"`
	MasterName                          string                                   `json:"mNm"`
	MaxWeaponLimitBreakCount            map[int]map[int]int                      `json:"mWLBC"`
	QuestClearCountTotal                int                                      `json:"qCCT"`
	QuestClearRarityCountTotal          map[string]map[int]int                   `json:"qCRT"`
	QuestMissionClearCountTotal         int                                      `json:"qMCT"`
	QuestRentalSupportCount             int                                      `json:"qRPC"`
	QuestResultCountTypeDifficultyTotal map[string]map[string]map[string]int     `json:"qRCTD"`
	QuestResultCountTypeTotal           map[string]map[string]int                `json:"qRCTT"`
	QuestSupportCountTotal              int                                      `json:"qSPC"`
	RaidLastAttackCount                 int                                      `json:"rLAC"`
	RaidOwnerCount                      int                                      `json:"rOWC"`
	RaidQuestClearCount                 map[string]map[string]int                `json:"rQCC"`
	SellGameCoinTotal                   int                                      `json:"sGCT"`
	SpecifiedItemLimitBreakCount        map[string]int                           `json:"tstrC"`
	StrengthenCount                     map[string]int                           `json:"strC"`
	TradeItemTradeCountTotal            map[string]int                           `json:"tITCT"`
	UserId                              string                                   `json:"id"`
	WeaponStrengthenCount               map[string]map[string]int                `json:"wStrC"`
}
