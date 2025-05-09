package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type PlayerReturnInvite struct {
	APILabel           string                               `json:"APILabel"`
	ActivityModuleID   float64                              `json:"ActivityModuleID"`
	DisplayRewardItems PlayerReturnInviteDisplayRewardItems `json:"DisplayRewardItems"`
	ID                 float64                              `json:"ID"`
}
type PlayerReturnInviteDisplayRewardItems struct {
	F1 float64 `json:"1"`
}
type PlayerReturnInviteAccessor struct {
	_data                 []PlayerReturnInvite
	_dataAPILabel         map[string]PlayerReturnInvite
	_dataActivityModuleID map[float64]PlayerReturnInvite
	_dataID               map[float64]PlayerReturnInvite
}

// LoadData retrieves the data. Must be called before PlayerReturnInvite.GroupData
func (a *PlayerReturnInviteAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PlayerReturnInvite.json")
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
func (a *PlayerReturnInviteAccessor) Raw() ([]PlayerReturnInvite, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PlayerReturnInvite{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PlayerReturnInviteAccessor.LoadData to preload everything
func (a *PlayerReturnInviteAccessor) GroupData() {
	a._dataAPILabel = map[string]PlayerReturnInvite{}
	a._dataActivityModuleID = map[float64]PlayerReturnInvite{}
	a._dataID = map[float64]PlayerReturnInvite{}
	for _, d := range a._data {
		a._dataAPILabel[d.APILabel] = d
		a._dataActivityModuleID[d.ActivityModuleID] = d
		a._dataID[d.ID] = d
	}
}

// ByAPILabel returns the PlayerReturnInvite uniquely identified by APILabel
//
// Error is only non-nil if the source errors out
func (a *PlayerReturnInviteAccessor) ByAPILabel(identifier string) (PlayerReturnInvite, error) {
	if a._dataAPILabel == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlayerReturnInvite{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAPILabel[identifier], nil
}

// ByActivityModuleID returns the PlayerReturnInvite uniquely identified by ActivityModuleID
//
// Error is only non-nil if the source errors out
func (a *PlayerReturnInviteAccessor) ByActivityModuleID(identifier float64) (PlayerReturnInvite, error) {
	if a._dataActivityModuleID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlayerReturnInvite{}, err
			}
		}
		a.GroupData()
	}
	return a._dataActivityModuleID[identifier], nil
}

// ByID returns the PlayerReturnInvite uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *PlayerReturnInviteAccessor) ByID(identifier float64) (PlayerReturnInvite, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlayerReturnInvite{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
