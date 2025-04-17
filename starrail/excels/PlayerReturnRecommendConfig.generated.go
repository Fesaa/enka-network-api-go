package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type PlayerReturnRecommendConfig struct {
	Condition   []PlayerReturnRecommendConfigCondition `json:"Condition"`
	GachaID     json.Number                            `json:"GachaID"`
	GotoID      json.Number                            `json:"GotoID"`
	ImagePath   string                                 `json:"ImagePath"`
	PanelID     json.Number                            `json:"PanelID"`
	RecommendID json.Number                            `json:"RecommendID"`
	Title       map[string]json.Number                 `json:"Title"`
	Type        string                                 `json:"Type"`
	Weight      json.Number                            `json:"Weight"`
}
type PlayerReturnRecommendConfigCondition struct {
	Param string `json:"Param"`
	Type  string `json:"Type"`
}
type PlayerReturnRecommendConfigAccessor struct {
	_data            []PlayerReturnRecommendConfig
	_dataRecommendID map[json.Number]PlayerReturnRecommendConfig
	_dataWeight      map[json.Number]PlayerReturnRecommendConfig
}

// LoadData retrieves the data. Must be called before PlayerReturnRecommendConfig.GroupData
func (a *PlayerReturnRecommendConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PlayerReturnRecommendConfig.json")
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
func (a *PlayerReturnRecommendConfigAccessor) Raw() ([]PlayerReturnRecommendConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PlayerReturnRecommendConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PlayerReturnRecommendConfigAccessor.LoadData to preload everything
func (a *PlayerReturnRecommendConfigAccessor) GroupData() {
	a._dataRecommendID = map[json.Number]PlayerReturnRecommendConfig{}
	a._dataWeight = map[json.Number]PlayerReturnRecommendConfig{}
	for _, d := range a._data {
		a._dataRecommendID[d.RecommendID] = d
		a._dataWeight[d.Weight] = d
	}
}

// ByRecommendID returns the PlayerReturnRecommendConfig uniquely identified by RecommendID
//
// Error is only non-nil if the source errors out
func (a *PlayerReturnRecommendConfigAccessor) ByRecommendID(identifier json.Number) (PlayerReturnRecommendConfig, error) {
	if a._dataRecommendID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlayerReturnRecommendConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRecommendID[identifier], nil
}

// ByWeight returns the PlayerReturnRecommendConfig uniquely identified by Weight
//
// Error is only non-nil if the source errors out
func (a *PlayerReturnRecommendConfigAccessor) ByWeight(identifier json.Number) (PlayerReturnRecommendConfig, error) {
	if a._dataWeight == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlayerReturnRecommendConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataWeight[identifier], nil
}
