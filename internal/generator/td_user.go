package generator

type TD_User struct {
	FriendRequestId     int64  `json:"frId"`
	UserLv              int    `json:"lv"`
	UserExp             int    `json:"exp"`
	GameCoin            int    `json:"gCoin"`
	RealCoin            int    `json:"rCoin"`
	RealCoinFree        int    `json:"fCoin"`
	GachaFreeDrawDate   string `json:"gFDD"`
	LastLoginDate       string `json:"logD"`
	Stamina             int    `json:"stm"`
	StaminaMax          int    `json:"stmMx"`
	LastRecoveryDate    string `json:"recD"`
	LastConsumeDate     string `json:"lscD"`
	PlayQuestId         string `json:"pQId"`
	PlayQuestMasterName string `json:"pQNm"`
	PlayQuestTeamNumber string `json:"pQTNm"`
	PlayerQuestEnemies  string `json:"pqe"`
	PlayQuestSupporter  int    `json:"pqSpt"`
	PlayerQuestBuf      string `json:"pqBf"`
	PlayQuestContinue   int    `json:"pqCon"`
	PlayRankingID       string `json:"pRkId"`
	PlayQuestTempId     string `json:"pqTmp"`
	PlayedRankingEvents string `json:"pRkEv"`
	FriendMax           int    `json:"frMx"`
	// FrPnt is unknown field
	FrPnt                int    `json:"frPnt"`
	FriendLeasePoint     int    `json:"frLpt"`
	FriendLeasePointMax  int    `json:"frLmx"`
	CostumeInventory     int    `json:"csInv"`
	AccessoryInventory   int    `json:"acInv"`
	WeaponInventory      int    `json:"wpInv"`
	ComboLobinNum        int    `json:"cLogn"`
	SumLobinNum          int    `json:"sLogn"`
	NativeTagName        string `json:"ntNm"`
	CharacterExpResource string `json:"cExRs"`
	Lesson               string `json:"lesn"`
	CreateDate           string `json:"creD"`
	// LogDH contains unix time stamps as json encoded string. same as "logD"?
	LogDH    string `json:"logDH"`
	Tutorial string `json:"tuPnt"`
	// LrcuD contains some unix time stamp value.
	LrcuD                int64 `json:"lrcuD"`
	PlayQuestStaminaRate int   `json:"pqsr"`
	// BillA is unknown field
	BillA int `json:"billA"`
	// MBill is unknown field
	MBill int `json:"mBill"`
	// LrcgD contains some datetime string
	LrcgD string `json:"lrcgD"`
	// EndTimes seems to be a dictionary defined in TD_User_GachaEndTime.EndTimes
	EndTimes string `json:"gEd"`
	// TPasD contains some date time string
	TPasD string `json:"tPasD"`
	// SBE is unknown field
	SBE int `json:"sBE"`
	// PQTy is unknown field
	PQTy int `json:"pQTy"`
	// ChtP is unknown field
	ChtP             int    `json:"chtP"`
	CostumeVersion   int    `json:"cVer"`
	ConsumeVP0Counts string `json:"vp0"`
	// Pqvp0 is unknown field
	Pqvp0 bool `json:"pqvp0"`
	// ClVer seems to be idol class version
	ClVer                  int    `json:"clVer"`
	RealGoodsPointRateInfo string `json:"rgpri"`
	// PcbID contains whitespace
	PcbID string `json:"pcbId"`
	// PvpID contains whitespace
	PvpID string `json:"pvpId"`
	// EB is unknown field
	Eb                     int      `json:"eb"`
	UserId                 string   `json:"id"`
	CharacterPrivate       []string `json:"pr_ch"`
	TeamList               string   `json:"team"`
	CharacterStatusList    string   `json:"csts"`
	GachaEventCount        string   `json:"gEC"`
	GachaBonusCount        string   `json:"gBCnt"`
	EquipmentAvatarsList   string   `json:"eqAv"`
	OwnCharacterIds        string   `json:"crIds"`
	FavoriteCharacterId    string   `json:"fCrId"`
	PushNotificationToken  string   `json:"push"`
	ClientDeviceOS         string   `json:"cdos"`
	ClientDeviceName       string   `json:"cdnm"`
	PurchaseLimitCounts    string   `json:"pLmt"`
	StaminaItem            string   `json:"irs"`
	AreaReleaseItem        string   `json:"ari"`
	UserGachaTicket        string   `json:"gt"`
	BoostItem              string   `json:"bst"`
	TradeItem              string   `json:"tt"`
	GiftItem               string   `json:"gfti"`
	RecaptureItem          string   `json:"rcpi"`
	WeaponMaterial         string   `json:"wm"`
	WeaponLimitBreakItem   string   `json:"wlb"`
	CostumeLimitBreakItem  string   `json:"cslb"`
	CostumeAwakeItem       string   `json:"csAw"`
	QuestTicketItem        string   `json:"qt"`
	CostumeTradeDollsPoint string   `json:"ctp"`
	CostumeStone           string   `json:"csSt"`
	WeaponStone            string   `json:"wpSt"`
	EventVotePoint         string   `json:"eVvp"`
	AreaReleaseItemUseDate string   `json:"ariud"`
	ItmBf                  string   `json:"itmBf"`
	LoginBonus             string   `json:"lbNs"`
	ItemTradeCount         string   `json:"itc"`
	ItemBuyCount           string   `json:"ibc"`
	MainTitle              string   `json:"tiId"`
	BirthYM                string   `json:"bthYM"`
	TiAs                   int      `json:"tiAs"`
	Name                   string   `json:"name"`
	Profile                string   `json:"prof"`
}
