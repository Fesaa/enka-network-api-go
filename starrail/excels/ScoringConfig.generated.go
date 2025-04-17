package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type ScoringConfig struct {
	AbilityName     string        `json:"AbilityName"`
	DisplayTypeList []string      `json:"DisplayTypeList"`
	GameModeGroup   json.Number   `json:"GameModeGroup"`
	ParamList       []interface{} `json:"ParamList"`
	ScoringID       json.Number   `json:"ScoringID"`
}
type ScoringConfigAccessor struct {
	_data          []ScoringConfig
	_dataScoringID map[json.Number]ScoringConfig
}

// LoadData retrieves the data. Must be called before ScoringConfig.GroupData
func (a *ScoringConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ScoringConfig.json")
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
func (a *ScoringConfigAccessor) Raw() ([]ScoringConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ScoringConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ScoringConfigAccessor.LoadData to preload everything
func (a *ScoringConfigAccessor) GroupData() {
	a._dataScoringID = map[json.Number]ScoringConfig{}
	for _, d := range a._data {
		a._dataScoringID[d.ScoringID] = d
	}
}

// ByScoringID returns the ScoringConfig uniquely identified by ScoringID
//
// Error is only non-nil if the source errors out
func (a *ScoringConfigAccessor) ByScoringID(identifier json.Number) (ScoringConfig, error) {
	if a._dataScoringID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ScoringConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataScoringID[identifier], nil
}
