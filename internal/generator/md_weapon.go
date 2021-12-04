package generator

type MD_Weapon struct {
	ID            string `json:"_ID"`
	Name          string `json:"_Name"`
	Rarity        int    `json:"_Rarity"`
	ItemType      int    `json:"_ItemType"`
	SkillContent1 string `json:"_SkillContent1"`
	SkillContent2 string `json:"_SkillContent2"`
	SkillContent3 string `json:"_SkillContent3"`
}
