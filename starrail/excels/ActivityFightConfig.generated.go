package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ActivityFightConfig struct {
	ActivityFightGroupID float64 `json:"ActivityFightGroupID"`
	DifficultyLevel      string  `json:"DifficultyLevel"`
	FightEventID         float64 `json:"FightEventID"`
	OffsetLevel          float64 `json:"OffsetLevel"`
	RewardID             float64 `json:"RewardID"`
	RewardQuest          float64 `json:"RewardQuest"`
	RewardWave           float64 `json:"RewardWave"`
	RewardWave2          float64 `json:"RewardWave2"`
	RoundsLimit          float64 `json:"RoundsLimit"`
	TotalWave            float64 `json:"TotalWave"`
}
type ActivityFightConfigAccessor struct {
	_data             []ActivityFightConfig
	_dataFightEventID map[float64]ActivityFightConfig
}

// LoadData retrieves the data. Must be called before ActivityFightConfig.GroupData
func (a *ActivityFightConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActivityFightConfig.json")
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
func (a *ActivityFightConfigAccessor) Raw() ([]ActivityFightConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActivityFightConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ActivityFightConfigAccessor.LoadData to preload everything
func (a *ActivityFightConfigAccessor) GroupData() {
	a._dataFightEventID = map[float64]ActivityFightConfig{}
	for _, d := range a._data {
		a._dataFightEventID[d.FightEventID] = d
	}
}

// ByFightEventID returns the ActivityFightConfig uniquely identified by FightEventID
//
// Error is only non-nil if the source errors out
func (a *ActivityFightConfigAccessor) ByFightEventID(identifier float64) (ActivityFightConfig, error) {
	if a._dataFightEventID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityFightConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataFightEventID[identifier], nil
}
