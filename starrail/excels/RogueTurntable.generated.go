package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type RogueTurntable struct {
	ParamGroupID float64 `json:"ParamGroupID"`
	PrepareTime  float64 `json:"PrepareTime"`
	RewardLevel  string  `json:"RewardLevel"`
}
type RogueTurntableAccessor struct {
	_data []RogueTurntable
}

// LoadData retrieves the data. Must be called before RogueTurntable.GroupData
func (a *RogueTurntableAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueTurntable.json")
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
func (a *RogueTurntableAccessor) Raw() ([]RogueTurntable, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueTurntable{}, err
		}
	}
	return a._data, nil
}
