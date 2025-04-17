package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type RecolorConfig struct {
	DefaultColor string  `json:"DefaultColor"`
	ID           float64 `json:"ID"`
	WhiteBGColor string  `json:"WhiteBGColor"`
}
type RecolorConfigAccessor struct {
	_data   []RecolorConfig
	_dataID map[float64]RecolorConfig
}

// LoadData retrieves the data. Must be called before RecolorConfig.GroupData
func (a *RecolorConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RecolorConfig.json")
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
func (a *RecolorConfigAccessor) Raw() ([]RecolorConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RecolorConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RecolorConfigAccessor.LoadData to preload everything
func (a *RecolorConfigAccessor) GroupData() {
	a._dataID = map[float64]RecolorConfig{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the RecolorConfig uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *RecolorConfigAccessor) ByID(identifier float64) (RecolorConfig, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RecolorConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
