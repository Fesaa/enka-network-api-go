package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type PreAvatarLevelingTemplate struct {
	BossMaterialAmount        float64 `json:"BossMaterialAmount"`
	CoinAmount                float64 `json:"CoinAmount"`
	ExpAmount                 float64 `json:"ExpAmount"`
	PromotionMaterialAmount   float64 `json:"PromotionMaterialAmount"`
	SkillMaterialLargeAmount  float64 `json:"SkillMaterialLargeAmount"`
	SkillMaterialMediumAmount float64 `json:"SkillMaterialMediumAmount"`
	SkillMaterialSmallAmount  float64 `json:"SkillMaterialSmallAmount"`
	TemplateID                float64 `json:"TemplateID"`
	TracksDestinyAmount       float64 `json:"TracksDestinyAmount"`
	WorldLevel                float64 `json:"WorldLevel"`
	WorldMaterialLargeAmount  float64 `json:"WorldMaterialLargeAmount"`
	WorldMaterialMediumAmount float64 `json:"WorldMaterialMediumAmount"`
	WorldMaterialSmallAmount  float64 `json:"WorldMaterialSmallAmount"`
}
type PreAvatarLevelingTemplateAccessor struct {
	_data []PreAvatarLevelingTemplate
}

// LoadData retrieves the data. Must be called before PreAvatarLevelingTemplate.GroupData
func (a *PreAvatarLevelingTemplateAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PreAvatarLevelingTemplate.json")
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
func (a *PreAvatarLevelingTemplateAccessor) Raw() ([]PreAvatarLevelingTemplate, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PreAvatarLevelingTemplate{}, err
		}
	}
	return a._data, nil
}
