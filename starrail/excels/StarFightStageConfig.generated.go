package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type StarFightStageConfig struct {
	BattleAreaID    float64   `json:"BattleAreaID"`
	DifficultyLevel string    `json:"DifficultyLevel"`
	EventID         float64   `json:"EventID"`
	GroupID         float64   `json:"GroupID"`
	QuestList       []float64 `json:"QuestList"`
	UnlockQuest     float64   `json:"UnlockQuest"`
}
type StarFightStageConfigAccessor struct {
	_data        []StarFightStageConfig
	_dataEventID map[float64]StarFightStageConfig
}

// LoadData retrieves the data. Must be called before StarFightStageConfig.GroupData
func (a *StarFightStageConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/StarFightStageConfig.json")
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
func (a *StarFightStageConfigAccessor) Raw() ([]StarFightStageConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []StarFightStageConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with StarFightStageConfigAccessor.LoadData to preload everything
func (a *StarFightStageConfigAccessor) GroupData() {
	a._dataEventID = map[float64]StarFightStageConfig{}
	for _, d := range a._data {
		a._dataEventID[d.EventID] = d
	}
}

// ByEventID returns the StarFightStageConfig uniquely identified by EventID
//
// Error is only non-nil if the source errors out
func (a *StarFightStageConfigAccessor) ByEventID(identifier float64) (StarFightStageConfig, error) {
	if a._dataEventID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return StarFightStageConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataEventID[identifier], nil
}
