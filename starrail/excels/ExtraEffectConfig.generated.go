package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
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
	_data []ExtraEffectConfig
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
