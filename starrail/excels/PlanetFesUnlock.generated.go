package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type PlanetFesUnlock struct {
	FinishWayID float64   `json:"FinishWayID"`
	UnlockDesc  hash.Hash `json:"UnlockDesc"`
	UnlockID    float64   `json:"UnlockID"`
}
type PlanetFesUnlockAccessor struct {
	_data            []PlanetFesUnlock
	_dataFinishWayID map[float64]PlanetFesUnlock
	_dataUnlockID    map[float64]PlanetFesUnlock
}

// LoadData retrieves the data. Must be called before PlanetFesUnlock.GroupData
func (a *PlanetFesUnlockAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PlanetFesUnlock.json")
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
func (a *PlanetFesUnlockAccessor) Raw() ([]PlanetFesUnlock, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PlanetFesUnlock{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PlanetFesUnlockAccessor.LoadData to preload everything
func (a *PlanetFesUnlockAccessor) GroupData() {
	a._dataFinishWayID = map[float64]PlanetFesUnlock{}
	a._dataUnlockID = map[float64]PlanetFesUnlock{}
	for _, d := range a._data {
		a._dataFinishWayID[d.FinishWayID] = d
		a._dataUnlockID[d.UnlockID] = d
	}
}

// ByFinishWayID returns the PlanetFesUnlock uniquely identified by FinishWayID
//
// Error is only non-nil if the source errors out
func (a *PlanetFesUnlockAccessor) ByFinishWayID(identifier float64) (PlanetFesUnlock, error) {
	if a._dataFinishWayID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlanetFesUnlock{}, err
			}
		}
		a.GroupData()
	}
	return a._dataFinishWayID[identifier], nil
}

// ByUnlockID returns the PlanetFesUnlock uniquely identified by UnlockID
//
// Error is only non-nil if the source errors out
func (a *PlanetFesUnlockAccessor) ByUnlockID(identifier float64) (PlanetFesUnlock, error) {
	if a._dataUnlockID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlanetFesUnlock{}, err
			}
		}
		a.GroupData()
	}
	return a._dataUnlockID[identifier], nil
}
