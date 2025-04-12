package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Fesaa/enka-network-api-go/starrail"
	"github.com/Fesaa/enka-network-api-go/utils"
	"io"
	"log/slog"
	"net/http"
)

var (
	ExcelOutputURL = "https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/%s.json"
)

type hsrExcels struct {
	log *slog.Logger

	skillTree      *utils.Map[string, map[starrail.SkillTreeAnchor]starrail.SkillTreeNode]
	relicMainAffix []starrail.RelicMainAffixConfig
	relicSubAffix  []starrail.RelicSubAffixConfig
	setConfig      *utils.Map[string, *starrail.RelicSetConfig]
	relicConfig    *utils.Map[string, *starrail.RelicConfig]
}

func NewExcels(logs ...*slog.Logger) HSRExcels {
	log := func() *slog.Logger {
		if len(logs) > 0 {
			return logs[0]
		}
		return slog.Default()
	}()

	return &hsrExcels{
		log: log,
	}
}

func (e *hsrExcels) SkillTree(id string) map[starrail.SkillTreeAnchor]starrail.SkillTreeNode {
	if e.skillTree == nil {
		if err := e.parseHSRSkillTree(); err != nil {
			e.log.Error("Failed to parse HSRSkillTree", "error", err)
			return nil
		}
	}

	if data, ok := e.skillTree.Get(id); ok {
		return data
	}
	return nil
}

func (e *hsrExcels) RelicMainAffix() []starrail.RelicMainAffixConfig {
	if e.relicMainAffix == nil {
		if err := e.parseFromExcels("RelicMainAffixConfig", &e.relicMainAffix); err != nil {
			e.log.Error("Failed to parse RelicMainAffixConfig", "error", err)
			return nil
		}
	}

	return e.relicMainAffix
}

func (e *hsrExcels) RelicSubAffixConfig() []starrail.RelicSubAffixConfig {
	if e.relicSubAffix == nil {
		if err := e.parseFromExcels("RelicSubAffixConfig", &e.relicSubAffix); err != nil {
			e.log.Error("Failed to parse RelicMainAffixConfig", "error", err)
			return nil
		}
	}

	return e.relicSubAffix
}

func (e *hsrExcels) RelicSetConfig(i string) (*starrail.RelicSetConfig, bool) {
	if e.setConfig == nil {
		if err := e.parseRelicSetConfig(); err != nil {
			e.log.Error("Failed to parse RelicSetConfig", "error", err)
			return nil, false
		}
	}

	return e.setConfig.Get(i)
}

func (e *hsrExcels) RelicConfig(i string) (*starrail.RelicConfig, bool) {
	if e.relicConfig == nil {
		if err := e.parseRelicConfig(); err != nil {
			e.log.Error("Failed to parse RelicConfig", "error", err)
			return nil, false
		}
	}

	return e.relicConfig.Get(i)
}

func (e *hsrExcels) parseRelicSetConfig() error {
	var sets []*starrail.RelicSetConfig
	if err := e.parseFromExcels("RelicSetConfig", &sets); err != nil {
		return err
	}

	m := make(map[string]*starrail.RelicSetConfig)
	for _, set := range sets {
		m[fmt.Sprintf("%d", set.SetID)] = set
	}

	e.setConfig = utils.FromMap(m)
	return nil
}

func (e *hsrExcels) parseRelicConfig() error {
	var sets []*starrail.RelicConfig
	if err := e.parseFromExcels("RelicConfig", &sets); err != nil {
		return err
	}

	m := make(map[string]*starrail.RelicConfig)
	for _, set := range sets {
		m[fmt.Sprintf("%d", set.ID)] = set
	}

	e.relicConfig = utils.FromMap(m)
	return nil
}

func (e *hsrExcels) parseHSRSkillTree() error {
	var skillTree []starrail.SkillTreeNode
	if err := e.parseFromExcels("AvatarSkillTreeConfig", &skillTree); err != nil {
		return err
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

	e.skillTree = utils.FromMap(m)
	return nil
}

func (e *hsrExcels) parseFromExcels(file string, out any) error {
	res, err := http.Get(fmt.Sprintf(ExcelOutputURL, file))
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New(res.Status)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, out)
}
