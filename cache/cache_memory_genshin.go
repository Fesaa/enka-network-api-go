package cache

import (
	_ "embed"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/Fesaa/enka-network-api-go/genshin"
	"github.com/Fesaa/enka-network-api-go/utils"
)

const MATERIAL_URL = "https://gitlab.com/Dimbreath/AnimeGameData/-/raw/master/ExcelBinOutput/MaterialExcelConfigData.json"

func (m *memoryCache) loadGenshinResources() error {
	cards, err := loadCards()
	if err != nil {
		return err
	}
	m.GenshinNameCards = cards

	profileIcons, err := loadProfileIdentifiers()
	if err != nil {
		return err
	}
	m.GenshinProfileIcons = profileIcons

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

func loadCards() (*utils.Map[int, string], error) {
	type Icon struct {
		IconKey string `json:"icon"`
	}
	var genshinNameCards map[string]Icon
	err := json.Unmarshal(genshinNameCardJson, &genshinNameCards)
	if err != nil {
		return nil, err
	}

	cards := utils.NewMap[int, string]()
	for id, card := range genshinNameCards {
		intId, _ := strconv.Atoi(id)
		cards.Set(intId, card.IconKey)
	}
	return cards, nil
}

func loadProfileIdentifiers() (*utils.Map[string, *genshin.ProfilePicture], error) {
	var genshinProfileIcons map[string]genshin.ProfilePicture
	err := json.Unmarshal(genshinProfileIdentifiersJson, &genshinProfileIcons)
	if err != nil {
		return nil, err
	}

	icons := utils.NewMap[string, *genshin.ProfilePicture]()
	for id, profile := range genshinProfileIcons {
		icons.Set(id, &genshin.ProfilePicture{
			Id:       id,
			IconPath: profile.IconPath,
		})
	}
	return icons, nil
}

func loadCharacters() (*utils.Map[string, *genshin.CharacterData], error) {
	genshinCharactersMap := make(map[string]*genshin.CharacterData)
	err := json.Unmarshal(genshinCharactersJson, &genshinCharactersMap)
	if err != nil {
		return nil, err
	}

	genshinCharacters := utils.FromMap(genshinCharactersMap)
	genshinCharacters.ForEach(func(key string, character *genshin.CharacterData) {
		character.WeaponType = genshin.WrapWeaponType(character.RawWeaponType)
	})

	return genshinCharacters, nil
}

func loadMaterials() (*utils.Map[int, *genshin.RawMaterial], error) {
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

	// For some reasons we can't store them as pointers
	// Maybe because they're from a Get request?
	// I'm not sure, pretty confused
	// Do tell me if you know why
	materialsMap := utils.NewMap[int, *genshin.RawMaterial]()
	for _, material := range materials {
		materialsMap.Set(material.Id, &material)
	}
	return materialsMap, nil
}

func (m *memoryCache) GetNameCardName(id int) *string {
	if name, ok := m.GenshinNameCards.Get(id); ok {
		return &name
	}
	return nil
}

func (m *memoryCache) HasNameCard(id int) bool {
	_, ok := m.GenshinNameCards.Get(id)
	return ok
}

func (m *memoryCache) AddGenshinUser(user *genshin.RawGenshinUser) {
	m.GenshinUsers.Set(user.Uid, NewCachedData[*genshin.RawGenshinUser](user))
}

func (m *memoryCache) GetGenshinUser(uid string) *genshin.RawGenshinUser {
	if cache, ok := m.GenshinUsers.Get(uid); ok {
		if !cache.IsExpired() {
			return cache.GetData()
		}
		m.GenshinUsers.Delete(uid)
		return nil
	}
	return nil
}

// If this returns an error, I have the update the lib.
// You can mostly assume it'll return succesfully.
func (m *memoryCache) GetProfileIcon(id string) string {
	icon, ok := m.GenshinProfileIcons.Get(id)
	if !ok {
		return ""
	}
	return icon.IconPath
}

func (m *memoryCache) GetGenshinCharacterData(name string) *genshin.CharacterData {
	if character, ok := m.GenshinCharacterData.Get(name); ok {
		return character
	}
	return nil
}

func (m *memoryCache) GetAllGenshinCharacterData() []*genshin.CharacterData {
	var characters []*genshin.CharacterData = make([]*genshin.CharacterData, 0, m.GenshinCharacterData.Len())
	m.GenshinCharacterData.ForEach(func(_ string, character *genshin.CharacterData) {
		characters = append(characters, character)
	})
	return characters
}

func (m *memoryCache) GetGenshinMaterial(id int) *genshin.RawMaterial {
	if material, ok := m.GenshinMaterials.Get(id); ok {
		return material
	}
	return nil
}
