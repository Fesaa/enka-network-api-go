package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type OverrideFloorConfig struct {
	ContentID       json.Number `json:"ContentID"`
	DimensionID     json.Number `json:"DimensionID"`
	EnableCondition string      `json:"EnableCondition"`
	FloorID         json.Number `json:"FloorID"`
}
type OverrideFloorConfigAccessor struct {
	_data []OverrideFloorConfig
}

// LoadData retrieves the data. Must be called before OverrideFloorConfig.GroupData
func (a *OverrideFloorConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/OverrideFloorConfig.json")
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
func (a *OverrideFloorConfigAccessor) Raw() ([]OverrideFloorConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []OverrideFloorConfig{}, err
		}
	}
	return a._data, nil
}
