package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type StaminaSaleConfig struct {
	Price     StaminaSaleConfigPrice `json:"Price"`
	Times     float64                `json:"Times"`
	ToStamina float64                `json:"ToStamina"`
}
type StaminaSaleConfigPrice struct {
	F1 float64 `json:"1"`
}
type StaminaSaleConfigAccessor struct {
	_data      []StaminaSaleConfig
	_dataTimes map[float64]StaminaSaleConfig
}

// LoadData retrieves the data. Must be called before StaminaSaleConfig.GroupData
func (a *StaminaSaleConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/StaminaSaleConfig.json")
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
func (a *StaminaSaleConfigAccessor) Raw() ([]StaminaSaleConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []StaminaSaleConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with StaminaSaleConfigAccessor.LoadData to preload everything
func (a *StaminaSaleConfigAccessor) GroupData() {
	a._dataTimes = map[float64]StaminaSaleConfig{}
	for _, d := range a._data {
		a._dataTimes[d.Times] = d
	}
}

// ByTimes returns the StaminaSaleConfig uniquely identified by Times
//
// Error is only non-nil if the source errors out
func (a *StaminaSaleConfigAccessor) ByTimes(identifier float64) (StaminaSaleConfig, error) {
	if a._dataTimes == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return StaminaSaleConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataTimes[identifier], nil
}
