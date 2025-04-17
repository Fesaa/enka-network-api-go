package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type PasterConfig struct {
	DefaultUnlock      bool                   `json:"DefaultUnlock"`
	ID                 json.Number            `json:"ID"`
	IncreaseCompletion json.Number            `json:"IncreaseCompletion"`
	PasterTextmap      map[string]json.Number `json:"PasterTextmap"`
	PasterUnlockDesc   map[string]json.Number `json:"PasterUnlockDesc"`
	TextPasterPrefab   string                 `json:"TextPasterPrefab"`
	TravelBrochureID   []json.Number          `json:"TravelBrochureID"`
	Type               string                 `json:"Type"`
}
type PasterConfigAccessor struct {
	_data []PasterConfig
}

// LoadData retrieves the data. Must be called before PasterConfig.GroupData
func (a *PasterConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PasterConfig.json")
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
func (a *PasterConfigAccessor) Raw() ([]PasterConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PasterConfig{}, err
		}
	}
	return a._data, nil
}
