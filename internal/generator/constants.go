package generator

var costumeRarities map[int]Rarity = map[int]Rarity{
	1: {
		MaxLevel: 1,
		MaxExp:   0,
	},
	2: {
		MaxLevel: 20,
		MaxExp:   13220,
	},
	3: {
		MaxLevel: 30,
		MaxExp:   40080,
	},
	4: {
		MaxLevel: 45,
		MaxExp:   139910,
	},
	5: {
		MaxLevel: 50,
		MaxExp:   198040,
	},
	6: {
		MaxLevel: 60,
		MaxExp:   368560,
	},
}

var weaponRarities map[int]Rarity = map[int]Rarity{
	1: {
		MaxLevel: 20,
		MaxExp:   4044,
	},
	2: {
		MaxLevel: 20,
		MaxExp:   40440,
	},
	3: {
		MaxLevel: 20,
		MaxExp:   404400,
	},
}

const serverTimeFormat = "2006-01-02 15:04:05"

var questGroupIDs = []string{
	"GLBQuestc001",
	"GLBQuestc002",
	"GLBQuestc003",
	"GLBQuestc004",
	"GLBQuestc005",
	"GLBQuestc006",
	"GLBQuestc007",
	"GLBQuestc008",
	"GLBQuestc009",
	"ScenarioQuest",
	"ScenarioQuest2",
	"ScenarioQuest3",
	"ScenarioQuest4",
	"ScenarioQuest5",
	"ScenarioQuest6",
}
