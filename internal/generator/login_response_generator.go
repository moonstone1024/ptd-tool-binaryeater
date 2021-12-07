package generator

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path"
	"sort"
	"strings"
	"time"
)

type loginResponseGenerator struct {
	decodedMDDir string
	cfg          LoginResponseGeneratorConfig
}

type LoginResponseGeneratorConfig struct {
	ServerTime              CustomTime                   `json:"serverTime"`
	CostumeUIDSalt          string                       `json:"costumeUIDSalt"`
	WeaponUIDSalt           string                       `json:"weaponUIDSalt"`
	EnableDebugStockWeapons bool                         `json:"enableDebugStockWeapons"`
	EquipmentAvatarsList    map[string]map[string]string `json:"equipmentAvatarsList"`
	Teams                   map[string]TD_User_Team      `json:"teams"`
	IDolClasses             map[string]string            `json:"idolClasses"`
	FavoriteCharacterId     string                       `json:"favoriteCharacterId"`
}

func GenearateLogin(decodedMDDir string, cfg LoginResponseGeneratorConfig) *LoginResponse {
	l := loginResponseGenerator{
		decodedMDDir: decodedMDDir,
		cfg:          cfg,
	}

	return l.generate()
}

func (l *loginResponseGenerator) generate() *LoginResponse {
	inCostumes := l.loadMDCostumes()
	costumes := l.createCostumeUIDToCostumeMap(inCostumes)
	inWeapons := l.loadWeapons()
	weapons := l.createWeaponUIDToWeaponMap(inWeapons)
	avatarIDs := l.loadAvatarIDs()
	subScenarios := l.loadSubScenarios()
	inMemoryLocks := l.loadCharacterMemoryLocks()
	inCostumeStones := l.loadCostumeStones()
	inWeaponStones := l.loadWeaponStones()
	serverTimeText := l.cfg.ServerTime.Format(serverTimeFormat)
	serverTimeStamp := l.cfg.ServerTime.Unix()
	resp := &LoginResponse{
		ServerTime: serverTimeText,
		ResponseParameter: LoginResponseParameter{
			ClassUpdate: map[string]int{
				"o": 11,
				"n": 11,
			},
			TowerRankingRank:     -1,
			SessionId:            "dummy-session-id",
			LoginTime:            serverTimeStamp,
			TachimotoInsertIndex: 3,
			BridgeToken:          "dummy-bridge-token",
			PresentResponse: TD_UserPresentResponseData{
				NotReceived: []TD_UserPresent{},
				Received:    []TD_UserPresent{},
			},
			AssetBundleDomainName: "cache.example.com",
			AssetBundleBacketName: "non-existent-cdn-bucket",
			AssetBundleVersion:    95500,
			AssetBundleVersion2:   129900,
			RealCoin:              0,
			RealCoinFree:          10000,
		},
		TransData: TransData{
			User: TD_User{
				FriendRequestId:        1,
				UserLv:                 1,
				UserExp:                0,
				GameCoin:               0,
				RealCoin:               0,
				RealCoinFree:           0,
				GachaFreeDrawDate:      "2000-01-01 00:00:00",
				LastLoginDate:          serverTimeText,
				Stamina:                100,
				StaminaMax:             100,
				LastRecoveryDate:       serverTimeText,
				LastConsumeDate:        serverTimeText,
				PlayQuestId:            "",
				PlayQuestMasterName:    " ",
				PlayQuestTeamNumber:    " ",
				PlayerQuest:            "[]",
				PlayQuestSupporter:     -1,
				PlayerQuestBuf:         "[]",
				PlayQuestContinue:      0,
				PlayRankingID:          " ",
				PlayQuestTempId:        " ",
				PlayedRankingEvents:    "{}",
				FriendMax:              0,
				FrPnt:                  0,
				FriendLeasePoint:       0,
				FriendLeasePointMax:    0,
				CostumeInventory:       2000,
				AccessoryInventory:     2000,
				WeaponInventory:        2000,
				ComboLobinNum:          0,
				SumLobinNum:            0,
				NativeTagName:          "dummy-native-tag",
				CharacterExpResource:   l.getCharacterExpResourceAsString(),
				Lesson:                 l.getUserLessonMapAsString(),
				CreateDate:             serverTimeText,
				LogDH:                  "[]",
				Tutorial:               "{}",
				LrcuD:                  serverTimeStamp,
				PlayQuestStaminaRate:   1,
				BillA:                  0,
				MBill:                  0,
				LrcgD:                  serverTimeText,
				EndTimes:               "{}",
				TPasD:                  serverTimeText,
				SBE:                    0,
				PQTy:                   0,
				ChtP:                   0,
				CostumeVersion:         210,
				ConsumeVP0Counts:       "{\"2\":1}",
				Pqvp0:                  false,
				ClVer:                  11,
				RealGoodsPointRateInfo: "{}",
				PcbID:                  " ",
				PvpID:                  " ",
				Eb:                     0,
				UserId:                 "dummy-user-id",
				CharacterPrivate:       []string{""},
				TeamList:               l.getTeamListAsString(costumes, weapons),
				CharacterStatusList:    l.getCharacterStatusListAsString(),
				GachaEventCount:        "{}",
				GachaBonusCount:        "{}",
				EquipmentAvatarsList:   l.getEquipmentAvatarListAsString(),
				OwnCharacterIds:        "c001,c002,c003,c004,c005,c006,c007,c008,c009,c010,c011,c012,c301,c302,c303,c304",
				FavoriteCharacterId:    l.cfg.FavoriteCharacterId,
				PushNotificationToken:  " ",
				ClientDeviceOS:         " ",
				ClientDeviceName:       " ",
				PurchaseLimitCounts:    " ",
				StaminaItem:            "{}",
				AreaReleaseItem:        "{}",
				UserGachaTicket:        "{}",
				BoostItem:              "[]",
				TradeItem:              "{}",
				GiftItem:               "{}",
				RecaptureItem:          l.getRecaptureItemsAsString(inMemoryLocks),
				WeaponMaterial:         "{}",
				WeaponLimitBreakItem:   "{}",
				CostumeLimitBreakItem:  "[]",
				CostumeAwakeItem:       "[]",
				QuestTicketItem:        "{}",
				CostumeTradeDollsPoint: "{}",
				CostumeStone:           "{}",
				WeaponStone:            "{}",
				EventVotePoint:         "[]",
				AreaReleaseItemUseDate: "{}",
				ItmBf:                  "[]",
				LoginBonus:             "{\"3\":{\"LBCB0001\":0,\"LBCB0002\":0},\"2\":{\"LBWC0001\":0,\"LBWC0002\":0},\"1\":{\"LBN001\":1}}",
				ItemTradeCount:         "{}",
				ItemBuyCount:           "[]",
				MainTitle:              " ",
				BirthYM:                "1970-01",
				TiAs:                   0,
				Name:                   "Master",
				Profile:                "Hi",
			},
			UserCharacterContents: l.getUserCharacterContents(subScenarios, inMemoryLocks),
			UserAchievement:       l.getUserAchievement(),
			UserAchievementAccumulate: UserAchievementAccumulateData{
				Achievement: map[string]TD_UserAchievementAccumulate{},
			},
			UserEmblem:    []interface{}{},
			UserFriend:    []interface{}{},
			UserQuest:     l.getUserQuest(),
			UserCostume:   costumes,
			UserAccessory: []interface{}{},
			UserWeapon:    weapons,
			UserLibrary: l.getUserLibrayList(avatarIDs,
				inCostumes,
				inWeapons,
				inCostumeStones,
				inWeaponStones),
			UserBoxGacha:         []interface{}{},
			UserTowerStatus:      l.getUserTowerStatus(),
			UserShootingModeItem: l.getUserShootingModeItems(),
		},
	}
	return resp
}

func (l *loginResponseGenerator) loadMDJson(filename string, v interface{}) {
	filePath := path.Join(l.decodedMDDir, filename)
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(bytes, v)
	if err != nil {
		panic(err)
	}
}

func (l *loginResponseGenerator) getCharacterExpResourceAsString() string {
	r := CharacterExpResource{
		A: 999999999,
		B: 999999999,
		C: 999999999,
		D: 999999999,
	}
	bytes, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func (l *loginResponseGenerator) getUserLessonMapAsString() string {
	lesson := map[string]TD_User_Lesson{
		"c001": {
			LessonID: "",
		},
		"c002": {
			LessonID: "",
		},
		"c003": {
			LessonID: "",
		},
		"c004": {
			LessonID: "",
		},
		"c005": {
			LessonID: "",
		},
		"c006": {
			LessonID: "",
		},
		"c007": {
			LessonID: "",
		},
		"c008": {
			LessonID: "",
		},
		"c009": {
			LessonID: "",
		},
	}
	bytes, err := json.Marshal(lesson)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func (l *loginResponseGenerator) loadMDCostumes() []MD_Costume {
	inCostumes := make([]MD_Costume, 0)
	l.loadMDJson("MD_Costume.json", &inCostumes)

	return inCostumes
}

func (l *loginResponseGenerator) createCostumeUIDToCostumeMap(inCostumes []MD_Costume) map[string]TD_User_Costume {
	h, err := CreateNewHashId(l.cfg.CostumeUIDSalt)
	if err != nil {
		panic(err)
	}

	index := 0
	costumes := map[string]TD_User_Costume{}
	t := l.cfg.ServerTime.Time
	for _, costume := range inCostumes {
		costumeUID, err := h.Encode([]int{index})
		index++
		if err != nil {
			panic(err)
		}
		if costume.ItemKindID != 0 {
			continue
		}
		rarity := costumeRarities[costume.Rarity]
		costumes[costumeUID] = TD_User_Costume{
			CostumeUId:      costumeUID,
			LimitBreakCount: 0,
			IsSalable:       true,
			AUID:            []interface{}{},
			CostumeLevel:    rarity.MaxLevel,
			IsLock:          true,
			GrowthPoint:     0,
			CostumeId:       costume.ID,
			CreateDate:      t.Format(serverTimeFormat),
			DFLT:            0,
			MaxLevel:        rarity.MaxLevel,
			UpdateDate:      t.Format(serverTimeFormat),
			AwakeLevel:      0,
			CostumeExp:      rarity.MaxExp,
			EXS:             []interface{}{},
		}
		t = t.Add(time.Duration(1) * time.Second)
	}

	return costumes
}

func (l *loginResponseGenerator) loadWeapons() []MD_Weapon {
	inWeapons := make([]MD_Weapon, 0)
	l.loadMDJson("MD_Weapon.json", &inWeapons)

	return inWeapons
}

func (l *loginResponseGenerator) createWeaponUIDToWeaponMap(inWeapons []MD_Weapon) map[string]TD_UserWeapon {
	h, err := CreateNewHashId(l.cfg.WeaponUIDSalt)
	if err != nil {
		panic(err)
	}
	inWeaponSkillContents := make([]MD_WeaponSkillContent, 0)
	l.loadMDJson("MD_WeaponSkillContent.json", &inWeaponSkillContents)
	weaponSkillContenIDMap := map[string]MD_WeaponSkillContent{}
	for _, v := range inWeaponSkillContents {
		weaponSkillContenIDMap[v.ID] = v
	}

	index := 0
	weapons := map[string]TD_UserWeapon{}
	t := l.cfg.ServerTime.Time
	for _, weapon := range inWeapons {
		weaponUID, err := h.Encode([]int{index})
		index++
		if err != nil {
			panic(err)
		}
		if weapon.ItemType != 1 {
			continue
		}
		if !l.cfg.EnableDebugStockWeapons && (strings.Contains(weapon.Name, "デバッグ") || strings.Contains(weapon.Name, "ストック")) {
			continue
		}
		skillID1 := " "
		skillID2 := " "
		skillID3 := " "

		if v, ok := weaponSkillContenIDMap[weapon.SkillContent1]; ok {
			skillID1 = v.SkillID1
		}
		if v, ok := weaponSkillContenIDMap[weapon.SkillContent2]; ok {
			skillID2 = v.SkillID1
		}
		if v, ok := weaponSkillContenIDMap[weapon.SkillContent3]; ok {
			skillID3 = v.SkillID1
		}

		weapons[weaponUID] = TD_UserWeapon{
			UCR:                  " ",
			WeaponUId:            weaponUID,
			WeaponLevel:          20,
			IsLock:               false,
			IsEquipment:          0,
			WeaponId:             weapon.ID,
			CreateDate:           t.Format(serverTimeFormat),
			Dflt:                 0,
			LimitBreakCount:      0,
			MaxLevel:             20,
			LimitBreakSkillTable: nil,
			UpdateDate:           t.Format(serverTimeFormat),
			WeaponExp:            weaponRarities[weapon.Rarity].MaxExp,
			SkillId1:             skillID1,
			SkillId2:             skillID2,
			SkillId3:             skillID3,
			SkillId4:             " ",
			SkillId5:             " ",
			SkillId6:             " ",
			SkillId7:             " ",
			SkillId8:             " ",
		}
		t = t.Add(time.Duration(1) * time.Second)
	}

	return weapons
}

func (l *loginResponseGenerator) loadAvatarIDs() []string {
	inAvatars := make([]MD_Avatar, 0)
	l.loadMDJson("MD_Avatar.json", &inAvatars)

	avatarLen := len(inAvatars)
	avatars := make([]string, avatarLen)
	index := 0
	for _, avatar := range inAvatars {
		avatars[index] = avatar.ID
		index++
	}

	return avatars
}

func (l *loginResponseGenerator) getTeamListAsString(costumeUIDToCostumeMap map[string]TD_User_Costume,
	weaponUIDToWeaponMap map[string]TD_UserWeapon) string {
	costumeIDToCostumeMap := l.createCostumeIdToCostumeMap(costumeUIDToCostumeMap)
	weaponIDToWeaponMap := l.createWeaponIDToWeaponMap(weaponUIDToWeaponMap)
	teams := map[string]TD_User_Team{}
	for id, team := range l.cfg.Teams {
		l.mapTeamCharaFields(&team.Order1st, costumeIDToCostumeMap, weaponIDToWeaponMap)
		l.mapTeamCharaFields(&team.Order2nd, costumeIDToCostumeMap, weaponIDToWeaponMap)
		l.mapTeamCharaFields(&team.Order3rd, costumeIDToCostumeMap, weaponIDToWeaponMap)
		teams[id] = team
	}
	bytes, err := json.Marshal(teams)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func (l *loginResponseGenerator) mapTeamCharaFields(chara *TD_User_Team_Chara,
	costumeIDToCostumeMap map[string]TD_User_Costume,
	weaponIDToWeaponMap map[string]TD_UserWeapon) {
	chara.MCrCs = costumeIDToCostumeMap[chara.MCrCs].CostumeUId
	chara.MCrWp = weaponIDToWeaponMap[chara.MCrWp].WeaponUId
}

func (l *loginResponseGenerator) createCostumeIdToCostumeMap(costumeUIDToCostumeMap map[string]TD_User_Costume) map[string]TD_User_Costume {
	m := map[string]TD_User_Costume{}
	for _, costume := range costumeUIDToCostumeMap {
		m[costume.CostumeId] = costume
	}

	return m
}

func (l *loginResponseGenerator) createWeaponIDToWeaponMap(costumeUIDToCostumeMap map[string]TD_UserWeapon) map[string]TD_UserWeapon {
	m := map[string]TD_UserWeapon{}
	for _, weapon := range costumeUIDToCostumeMap {
		m[weapon.WeaponId] = weapon
	}

	return m
}

func (l *loginResponseGenerator) getCharacterStatusListAsString() string {
	lv60FeelAmount := 8226000
	lv70FeelAmount := 17851600

	lv40FeelAmount := 2184400
	lv30FeelAmount := 699600

	statuses := map[string]TD_User_CharacterStatus{
		"c001": {
			Biorhythm: 80,
			ExpA:      lv60FeelAmount,
			ExpB:      lv70FeelAmount,
			ExpC:      lv60FeelAmount,
			ExpD:      lv70FeelAmount,
			IdolClass: l.cfg.IDolClasses["c001"],
			LevelA:    60,
			LevelB:    70,
			LevelC:    60,
			LevelD:    70,
			Cmp:       []interface{}{},
			MaxLv:     100,
		},
		"c002": {
			Biorhythm: 80,
			ExpA:      lv60FeelAmount,
			ExpB:      lv60FeelAmount,
			ExpC:      lv70FeelAmount,
			ExpD:      lv70FeelAmount,
			IdolClass: l.cfg.IDolClasses["c002"],
			LevelA:    60,
			LevelB:    60,
			LevelC:    70,
			LevelD:    70,
			Cmp:       []interface{}{},
			MaxLv:     100,
		},
		"c003": {
			Biorhythm: 80,
			ExpA:      lv70FeelAmount,
			ExpB:      lv60FeelAmount,
			ExpC:      lv60FeelAmount,
			ExpD:      lv70FeelAmount,
			IdolClass: l.cfg.IDolClasses["c003"],
			LevelA:    70,
			LevelB:    60,
			LevelC:    60,
			LevelD:    70,
			Cmp:       []interface{}{},
			MaxLv:     100,
		},
		"c004": {
			Biorhythm: 80,
			ExpA:      lv60FeelAmount,
			ExpB:      lv60FeelAmount,
			ExpC:      lv70FeelAmount,
			ExpD:      lv70FeelAmount,
			IdolClass: l.cfg.IDolClasses["c004"],
			LevelA:    60,
			LevelB:    60,
			LevelC:    70,
			LevelD:    70,
			Cmp:       []interface{}{},
			MaxLv:     100,
		},
		"c005": {
			Biorhythm: 80,
			ExpA:      lv60FeelAmount,
			ExpB:      lv70FeelAmount,
			ExpC:      lv60FeelAmount,
			ExpD:      lv70FeelAmount,
			IdolClass: l.cfg.IDolClasses["c005"],
			LevelA:    60,
			LevelB:    70,
			LevelC:    60,
			LevelD:    70,
			Cmp:       []interface{}{},
			MaxLv:     100,
		},
		"c006": {
			Biorhythm: 80,
			ExpA:      lv70FeelAmount,
			ExpB:      lv60FeelAmount,
			ExpC:      lv60FeelAmount,
			ExpD:      lv70FeelAmount,
			IdolClass: l.cfg.IDolClasses["c006"],
			LevelA:    70,
			LevelB:    60,
			LevelC:    60,
			LevelD:    70,
			Cmp:       []interface{}{},
			MaxLv:     100,
		},
		"c007": {
			Biorhythm: 80,
			ExpA:      lv70FeelAmount,
			ExpB:      lv60FeelAmount,
			ExpC:      lv60FeelAmount,
			ExpD:      lv70FeelAmount,
			IdolClass: l.cfg.IDolClasses["c007"],
			LevelA:    70,
			LevelB:    60,
			LevelC:    60,
			LevelD:    70,
			Cmp:       []interface{}{},
			MaxLv:     100,
		},
		"c008": {
			Biorhythm: 80,
			ExpA:      lv60FeelAmount,
			ExpB:      lv60FeelAmount,
			ExpC:      lv70FeelAmount,
			ExpD:      lv70FeelAmount,
			IdolClass: l.cfg.IDolClasses["c008"],
			LevelA:    60,
			LevelB:    60,
			LevelC:    70,
			LevelD:    70,
			Cmp:       []interface{}{},
			MaxLv:     100,
		},
		"c009": {
			Biorhythm: 80,
			ExpA:      lv60FeelAmount,
			ExpB:      lv70FeelAmount,
			ExpC:      lv60FeelAmount,
			ExpD:      lv70FeelAmount,
			IdolClass: l.cfg.IDolClasses["c009"],
			LevelA:    60,
			LevelB:    70,
			LevelC:    60,
			LevelD:    70,
			Cmp:       []interface{}{},
			MaxLv:     100,
		},
		"c010": {
			Biorhythm: 80,
			ExpA:      lv30FeelAmount,
			ExpB:      lv40FeelAmount,
			ExpC:      lv30FeelAmount,
			ExpD:      lv40FeelAmount,
			IdolClass: l.cfg.IDolClasses["c010"],
			LevelA:    30,
			LevelB:    40,
			LevelC:    30,
			LevelD:    40,
			Cmp:       []interface{}{},
			MaxLv:     100,
		},
		"c011": {
			Biorhythm: 80,
			ExpA:      lv40FeelAmount,
			ExpB:      lv30FeelAmount,
			ExpC:      lv30FeelAmount,
			ExpD:      lv40FeelAmount,
			IdolClass: l.cfg.IDolClasses["c011"],
			LevelA:    40,
			LevelB:    30,
			LevelC:    30,
			LevelD:    40,
			Cmp:       []interface{}{},
			MaxLv:     100,
		},
		"c012": {
			Biorhythm: 80,
			ExpA:      lv30FeelAmount,
			ExpB:      lv30FeelAmount,
			ExpC:      lv40FeelAmount,
			ExpD:      lv40FeelAmount,
			IdolClass: l.cfg.IDolClasses["c012"],
			LevelA:    30,
			LevelB:    30,
			LevelC:    40,
			LevelD:    40,
			Cmp:       []interface{}{},
			MaxLv:     100,
		},
		"c301": {
			Biorhythm: 80,
			ExpA:      0,
			ExpB:      0,
			ExpC:      0,
			ExpD:      0,
			IdolClass: "MD_IDOLCLASS_c301_0001",
			LevelA:    1,
			LevelB:    1,
			LevelC:    1,
			LevelD:    1,
			Cmp:       []interface{}{},
			MaxLv:     100,
		},
		"c302": {
			Biorhythm: 80,
			ExpA:      0,
			ExpB:      0,
			ExpC:      0,
			ExpD:      0,
			IdolClass: "MD_IDOLCLASS_c302_0001",
			LevelA:    1,
			LevelB:    1,
			LevelC:    1,
			LevelD:    1,
			Cmp:       []interface{}{},
			MaxLv:     100,
		},
		"c303": {
			Biorhythm: 80,
			ExpA:      0,
			ExpB:      0,
			ExpC:      0,
			ExpD:      0,
			IdolClass: "MD_IDOLCLASS_c303_0001",
			LevelA:    1,
			LevelB:    1,
			LevelC:    1,
			LevelD:    1,
			Cmp:       []interface{}{},
			MaxLv:     100,
		},
		"c304": {
			Biorhythm: 80,
			ExpA:      0,
			ExpB:      0,
			ExpC:      0,
			ExpD:      0,
			IdolClass: "MD_IDOLCLASS_c304_0001",
			LevelA:    1,
			LevelB:    1,
			LevelC:    1,
			LevelD:    1,
			Cmp:       []interface{}{},
			MaxLv:     100,
		},
	}
	bytes, err := json.Marshal(statuses)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func (l *loginResponseGenerator) getUserLibrayList(avatarIDs []string,
	inCostumes []MD_Costume,
	inWeapons []MD_Weapon,
	inCostumeStones []MD_CostumeStone,
	inWeaponStones []MD_WeaponStone) []TD_UserLibrary {
	libraries := make([]TD_UserLibrary, 0)
	for _, avatar := range avatarIDs {
		libraries = append(libraries, TD_UserLibrary{
			ItemType: 4,
			MasterId: avatar,
		})
	}
	for _, costume := range inCostumes {
		if costume.ItemKindID != 0 {
			continue
		}
		libraries = append(libraries, TD_UserLibrary{
			ItemType: 1,
			MasterId: costume.ID,
		})
	}
	for _, weapon := range inWeapons {
		if weapon.ItemType != 1 {
			continue
		}
		libraries = append(libraries, TD_UserLibrary{
			ItemType: 2,
			MasterId: weapon.ID,
		})
	}
	for _, costumeStone := range inCostumeStones {
		libraries = append(libraries, TD_UserLibrary{
			ItemType: 26,
			MasterId: costumeStone.ID,
		})
	}
	for _, weaponStone := range inWeaponStones {
		libraries = append(libraries, TD_UserLibrary{
			ItemType: 27,
			MasterId: weaponStone.ID,
		})
	}

	return libraries
}

func (l *loginResponseGenerator) getEquipmentAvatarListAsString() string {
	bytes, err := json.Marshal(l.cfg.EquipmentAvatarsList)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func (l *loginResponseGenerator) loadSubScenarios() []MD_SubScenario {
	subScenarios := make([]MD_SubScenario, 0)
	l.loadMDJson("MD_SubScenario.json", &subScenarios)

	return subScenarios
}

func (l *loginResponseGenerator) getUserCharacterContents(subScenarios []MD_SubScenario, inMemoryLocks []MD_CharacterMemoryLock) []TD_UserCharacterContents {
	memoryLockIDToKeyIDMap := l.createMemoryLockIDToKeyIDMap(inMemoryLocks)
	characterIDToSubScenariosMap := l.createCharacterIDToSubScenariosMap(subScenarios)
	characterIDs := make([]string, len(characterIDToSubScenariosMap))
	i := 0
	for key := range characterIDToSubScenariosMap {
		characterIDs[i] = key
		i++
	}
	sort.Strings(characterIDs)

	characters := make([]TD_UserCharacterContents, 0)
	for _, characterID := range characterIDs {
		subScenarios := characterIDToSubScenariosMap[characterID]
		recaptureSubScenario := map[string]int{}
		memories := map[string]TD_MemoryData{}
		feelSubScenario := map[string]int{}
		costumesSubScenario := map[string]int{}

		for _, subScenario := range subScenarios {
			switch subScenario.Type {
			case 1:
				recaptureSubScenario[subScenario.ID] = 1
				memoryLockID := "MEM_LOCK_" + strings.TrimPrefix(subScenario.MemoryID, "MEM_")
				keyID := memoryLockIDToKeyIDMap[memoryLockID]
				if keyID == "" {
					panic(errors.New("Could not find key ID for memory: " + memoryLockID))
				}
				memories[subScenario.MemoryID] = TD_MemoryData{
					Status:         2,
					UnLockMemories: []string{keyID},
				}
			case 2:
				feelSubScenario[subScenario.ID] = 1
			case 3:
				costumesSubScenario[subScenario.ID] = 1
			}
		}

		c := TD_UserCharacterContents{
			FeelSubScenario:      feelSubScenario,
			Memories:             memories,
			CostumeSubScenario:   costumesSubScenario,
			CharacterId:          characterID,
			RecaptureSubScenario: recaptureSubScenario,
		}
		characters = append(characters, c)
	}

	return characters
}

func (l *loginResponseGenerator) createCharacterIDToSubScenariosMap(subScenarios []MD_SubScenario) map[string][]MD_SubScenario {
	characterIDToSubScenariosMap := map[string][]MD_SubScenario{}

	for _, subScenario := range subScenarios {
		characterSubScenarios, ok := characterIDToSubScenariosMap[subScenario.CharacterID]
		if !ok {
			characterSubScenarios = make([]MD_SubScenario, 0)
		}
		characterSubScenarios = append(characterSubScenarios, subScenario)
		characterIDToSubScenariosMap[subScenario.CharacterID] = characterSubScenarios
	}

	return characterIDToSubScenariosMap
}

func (l *loginResponseGenerator) createMemoryLockIDToKeyIDMap(inMemoryLocks []MD_CharacterMemoryLock) map[string]string {
	m := map[string]string{}
	for _, memoryLock := range inMemoryLocks {
		m[memoryLock.ID] = memoryLock.KeyID
	}

	return m
}

func (l *loginResponseGenerator) getUserQuest() map[string]map[string]TD_UserQuestInfo {
	quests := map[string]map[string]TD_UserQuestInfo{}
	for _, questGroupID := range questGroupIDs {
		questInfoMap := map[string]TD_UserQuestInfo{}
		inQuests := make([]MD_Quest, 0)
		l.loadMDJson("Quests/"+questGroupID+".json", &inQuests)

		for _, quest := range inQuests {
			questInfoMap[quest.ID] = TD_UserQuestInfo{
				ClearMissions:       []string{},
				TotalClearCount:     1,
				TotalChallengeCount: 1,
				Lcd:                 " ",
			}
		}
		quests[questGroupID] = questInfoMap
	}

	return quests
}

func (l *loginResponseGenerator) getUserTowerStatus() map[string]TD_UserTowerStatus {
	towerQuests := make([]MD_TowerQuest, 0)
	l.loadMDJson("MD_TowerQuest.json", &towerQuests)
	lastTowerQuest := towerQuests[len(towerQuests)-1]

	userTowerStatus := map[string]TD_UserTowerStatus{}
	questDataMap := map[string]map[string]TD_UserTowerStatus_QuestData{}
	for _, towerQuest := range towerQuests {
		questIDToQuestDataMap, ok := questDataMap[towerQuest.StageID]
		if !ok {
			questIDToQuestDataMap = map[string]TD_UserTowerStatus_QuestData{}
			questDataMap[towerQuest.StageID] = questIDToQuestDataMap
		}
		questIDToQuestDataMap["TowerQuest::"+towerQuest.ID] = TD_UserTowerStatus_QuestData{
			StarCount:   1,
			MinimumTurn: 1,
			MaxDamage:   1,
		}
	}

	stageIDs := make([]string, len(questDataMap))
	i := 0
	for stageID := range questDataMap {
		stageIDs[i] = stageID
		i++
	}
	sort.Strings(stageIDs)

	statusEntry := TD_UserTowerStatus{
		ResetData: l.cfg.ServerTime.Time.Add(time.Duration(2) * time.Hour).Format(serverTimeFormat),
		Data: TD_UserTowerStatus_MyData{
			SeasonID:         "S202109xx",
			QuestStageID:     lastTowerQuest.StageID,
			LastQuestStageID: lastTowerQuest.StageID,
			QuestData:        questDataMap,
			SR:               stageIDs,
		},
	}
	userTowerStatus["TWTEST001"] = statusEntry

	return userTowerStatus
}

func (l *loginResponseGenerator) getUserAchievement() map[string]TD_UserAchievement {
	inAchievements := make([]MD_Achievement, 0)
	l.loadMDJson("MD_Achievement.json", &inAchievements)

	achievements := map[string]TD_UserAchievement{}
	for _, inAchievement := range inAchievements {
		achievementID := "Achievement::" + inAchievement.ID
		achievements[achievementID] = TD_UserAchievement{
			AchievementDate:     l.cfg.ServerTime.Format(serverTimeFormat),
			Category:            1,
			ResetDayOfWeek:      1,
			ConditionId:         inAchievement.ConditionID,
			GainReward:          1,
			AchievementID:       achievementID,
			MasterName:          "Achievement",
			MasterAchievementId: inAchievement.ID,
		}
	}

	return achievements
}

func (l *loginResponseGenerator) loadCharacterMemoryLocks() []MD_CharacterMemoryLock {
	memoryLocks := make([]MD_CharacterMemoryLock, 0)
	l.loadMDJson("MD_CharacterMemoryLock.json", &memoryLocks)
	return memoryLocks
}

func (l *loginResponseGenerator) loadCostumeStones() []MD_CostumeStone {
	costumeStones := make([]MD_CostumeStone, 0)
	l.loadMDJson("MD_CostumeStone.json", &costumeStones)
	return costumeStones
}

func (l *loginResponseGenerator) loadWeaponStones() []MD_WeaponStone {
	weaponStones := make([]MD_WeaponStone, 0)
	l.loadMDJson("MD_WeaponStone.json", &weaponStones)
	return weaponStones
}

func (l *loginResponseGenerator) getRecaptureItemsAsString(inMemoryLocks []MD_CharacterMemoryLock) string {
	items := map[string]int{}
	for _, memoryLock := range inMemoryLocks {
		items[memoryLock.KeyID] = 2
	}

	itemsBytes, err := json.Marshal(items)
	if err != nil {
		panic(err)
	}

	return string(itemsBytes)
}

func (l *loginResponseGenerator) getUserShootingModeItems() map[string][]string {
	return map[string][]string{
		"30": l.getUserShootingModeItemForType("MD_ShootingModeItemMotion.json"),
		"31": l.getUserShootingModeItemForType("MD_ShootingModeItemFacialExpression.json"),
		"32": l.getUserShootingModeItemForType("MD_ShootingModeItemEffect.json"),
		"33": l.getUserShootingModeItemForType("MD_ShootingModeItemPosing.json"),
		"34": l.getUserShootingModeItemForType("MD_ShootingModeItemBackGround.json"),
	}
}

func (l *loginResponseGenerator) getUserShootingModeItemForType(filename string) []string {
	items := make([]MD_ShootingModeItem, 0)
	l.loadMDJson(filename, &items)

	itemIDs := make([]string, len(items))
	for i, item := range items {
		if item.ID == "" {
			panic(errors.New(filename))
		}
		itemIDs[i] = item.ID
	}

	return itemIDs
}
