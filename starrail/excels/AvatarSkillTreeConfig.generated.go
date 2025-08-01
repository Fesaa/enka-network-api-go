package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type AvatarSkillTreeConfig struct {
	AbilityName             string                               `json:"AbilityName"`
	AnchorType              string                               `json:"AnchorType"`
	AvatarID                float64                              `json:"AvatarID"`
	AvatarLevelLimit        float64                              `json:"AvatarLevelLimit"`
	AvatarPromotionLimit    float64                              `json:"AvatarPromotionLimit"`
	DefaultUnlock           bool                                 `json:"DefaultUnlock"`
	EnhancedID              float64                              `json:"EnhancedID"`
	ExtraEffectIDList       []float64                            `json:"ExtraEffectIDList"`
	IconPath                string                               `json:"IconPath"`
	Level                   float64                              `json:"Level"`
	LevelUpSkillID          []float64                            `json:"LevelUpSkillID"`
	MaterialList            []AvatarSkillTreeConfigMaterialList  `json:"MaterialList"`
	MaxLevel                float64                              `json:"MaxLevel"`
	ParamList               []hash.Value                         `json:"ParamList"`
	PointDesc               string                               `json:"PointDesc"`
	PointID                 float64                              `json:"PointID"`
	PointName               string                               `json:"PointName"`
	PointTriggerKey         string                               `json:"PointTriggerKey"`
	PointType               float64                              `json:"PointType"`
	PrePoint                []float64                            `json:"PrePoint"`
	RecommendPriority       float64                              `json:"RecommendPriority"`
	SimpleExtraEffectIDList []interface{}                        `json:"SimpleExtraEffectIDList"`
	SimplePointDesc         string                               `json:"SimplePointDesc"`
	StatusAddList           []AvatarSkillTreeConfigStatusAddList `json:"StatusAddList"`
}
type AvatarSkillTreeConfigMaterialList struct {
	ItemID  float64 `json:"ItemID"`
	ItemNum float64 `json:"ItemNum"`
}
type AvatarSkillTreeConfigStatusAddList struct {
	PropertyType string     `json:"PropertyType"`
	Value        hash.Value `json:"Value"`
}
type AvatarSkillTreeConfigAccessor struct {
	_data []AvatarSkillTreeConfig
}

// LoadData retrieves the data. Must be called before AvatarSkillTreeConfig.GroupData
func (a *AvatarSkillTreeConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AvatarSkillTreeConfig.json")
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
func (a *AvatarSkillTreeConfigAccessor) Raw() ([]AvatarSkillTreeConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AvatarSkillTreeConfig{}, err
		}
	}
	return a._data, nil
}
