package generator

type TD_AchievementCharacterStatus struct {
	AddFeel1                 int `json:"af1"`
	AddFeel2                 int `json:"af2"`
	AddFeel3                 int `json:"af3"`
	AddFeel4                 int `json:"af4"`
	AvatarChangeCount        int `json:"avc"`
	BattleUseCount           int `json:"buc"`
	ChangeHome               int `json:"ch"`
	DislikeAvatarChangeCount int `json:"dav"`
	FavoriteTime             int `json:"fvt"`
	GiftReturn               int `json:"grc"`
	GiveGiftCount            int `json:"ggc"`
	LessonGoodCount          int `json:"lgc"`
	LessonSuccessCount       int `json:"lsc"`
	LikeAvatarChangeCount    int `json:"lav"`
	TapCount                 int `json:"tap"`
	VisitRoomCount           int `json:"vis"`
	VoteFeelCount            int `json:"vfc"`
	VoteLevelUpCount         int `json:"vlu"`
}
