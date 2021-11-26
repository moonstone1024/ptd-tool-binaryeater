package generator

type MD_Achievement struct {
	ID          string `json:"_ID"`
	TextData    string `json:"_TextData"`
	Category    int    `json:"_Category"`
	ConditionID int    `json:"_ConditionID"`
}
