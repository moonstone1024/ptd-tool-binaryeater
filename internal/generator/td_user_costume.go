package generator

type TD_User_Costume struct {
	CostumeUId      string        `json:"csUId"`
	LimitBreakCount int           `json:"lbCnt"`
	IsSalable       bool          `json:"bSale"`
	AUID            []interface{} `json:"aUId"`
	CostumeLevel    int           `json:"lv"`
	IsLock          bool          `json:"bL"`
	GrowthPoint     int           `json:"gPnt"`
	CostumeId       string        `json:"csId"`
	CreateDate      string        `json:"creD"`
	DFLT            int           `json:"dflt"`
	MaxLevel        int           `json:"mxlv"`
	UpdateDate      string        `json:"updD"`
	AwakeLevel      int           `json:"awlv"`
	CostumeExp      int           `json:"exp"`
	EXS             []interface{} `json:"exs"`
}
