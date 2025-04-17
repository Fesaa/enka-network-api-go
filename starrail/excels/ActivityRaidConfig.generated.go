package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ActivityRaidConfig struct {
	ActivityModuleID float64 `json:"ActivityModuleID"`
	HardLevel        float64 `json:"HardLevel"`
	RaidID           float64 `json:"RaidID"`
}
type ActivityRaidConfigAccessor struct {
	_data []ActivityRaidConfig
}

// LoadData retrieves the data. Must be called before ActivityRaidConfig.GroupData
func (a *ActivityRaidConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActivityRaidConfig.json")
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
func (a *ActivityRaidConfigAccessor) Raw() ([]ActivityRaidConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActivityRaidConfig{}, err
		}
	}
	return a._data, nil
}
