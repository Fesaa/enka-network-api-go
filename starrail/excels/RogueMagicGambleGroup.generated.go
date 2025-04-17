package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type RogueMagicGambleGroup struct {
	GambleGroupID    float64 `json:"GambleGroupID"`
	GambleGroupIcon  string  `json:"GambleGroupIcon"`
	GambleGroupLevel string  `json:"GambleGroupLevel"`
	GambleGroupType  string  `json:"GambleGroupType"`
}
type RogueMagicGambleGroupAccessor struct {
	_data              []RogueMagicGambleGroup
	_dataGambleGroupID map[float64]RogueMagicGambleGroup
}

// LoadData retrieves the data. Must be called before RogueMagicGambleGroup.GroupData
func (a *RogueMagicGambleGroupAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueMagicGambleGroup.json")
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
func (a *RogueMagicGambleGroupAccessor) Raw() ([]RogueMagicGambleGroup, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueMagicGambleGroup{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueMagicGambleGroupAccessor.LoadData to preload everything
func (a *RogueMagicGambleGroupAccessor) GroupData() {
	a._dataGambleGroupID = map[float64]RogueMagicGambleGroup{}
	for _, d := range a._data {
		a._dataGambleGroupID[d.GambleGroupID] = d
	}
}

// ByGambleGroupID returns the RogueMagicGambleGroup uniquely identified by GambleGroupID
//
// Error is only non-nil if the source errors out
func (a *RogueMagicGambleGroupAccessor) ByGambleGroupID(identifier float64) (RogueMagicGambleGroup, error) {
	if a._dataGambleGroupID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueMagicGambleGroup{}, err
			}
		}
		a.GroupData()
	}
	return a._dataGambleGroupID[identifier], nil
}
