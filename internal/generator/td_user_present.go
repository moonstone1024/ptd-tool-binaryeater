package generator

type TD_UserPresent struct {
	Type             string `json:"typ"`
	PresentId        string `json:"prId"`
	ReceiptBeginDate string `json:"rBD"`
	ReceiptEndDate   string `json:"rED"`
	ReceiptDate      string `json:"rD"`
	ItemType         int    `json:"iType"`
	ItemId           string `json:"iId"`
	ItemAmount       int    `json:"iAmt"`
	Message          string `json:"msg"`
}
