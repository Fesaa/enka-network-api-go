package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ActivityRogueAreaOverride struct {
	RecommendLevel           float64            `json:"RecommendLevel"`
	RecommendSkillTreePoints float64            `json:"RecommendSkillTreePoints"`
	RogueAreaID              float64            `json:"RogueAreaID"`
	ScoreMap                 map[string]float64 `json:"ScoreMap"`
	WorldLevel               float64            `json:"WorldLevel"`
}
type ActivityRogueAreaOverrideAccessor struct {
	_data []ActivityRogueAreaOverride
}

// LoadData retrieves the data. Must be called before ActivityRogueAreaOverride.GroupData
func (a *ActivityRogueAreaOverrideAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActivityRogueAreaOverride.json")
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
func (a *ActivityRogueAreaOverrideAccessor) Raw() ([]ActivityRogueAreaOverride, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActivityRogueAreaOverride{}, err
		}
	}
	return a._data, nil
}
