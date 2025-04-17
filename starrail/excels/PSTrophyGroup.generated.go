package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type PSTrophyGroup struct {
	PSTrophyGroup float64   `json:"PSTrophyGroup"`
	TrophyGroup   hash.Hash `json:"TrophyGroup"`
}
type PSTrophyGroupAccessor struct {
	_data              []PSTrophyGroup
	_dataPSTrophyGroup map[float64]PSTrophyGroup
}

// LoadData retrieves the data. Must be called before PSTrophyGroup.GroupData
func (a *PSTrophyGroupAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PSTrophyGroup.json")
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
func (a *PSTrophyGroupAccessor) Raw() ([]PSTrophyGroup, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PSTrophyGroup{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PSTrophyGroupAccessor.LoadData to preload everything
func (a *PSTrophyGroupAccessor) GroupData() {
	a._dataPSTrophyGroup = map[float64]PSTrophyGroup{}
	for _, d := range a._data {
		a._dataPSTrophyGroup[d.PSTrophyGroup] = d
	}
}

// ByPSTrophyGroup returns the PSTrophyGroup uniquely identified by PSTrophyGroup
//
// Error is only non-nil if the source errors out
func (a *PSTrophyGroupAccessor) ByPSTrophyGroup(identifier float64) (PSTrophyGroup, error) {
	if a._dataPSTrophyGroup == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PSTrophyGroup{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPSTrophyGroup[identifier], nil
}
