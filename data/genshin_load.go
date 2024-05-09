package data

import (
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
