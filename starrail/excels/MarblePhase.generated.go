package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type MarblePhase struct {
	ID   string    `json:"ID"`
	Name hash.Hash `json:"Name"`
}
type MarblePhaseAccessor struct {
	_data   []MarblePhase
	_dataID map[string]MarblePhase
}

// LoadData retrieves the data. Must be called before MarblePhase.GroupData
func (a *MarblePhaseAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MarblePhase.json")
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
func (a *MarblePhaseAccessor) Raw() ([]MarblePhase, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MarblePhase{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MarblePhaseAccessor.LoadData to preload everything
func (a *MarblePhaseAccessor) GroupData() {
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the MarblePhase uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *MarblePhaseAccessor) ByID(identifier string) (MarblePhase, error) {
	if a._dataID == nil {
		err := a.LoadData()
		if err != nil {
			return MarblePhase{}, err
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
