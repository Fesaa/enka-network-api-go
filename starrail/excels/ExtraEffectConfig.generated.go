package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type ExtraEffectConfig struct {
	DescParamList       []map[string]json.Number `json:"DescParamList"`
	ExtraEffectDesc     map[string]json.Number   `json:"ExtraEffectDesc"`
	ExtraEffectID       json.Number              `json:"ExtraEffectID"`
	ExtraEffectIconPath string                   `json:"ExtraEffectIconPath"`
	ExtraEffectName     map[string]json.Number   `json:"ExtraEffectName"`
	ExtraEffectType     json.Number              `json:"ExtraEffectType"`
}
type ExtraEffectConfigAccessor struct {
	_data              []ExtraEffectConfig
	_dataExtraEffectID map[json.Number]ExtraEffectConfig
}

// LoadData retrieves the data. Must be called before ExtraEffectConfig.GroupData
func (a *ExtraEffectConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ExtraEffectConfig.json")
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
func (a *ExtraEffectConfigAccessor) Raw() ([]ExtraEffectConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ExtraEffectConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ExtraEffectConfigAccessor.LoadData to preload everything
func (a *ExtraEffectConfigAccessor) GroupData() {
	a._dataExtraEffectID = map[json.Number]ExtraEffectConfig{}
	for _, d := range a._data {
		a._dataExtraEffectID[d.ExtraEffectID] = d
	}
}

// ByExtraEffectID returns the ExtraEffectConfig uniquely identified by ExtraEffectID
//
// Error is only non-nil if the source errors out
func (a *ExtraEffectConfigAccessor) ByExtraEffectID(identifier json.Number) (ExtraEffectConfig, error) {
	if a._dataExtraEffectID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ExtraEffectConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataExtraEffectID[identifier], nil
}
