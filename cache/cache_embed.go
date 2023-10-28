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
