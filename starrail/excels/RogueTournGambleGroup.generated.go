package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type RogueTournGambleGroup struct {
	GambleGroupID    float64 `json:"GambleGroupID"`
	GambleGroupIcon  string  `json:"GambleGroupIcon"`
	GambleGroupLevel string  `json:"GambleGroupLevel"`
	GambleGroupType  string  `json:"GambleGroupType"`
}
type RogueTournGambleGroupAccessor struct {
	_data              []RogueTournGambleGroup
	_dataGambleGroupID map[float64]RogueTournGambleGroup
}

// LoadData retrieves the data. Must be called before RogueTournGambleGroup.GroupData
func (a *RogueTournGambleGroupAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueTournGambleGroup.json")
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
func (a *RogueTournGambleGroupAccessor) Raw() ([]RogueTournGambleGroup, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueTournGambleGroup{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueTournGambleGroupAccessor.LoadData to preload everything
func (a *RogueTournGambleGroupAccessor) GroupData() {
	for _, d := range a._data {
		a._dataGambleGroupID[d.GambleGroupID] = d
	}
}

// ByGambleGroupID returns the RogueTournGambleGroup uniquely identified by GambleGroupID
//
// Error is only non-nil if the source errors out
func (a *RogueTournGambleGroupAccessor) ByGambleGroupID(identifier float64) (RogueTournGambleGroup, error) {
	if a._dataGambleGroupID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueTournGambleGroup{}, err
			}
		}
		a.GroupData()
	}
	return a._dataGambleGroupID[identifier], nil
}
