package generator

type LoginResponseParameter struct {
	ClassUpdate           map[string]int             `json:"clUp"`
	TowerRankingRank      int                        `json:"trRnk"`
	SessionId             string                     `json:"sid"`
	LoginTime             int64                      `json:"logt"`
	TachimotoInsertIndex  int                        `json:"tii"`
	BridgeToken           string                     `json:"bt"`
	PresentResponse       TD_UserPresentResponseData `json:"pres"`
	AssetBundleDomainName string                     `json:"asbDm"`
	AssetBundleBacketName string                     `json:"asbBk"`
	AssetBundleVersion    int                        `json:"asbVr"`
	AssetBundleVersion2   int                        `json:"asbVr2"`
	RealCoin              int                        `json:"rCoin"`
	RealCoinFree          int                        `json:"fCoin"`
}
