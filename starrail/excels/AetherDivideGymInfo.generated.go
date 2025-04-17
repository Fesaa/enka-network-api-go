package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type AetherDivideGymInfo struct {
	ActivityModuleID   json.Number            `json:"ActivityModuleID"`
	BGPath             string                 `json:"BGPath"`
	BadgeUnlockID      json.Number            `json:"BadgeUnlockID"`
	ChallengeQuestList []json.Number          `json:"ChallengeQuestList"`
	Description        map[string]json.Number `json:"Description"`
	DisplayMonsterMap  map[string]json.Number `json:"DisplayMonsterMap"`
	EntranceID         json.Number            `json:"EntranceID"`
	ID                 json.Number            `json:"ID"`
	IconPath           string                 `json:"IconPath"`
	Name               map[string]json.Number `json:"Name"`
	SpiritQuest        json.Number            `json:"SpiritQuest"`
	TabIconPath        string                 `json:"TabIconPath"`
	TrainerQuest       json.Number            `json:"TrainerQuest"`
	UnlockID           json.Number            `json:"UnlockID"`
}
type AetherDivideGymInfoAccessor struct {
	_data              []AetherDivideGymInfo
	_dataBGPath        map[string]AetherDivideGymInfo
	_dataBadgeUnlockID map[json.Number]AetherDivideGymInfo
	_dataEntranceID    map[json.Number]AetherDivideGymInfo
	_dataID            map[json.Number]AetherDivideGymInfo
	_dataIconPath      map[string]AetherDivideGymInfo
	_dataSpiritQuest   map[json.Number]AetherDivideGymInfo
	_dataTabIconPath   map[string]AetherDivideGymInfo
	_dataTrainerQuest  map[json.Number]AetherDivideGymInfo
	_dataUnlockID      map[json.Number]AetherDivideGymInfo
}

// LoadData retrieves the data. Must be called before AetherDivideGymInfo.GroupData
func (a *AetherDivideGymInfoAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AetherDivideGymInfo.json")
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
func (a *AetherDivideGymInfoAccessor) Raw() ([]AetherDivideGymInfo, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AetherDivideGymInfo{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AetherDivideGymInfoAccessor.LoadData to preload everything
func (a *AetherDivideGymInfoAccessor) GroupData() {
	a._dataBGPath = map[string]AetherDivideGymInfo{}
	a._dataBadgeUnlockID = map[json.Number]AetherDivideGymInfo{}
	a._dataEntranceID = map[json.Number]AetherDivideGymInfo{}
	a._dataID = map[json.Number]AetherDivideGymInfo{}
	a._dataIconPath = map[string]AetherDivideGymInfo{}
	a._dataSpiritQuest = map[json.Number]AetherDivideGymInfo{}
	a._dataTabIconPath = map[string]AetherDivideGymInfo{}
	a._dataTrainerQuest = map[json.Number]AetherDivideGymInfo{}
	a._dataUnlockID = map[json.Number]AetherDivideGymInfo{}
	for _, d := range a._data {
		a._dataBGPath[d.BGPath] = d
		a._dataBadgeUnlockID[d.BadgeUnlockID] = d
		a._dataEntranceID[d.EntranceID] = d
		a._dataID[d.ID] = d
		a._dataIconPath[d.IconPath] = d
		a._dataSpiritQuest[d.SpiritQuest] = d
		a._dataTabIconPath[d.TabIconPath] = d
		a._dataTrainerQuest[d.TrainerQuest] = d
		a._dataUnlockID[d.UnlockID] = d
	}
}

// ByBGPath returns the AetherDivideGymInfo uniquely identified by BGPath
//
// Error is only non-nil if the source errors out
func (a *AetherDivideGymInfoAccessor) ByBGPath(identifier string) (AetherDivideGymInfo, error) {
	if a._dataBGPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDivideGymInfo{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBGPath[identifier], nil
}

// ByBadgeUnlockID returns the AetherDivideGymInfo uniquely identified by BadgeUnlockID
//
// Error is only non-nil if the source errors out
func (a *AetherDivideGymInfoAccessor) ByBadgeUnlockID(identifier json.Number) (AetherDivideGymInfo, error) {
	if a._dataBadgeUnlockID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDivideGymInfo{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBadgeUnlockID[identifier], nil
}

// ByEntranceID returns the AetherDivideGymInfo uniquely identified by EntranceID
//
// Error is only non-nil if the source errors out
func (a *AetherDivideGymInfoAccessor) ByEntranceID(identifier json.Number) (AetherDivideGymInfo, error) {
	if a._dataEntranceID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDivideGymInfo{}, err
			}
		}
		a.GroupData()
	}
	return a._dataEntranceID[identifier], nil
}

// ByID returns the AetherDivideGymInfo uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *AetherDivideGymInfoAccessor) ByID(identifier json.Number) (AetherDivideGymInfo, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDivideGymInfo{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByIconPath returns the AetherDivideGymInfo uniquely identified by IconPath
//
// Error is only non-nil if the source errors out
func (a *AetherDivideGymInfoAccessor) ByIconPath(identifier string) (AetherDivideGymInfo, error) {
	if a._dataIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDivideGymInfo{}, err
			}
		}
		a.GroupData()
	}
	return a._dataIconPath[identifier], nil
}

// BySpiritQuest returns the AetherDivideGymInfo uniquely identified by SpiritQuest
//
// Error is only non-nil if the source errors out
func (a *AetherDivideGymInfoAccessor) BySpiritQuest(identifier json.Number) (AetherDivideGymInfo, error) {
	if a._dataSpiritQuest == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDivideGymInfo{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSpiritQuest[identifier], nil
}

// ByTabIconPath returns the AetherDivideGymInfo uniquely identified by TabIconPath
//
// Error is only non-nil if the source errors out
func (a *AetherDivideGymInfoAccessor) ByTabIconPath(identifier string) (AetherDivideGymInfo, error) {
	if a._dataTabIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDivideGymInfo{}, err
			}
		}
		a.GroupData()
	}
	return a._dataTabIconPath[identifier], nil
}

// ByTrainerQuest returns the AetherDivideGymInfo uniquely identified by TrainerQuest
//
// Error is only non-nil if the source errors out
func (a *AetherDivideGymInfoAccessor) ByTrainerQuest(identifier json.Number) (AetherDivideGymInfo, error) {
	if a._dataTrainerQuest == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDivideGymInfo{}, err
			}
		}
		a.GroupData()
	}
	return a._dataTrainerQuest[identifier], nil
}

// ByUnlockID returns the AetherDivideGymInfo uniquely identified by UnlockID
//
// Error is only non-nil if the source errors out
func (a *AetherDivideGymInfoAccessor) ByUnlockID(identifier json.Number) (AetherDivideGymInfo, error) {
	if a._dataUnlockID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDivideGymInfo{}, err
			}
		}
		a.GroupData()
	}
	return a._dataUnlockID[identifier], nil
}
