package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type AvatarSkillConfig struct {
	AttackType              string          `json:"AttackType"`
	BPAdd                   hash.IntValue   `json:"BPAdd"`
	BPNeed                  hash.IntValue   `json:"BPNeed"`
	CoolDown                float64         `json:"CoolDown"`
	DelayRatio              hash.IntValue   `json:"DelayRatio"`
	ExtraEffectIDList       []float64       `json:"ExtraEffectIDList"`
	HideInUI                bool            `json:"HideInUI"`
	InitCoolDown            float64         `json:"InitCoolDown"`
	Level                   float64         `json:"Level"`
	LevelUpCostList         []interface{}   `json:"LevelUpCostList"`
	MaxLevel                float64         `json:"MaxLevel"`
	ParamList               []hash.IntValue `json:"ParamList"`
	RatedRankID             []float64       `json:"RatedRankID"`
	RatedSkillTreeID        []float64       `json:"RatedSkillTreeID"`
	SPBase                  hash.IntValue   `json:"SPBase"`
	SPMultipleRatio         hash.IntValue   `json:"SPMultipleRatio"`
	SPNeed                  hash.IntValue   `json:"SPNeed"`
	ShowDamageList          []interface{}   `json:"ShowDamageList"`
	ShowHealList            []interface{}   `json:"ShowHealList"`
	ShowStanceList          []hash.IntValue `json:"ShowStanceList"`
	SimpleExtraEffectIDList []float64       `json:"SimpleExtraEffectIDList"`
	SimpleParamList         []hash.IntValue `json:"SimpleParamList"`
	SimpleSkillDesc         hash.Hash       `json:"SimpleSkillDesc"`
	SkillComboValueDelta    hash.IntValue   `json:"SkillComboValueDelta"`
	SkillDesc               hash.Hash       `json:"SkillDesc"`
	SkillEffect             string          `json:"SkillEffect"`
	SkillID                 float64         `json:"SkillID"`
	SkillIcon               string          `json:"SkillIcon"`
	SkillName               hash.Hash       `json:"SkillName"`
	SkillNeed               hash.Hash       `json:"SkillNeed"`
	SkillTag                hash.Hash       `json:"SkillTag"`
	SkillTriggerKey         string          `json:"SkillTriggerKey"`
	SkillTypeDesc           hash.Hash       `json:"SkillTypeDesc"`
	StanceDamageDisplay     float64         `json:"StanceDamageDisplay"`
	StanceDamageType        string          `json:"StanceDamageType"`
	UltraSkillIcon          string          `json:"UltraSkillIcon"`
}
type AvatarSkillConfigAccessor struct {
	_data []AvatarSkillConfig
}

// LoadData retrieves the data. Must be called before AvatarSkillConfig.GroupData
func (a *AvatarSkillConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AvatarSkillConfig.json")
	if err != nil {
		return err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &a._data)
}

// Raw returns the raw data.
func (a *AvatarSkillConfigAccessor) Raw() ([]AvatarSkillConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AvatarSkillConfig{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}
