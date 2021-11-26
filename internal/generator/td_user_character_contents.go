package generator

type TD_UserCharacterContents struct {
	FeelSubScenario      map[string]int           `json:"fSn"`
	Memories             map[string]TD_MemoryData `json:"mem"`
	CostumeSubScenario   map[string]int           `json:"cSn"`
	CharacterId          string                   `json:"cId"`
	RecaptureSubScenario map[string]int           `json:"rSn"`
}
