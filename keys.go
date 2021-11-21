package main

var apiKey0 = [...]string{
	"mfQMG92mJ6R",
	"D2L1L1aLJSBjLlLQ3fCdS8at",
	"52hHwSInpbOVRq",
	"M88yny3dvB5ZTVHpQD",
}
var apiKey1 = [...]string{
	"MW1LkPfxIrDB7PNTdupvw",
	"jAWLuydO",
	"sIQPwGnXVkOhMqks0p",
	"W7rkDhakfa1PI0",
}

func GetApiKey(idx int) string {
	return apiKey0[idx] + apiKey1[idx]
}

func GetApiKeyNext(idx int) string {
	idx = (idx + 1) % 4
	return apiKey0[idx] + apiKey1[idx]
}

var saveDataKey = []byte("HgJc2tsBsZh3QJRcancefgTnmy6zmgiG")
var saveDataIV = []byte("ancefgTnmy6zmgiG")
