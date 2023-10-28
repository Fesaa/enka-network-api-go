package cache

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/Fesaa/enka-network-api-go/genshin"
)

const MATERIAL_URL = "https://gitlab.com/Dimbreath/AnimeGameData/-/raw/master/ExcelBinOutput/MaterialExcelConfigData.json"

func (m *MemoryCache) loadGenshinResources() error {
	cards, err := loadCards()
	if err != nil {
		return err
	}
	m.GenshinNameCards = cards

	profileIcons, max, err := loadProfileIdentifiers()
	if err != nil {
		return err
	}
	m.GensshinProfileIcons = profileIcons
	m.MaxGenshinProfileId = *max

	characters, err := loadCharacters()
	if err != nil {
		return err
	}
	m.GenshinCharacterData = characters

	materials, err := loadMaterials()
	if err != nil {
		return err
	}
	m.GenshinMaterials = materials

	return nil
}

func loadCards() (map[int]string, error) {
	file, err := os.ReadFile("resources/genshin_namecards.json")
	if err != nil {
		return nil, err
	}

	type Icon struct {
		IconKey string `json:"icon"`
	}
	var genshinNameCards map[string]Icon
	err = json.Unmarshal(file, &genshinNameCards)
	if err != nil {
		return nil, err
	}

	var cards map[int]string = make(map[int]string, len(genshinNameCards))
	for id, card := range genshinNameCards {
		intId, _ := strconv.Atoi(id)
		cards[intId] = card.IconKey
	}
	return cards, nil
}

func loadProfileIdentifiers() (map[int]string, *int, error) {
	file, err := os.ReadFile("resources/genshin_profiles.json")
	if err != nil {
		return nil, nil, err
	}

	type Profile struct {
		Id       int    `json:"id"`
		IconPath string `json:"iconPath"`
	}
	var genshinProfileIcons []Profile
	err = json.Unmarshal(file, &genshinProfileIcons)
	if err != nil {
		return nil, nil, err
	}

	var icons map[int]string = make(map[int]string, len(genshinProfileIcons))
	var max int = 0
	for _, profile := range genshinProfileIcons {
		icons[profile.Id] = profile.IconPath
		if profile.Id > max {
			max = profile.Id
		}
	}
	return icons, &max, nil
}

func loadCharacters() (map[string]*genshin.CharacterData, error) {
	file, err := os.ReadFile("resources/genshin_characters.json")
	if err != nil {
		return nil, err
	}

	var genshinCharacters map[string]*genshin.CharacterData
	err = json.Unmarshal(file, &genshinCharacters)
	if err != nil {
		return nil, err
	}

	for _, character := range genshinCharacters {
		character.WeaponType = genshin.WrapWeaponType(character.RawWeaponType)
	}

	return genshinCharacters, nil
}

func loadMaterials() (map[int]*genshin.Material, error) {

	req, err := http.Get(MATERIAL_URL)
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	if req.StatusCode != 200 {
		return nil, errors.New("enka-network-api-go: Non 200 status code returned: " + req.Status)
	}

	data, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var materials []genshin.RawMaterial
	err = json.Unmarshal(data, &materials)
	if err != nil {
		return nil, err
	}

	var genshinMaterials map[int]*genshin.Material = make(map[int]*genshin.Material, len(materials))
	for _, material := range materials {
		converted := material.ToMaterial()
		genshinMaterials[material.Id] = &converted
	}

	return genshinMaterials, nil
}

func (m *MemoryCache) GetNameCardName(id int) *string {
	if name, ok := m.GenshinNameCards[id]; ok {
		return &name
	}
	return nil
}

func (m *MemoryCache) HasNameCard(id int) bool {
	_, ok := m.GenshinNameCards[id]
	return ok
}

func (m *MemoryCache) AddGenshinUser(user *genshin.RawGenshinUser) {
	m.GenshinUsers[user.Uid] = NewCachedData[*genshin.RawGenshinUser](user)
}

func (m *MemoryCache) GetGenshinUser(uid string) *genshin.RawGenshinUser {
	if cache, ok := m.GenshinUsers[uid]; ok {
		if !cache.IsExpired() {
			return cache.GetData()
		}
		delete(m.GenshinUsers, uid)
		return nil
	}
	return nil
}

func (m *MemoryCache) GetProfileIcon(id int) string {
	if id > m.MaxGenshinProfileId {
		return "Not Implemented"
	}

	if icon, ok := m.GensshinProfileIcons[id]; ok {
		return icon
	}
	return "UI_AvatarIcon_PlayerGirl_Circle"
}

func (m *MemoryCache) GetGenshinCharacterData(name string) *genshin.CharacterData {
	if character, ok := m.GenshinCharacterData[name]; ok {
		return character
	}
	return nil
}

func (m *MemoryCache) GetAllGenshinCharacterData() []*genshin.CharacterData {
	var characters []*genshin.CharacterData = make([]*genshin.CharacterData, 0, len(m.GenshinCharacterData))
	for _, character := range m.GenshinCharacterData {
		characters = append(characters, character)
	}
	return characters
}

func (m *MemoryCache) GetGenshinMaterial(id int) *genshin.Material {
	if material, ok := m.GenshinMaterials[id]; ok {
		return material
	}
	return nil
}
