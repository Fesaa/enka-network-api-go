package data

import (
	_ "embed"
	"encoding/json"

	"github.com/Fesaa/enka-network-api-go/starrail"
	"github.com/Fesaa/enka-network-api-go/utils"
)

func newStarRail() (StarRailData, error) {
	srData := &starRailData{}
	var starRailCharacterData map[string]*starrail.CharacterData
	err := json.Unmarshal(starRailCharacterJson, &starRailCharacterData)
	if err != nil {
		return nil, err
	}

	type Icon struct {
		Icon string `json:"icon"`
	}
	var starRailAvatars map[string]*Icon
	err = json.Unmarshal(starRailAvataJson, &starRailAvatars)
	if err != nil {
		return nil, err
	}

	var avatars map[string]string = make(map[string]string, len(starRailAvatars))
	for k, v := range starRailAvatars {
		avatars[k] = v.Icon
	}

	var relics map[string]*starrail.RelicData
	err = json.Unmarshal(starRailRelicJson, &relics)
	if err != nil {
		return nil, err
	}

	var lightCones map[string]*starrail.LightConeData
	err = json.Unmarshal(starRailLightconesJson, &lightCones)
	if err != nil {
		return nil, err
	}

	srData.StarRailCharacterData = utils.FromMap(starRailCharacterData)
	srData.StarRailAvatars = utils.FromMap(avatars)
	srData.StarRailRelics = utils.FromMap(relics)
	srData.StarRailLightCones = utils.FromMap(lightCones)
	return srData, nil
}
