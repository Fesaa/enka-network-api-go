package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type ActivityConfig struct {
	ActivityID           json.Number   `json:"ActivityID"`
	ActivityModuleIDList []json.Number `json:"ActivityModuleIDList"`
	ActivityPanelID      json.Number   `json:"ActivityPanelID"`
	EarlyAccessContentID json.Number   `json:"EarlyAccessContentID"`
	ResidentModuleList   []json.Number `json:"ResidentModuleList"`
}
type ActivityConfigAccessor struct {
	_data []ActivityConfig
}

// LoadData retrieves the data. Must be called before ActivityConfig.GroupData
func (a *ActivityConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActivityConfig.json")
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
func (a *ActivityConfigAccessor) Raw() ([]ActivityConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActivityConfig{}, err
		}
	}
	return a._data, nil
}
