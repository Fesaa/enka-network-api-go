package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type AetherDivideSpiritTrial struct {
	ID        float64 `json:"ID"`
	Promotion float64 `json:"Promotion"`
	SpiritID  float64 `json:"SpiritID"`
}
type AetherDivideSpiritTrialAccessor struct {
	_data         []AetherDivideSpiritTrial
	_dataID       map[float64]AetherDivideSpiritTrial
	_dataSpiritID map[float64]AetherDivideSpiritTrial
}

// LoadData retrieves the data. Must be called before AetherDivideSpiritTrial.GroupData
func (a *AetherDivideSpiritTrialAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AetherDivideSpiritTrial.json")
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
func (a *AetherDivideSpiritTrialAccessor) Raw() ([]AetherDivideSpiritTrial, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AetherDivideSpiritTrial{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AetherDivideSpiritTrialAccessor.LoadData to preload everything
func (a *AetherDivideSpiritTrialAccessor) GroupData() {
	a._dataID = map[float64]AetherDivideSpiritTrial{}
	a._dataSpiritID = map[float64]AetherDivideSpiritTrial{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
		a._dataSpiritID[d.SpiritID] = d
	}
}

// ByID returns the AetherDivideSpiritTrial uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *AetherDivideSpiritTrialAccessor) ByID(identifier float64) (AetherDivideSpiritTrial, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDivideSpiritTrial{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// BySpiritID returns the AetherDivideSpiritTrial uniquely identified by SpiritID
//
// Error is only non-nil if the source errors out
func (a *AetherDivideSpiritTrialAccessor) BySpiritID(identifier float64) (AetherDivideSpiritTrial, error) {
	if a._dataSpiritID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDivideSpiritTrial{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSpiritID[identifier], nil
}
