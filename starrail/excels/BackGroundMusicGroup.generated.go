package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type BackGroundMusicGroup struct {
	GroupIcon string    `json:"GroupIcon"`
	GroupName hash.Hash `json:"GroupName"`
	ID        float64   `json:"ID"`
}
type BackGroundMusicGroupAccessor struct {
	_data          []BackGroundMusicGroup
	_dataGroupIcon map[string]BackGroundMusicGroup
	_dataID        map[float64]BackGroundMusicGroup
}

// LoadData retrieves the data. Must be called before BackGroundMusicGroup.GroupData
func (a *BackGroundMusicGroupAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/BackGroundMusicGroup.json")
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
func (a *BackGroundMusicGroupAccessor) Raw() ([]BackGroundMusicGroup, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []BackGroundMusicGroup{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with BackGroundMusicGroupAccessor.LoadData to preload everything
func (a *BackGroundMusicGroupAccessor) GroupData() {
	a._dataGroupIcon = map[string]BackGroundMusicGroup{}
	a._dataID = map[float64]BackGroundMusicGroup{}
	for _, d := range a._data {
		a._dataGroupIcon[d.GroupIcon] = d
		a._dataID[d.ID] = d
	}
}

// ByGroupIcon returns the BackGroundMusicGroup uniquely identified by GroupIcon
//
// Error is only non-nil if the source errors out
func (a *BackGroundMusicGroupAccessor) ByGroupIcon(identifier string) (BackGroundMusicGroup, error) {
	if a._dataGroupIcon == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return BackGroundMusicGroup{}, err
			}
		}
		a.GroupData()
	}
	return a._dataGroupIcon[identifier], nil
}

// ByID returns the BackGroundMusicGroup uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *BackGroundMusicGroupAccessor) ByID(identifier float64) (BackGroundMusicGroup, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return BackGroundMusicGroup{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
