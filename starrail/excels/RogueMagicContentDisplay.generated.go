package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type RogueMagicContentDisplay struct {
	DisplayContent map[string]json.Number `json:"DisplayContent"`
	DisplayID      json.Number            `json:"DisplayID"`
}
type RogueMagicContentDisplayAccessor struct {
	_data []RogueMagicContentDisplay
}

// LoadData retrieves the data. Must be called before RogueMagicContentDisplay.GroupData
func (a *RogueMagicContentDisplayAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueMagicContentDisplay.json")
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
func (a *RogueMagicContentDisplayAccessor) Raw() ([]RogueMagicContentDisplay, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueMagicContentDisplay{}, err
		}
	}
	return a._data, nil
}
