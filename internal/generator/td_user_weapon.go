package generator

type TD_UserWeapon struct {
	UCR                  string  `json:"uCR"`
	WeaponUId            string  `json:"wpUId"`
	SkillId3             string  `json:"skId3"`
	SkillId2             string  `json:"skId2"`
	SkillId1             string  `json:"skId1"`
	WeaponLevel          int     `json:"lv"`
	IsLock               bool    `json:"bL"`
	IsEquipment          int     `json:"bEq"`
	WeaponId             string  `json:"wpId"`
	CreateDate           string  `json:"creD"`
	Dflt                 int     `json:"dflt"`
	LimitBreakCount      int     `json:"lBCnt"`
	MaxLevel             int     `json:"mxlv"`
	SkillId7             string  `json:"skId7"`
	LimitBreakSkillTable *string `json:"lbST"`
	UpdateDate           string  `json:"updD"`
	SkillId6             string  `json:"skId6"`
	SkillId5             string  `json:"skId5"`
	SkillId4             string  `json:"skId4"`
	WeaponExp            int     `json:"exp"`
	SkillId8             string  `json:"skId8"`
}
