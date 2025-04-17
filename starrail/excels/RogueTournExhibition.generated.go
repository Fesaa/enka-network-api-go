package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type RogueTournExhibition struct {
	ExhibitionID   json.Number `json:"ExhibitionID"`
	ExhibitionType string      `json:"ExhibitionType"`
	IconPath       string      `json:"IconPath"`
	ImagePath      string      `json:"ImagePath"`
	ProgramGroupID json.Number `json:"ProgramGroupID"`
	SlotIconPath   string      `json:"SlotIconPath"`
}
type RogueTournExhibitionAccessor struct {
	_data             []RogueTournExhibition
	_dataIconPath     map[string]RogueTournExhibition
	_dataImagePath    map[string]RogueTournExhibition
	_dataSlotIconPath map[string]RogueTournExhibition
}

// LoadData retrieves the data. Must be called before RogueTournExhibition.GroupData
func (a *RogueTournExhibitionAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueTournExhibition.json")
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
func (a *RogueTournExhibitionAccessor) Raw() ([]RogueTournExhibition, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueTournExhibition{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueTournExhibitionAccessor.LoadData to preload everything
func (a *RogueTournExhibitionAccessor) GroupData() {
	a._dataIconPath = map[string]RogueTournExhibition{}
	a._dataImagePath = map[string]RogueTournExhibition{}
	a._dataSlotIconPath = map[string]RogueTournExhibition{}
	for _, d := range a._data {
		a._dataIconPath[d.IconPath] = d
		a._dataImagePath[d.ImagePath] = d
		a._dataSlotIconPath[d.SlotIconPath] = d
	}
}

// ByIconPath returns the RogueTournExhibition uniquely identified by IconPath
//
// Error is only non-nil if the source errors out
func (a *RogueTournExhibitionAccessor) ByIconPath(identifier string) (RogueTournExhibition, error) {
	if a._dataIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueTournExhibition{}, err
			}
		}
		a.GroupData()
	}
	return a._dataIconPath[identifier], nil
}

// ByImagePath returns the RogueTournExhibition uniquely identified by ImagePath
//
// Error is only non-nil if the source errors out
func (a *RogueTournExhibitionAccessor) ByImagePath(identifier string) (RogueTournExhibition, error) {
	if a._dataImagePath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueTournExhibition{}, err
			}
		}
		a.GroupData()
	}
	return a._dataImagePath[identifier], nil
}

// BySlotIconPath returns the RogueTournExhibition uniquely identified by SlotIconPath
//
// Error is only non-nil if the source errors out
func (a *RogueTournExhibitionAccessor) BySlotIconPath(identifier string) (RogueTournExhibition, error) {
	if a._dataSlotIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueTournExhibition{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSlotIconPath[identifier], nil
}
