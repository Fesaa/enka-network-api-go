package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type RogueUnlockConfig struct {
	RogueUnlockDetail map[string]json.Number `json:"RogueUnlockDetail"`
	RogueUnlockID     json.Number            `json:"RogueUnlockID"`
	UnlockFinishWay   json.Number            `json:"UnlockFinishWay"`
}
type RogueUnlockConfigAccessor struct {
	_data []RogueUnlockConfig
}

// LoadData retrieves the data. Must be called before RogueUnlockConfig.GroupData
func (a *RogueUnlockConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueUnlockConfig.json")
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
func (a *RogueUnlockConfigAccessor) Raw() ([]RogueUnlockConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueUnlockConfig{}, err
		}
	}
	return a._data, nil
}
