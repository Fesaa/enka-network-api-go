package cache

import _ "embed"

//go:embed resources/genshin_namecards.json
var genshinNameCardJson []byte

//go:embed resources/genshin_profiles.json
var genshinProfileIdentifiersJson []byte

//go:embed resources/genshin_characters.json
var genshinCharactersJson []byte

//go:embed resources/honkai_characters.json
var starRailCharacterJson []byte

//go:embed resources/honkai_avatars.json
var starRailAvataJson []byte

//go:embed resources/honkai_relics.json
var starRailRelicJson []byte

//go:embed resources/honkai_lightcones.json
var starRailLightconesJson []byte

func GetGenshinNameCardJson() []byte {
	return genshinNameCardJson
}

func GetGenshinProfileIdentifiersJson() []byte {
	return genshinProfileIdentifiersJson
}

func GetGenshinCharactersJson() []byte {
	return genshinCharactersJson
}

func GetStarRailCharacterJson() []byte {
	return starRailCharacterJson
}

func GetStarRailAvatarJson() []byte {
	return starRailAvataJson
}

func GetStarRailRelicJson() []byte {
	return starRailRelicJson
}

func GetStarRailLightconesJson() []byte {
	return starRailLightconesJson
}
