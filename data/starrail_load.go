package data

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Fesaa/enka-network-api-go/starrail"
	"github.com/Fesaa/enka-network-api-go/utils"
)

var (
	SKILL_TREE_URL = "https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AvatarSkillTreeConfig.json"
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

	var avatars = make(map[string]string, len(starRailAvatars))
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
	srData.StarRailSkillTree, err = parseHSRSkillTree()
	if err != nil {
		return nil, err
	}

	return srData, nil
}

func parseHSRSkillTree() (*utils.Map[string, map[starrail.SkillTreeAnchor]starrail.SkillTreeNode], error) {
	res, err := http.Get(SKILL_TREE_URL)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New(res.Status)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var skillTree []starrail.SkillTreeNode
	if err = json.Unmarshal(body, &skillTree); err != nil {
		return nil, err
	}

	m := map[string]map[starrail.SkillTreeAnchor]starrail.SkillTreeNode{}
	for _, skillTreeNode := range skillTree {
		characterID := fmt.Sprintf("%d", skillTreeNode.AvatarID)
		anchorMap := m[characterID]
		if anchorMap == nil {
			anchorMap = make(map[starrail.SkillTreeAnchor]starrail.SkillTreeNode)
		}

		anchorMap[skillTreeNode.Anchor] = skillTreeNode
		m[characterID] = anchorMap
	}

	return utils.FromMap(m), nil
}
