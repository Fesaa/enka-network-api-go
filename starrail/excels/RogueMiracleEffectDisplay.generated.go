package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type RogueMiracleEffectDisplay struct {
	DescParamList          []map[string]json.Number `json:"DescParamList"`
	ExtraEffect            []json.Number            `json:"ExtraEffect"`
	MiracleDesc            map[string]json.Number   `json:"MiracleDesc"`
	MiracleEffectDisplayID json.Number              `json:"MiracleEffectDisplayID"`
	MiracleSimpleDesc      map[string]json.Number   `json:"MiracleSimpleDesc"`
}
type RogueMiracleEffectDisplayAccessor struct {
	_data []RogueMiracleEffectDisplay
}

// LoadData retrieves the data. Must be called before RogueMiracleEffectDisplay.GroupData
func (a *RogueMiracleEffectDisplayAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueMiracleEffectDisplay.json")
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
func (a *RogueMiracleEffectDisplayAccessor) Raw() ([]RogueMiracleEffectDisplay, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueMiracleEffectDisplay{}, err
		}
	}
	return a._data, nil
}
