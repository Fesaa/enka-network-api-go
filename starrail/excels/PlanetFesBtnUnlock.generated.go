package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type PlanetFesBtnUnlock struct {
	ID            string  `json:"ID"`
	UnlockQuestID float64 `json:"UnlockQuestID"`
}
type PlanetFesBtnUnlockAccessor struct {
	_data   []PlanetFesBtnUnlock
	_dataID map[string]PlanetFesBtnUnlock
}

// LoadData retrieves the data. Must be called before PlanetFesBtnUnlock.GroupData
func (a *PlanetFesBtnUnlockAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PlanetFesBtnUnlock.json")
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
func (a *PlanetFesBtnUnlockAccessor) Raw() ([]PlanetFesBtnUnlock, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PlanetFesBtnUnlock{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PlanetFesBtnUnlockAccessor.LoadData to preload everything
func (a *PlanetFesBtnUnlockAccessor) GroupData() {
	a._dataID = map[string]PlanetFesBtnUnlock{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the PlanetFesBtnUnlock uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *PlanetFesBtnUnlockAccessor) ByID(identifier string) (PlanetFesBtnUnlock, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlanetFesBtnUnlock{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
