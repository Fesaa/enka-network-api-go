package data

import (
	_ "embed"
	"encoding/json"

	"github.com/Fesaa/enka-network-api-go/starrail"
	"github.com/Fesaa/enka-network-api-go/utils"
)

func (m *memoryCache) loadStarRailResources() error {
	var starRailCharacterData map[string]*starrail.CharacterData
	err := json.Unmarshal(starRailCharacterJson, &starRailCharacterData)
	if err != nil {
		return err
	}

	type Icon struct {
		Icon string `json:"icon"`
	}
	var starRailAvatars map[string]*Icon
	err = json.Unmarshal(starRailAvataJson, &starRailAvatars)
	if err != nil {
		return err
	}

	var avatars map[string]string = make(map[string]string, len(starRailAvatars))
	for k, v := range starRailAvatars {
		avatars[k] = v.Icon
	}

	var relics map[string]*starrail.RelicData
	err = json.Unmarshal(starRailRelicJson, &relics)
	if err != nil {
		return err
	}

	var lightCones map[string]*starrail.LightConeData
	err = json.Unmarshal(starRailLightconesJson, &lightCones)
	if err != nil {
		return err
	}

	m.StarRailCharacterData = utils.FromMap(starRailCharacterData)
	m.StarRailAvatars = utils.FromMap(avatars)
	m.StarRailRelics = utils.FromMap(relics)
	m.StarRailLightCones = utils.FromMap(lightCones)
	return nil
}
