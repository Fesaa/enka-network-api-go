package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type HeliobusSkill struct {
	BGDescription        hash.Hash `json:"BGDescription"`
	HeliobusSkillID      float64   `json:"HeliobusSkillID"`
	RelatedEventID       float64   `json:"RelatedEventID"`
	SkillEffect          string    `json:"SkillEffect"`
	SkillIconPath        string    `json:"SkillIconPath"`
	SkillUIPosition      float64   `json:"SkillUIPosition"`
	UnlockMissionID      float64   `json:"UnlockMissionID"`
	UnlockToastMissionID float64   `json:"UnlockToastMissionID"`
}
type HeliobusSkillAccessor struct {
	_data                     []HeliobusSkill
	_dataHeliobusSkillID      map[float64]HeliobusSkill
	_dataRelatedEventID       map[float64]HeliobusSkill
	_dataSkillIconPath        map[string]HeliobusSkill
	_dataSkillUIPosition      map[float64]HeliobusSkill
	_dataUnlockMissionID      map[float64]HeliobusSkill
	_dataUnlockToastMissionID map[float64]HeliobusSkill
}

// LoadData retrieves the data. Must be called before HeliobusSkill.GroupData
func (a *HeliobusSkillAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/HeliobusSkill.json")
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
func (a *HeliobusSkillAccessor) Raw() ([]HeliobusSkill, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []HeliobusSkill{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with HeliobusSkillAccessor.LoadData to preload everything
func (a *HeliobusSkillAccessor) GroupData() {
	a._dataHeliobusSkillID = map[float64]HeliobusSkill{}
	a._dataRelatedEventID = map[float64]HeliobusSkill{}
	a._dataSkillIconPath = map[string]HeliobusSkill{}
	a._dataSkillUIPosition = map[float64]HeliobusSkill{}
	a._dataUnlockMissionID = map[float64]HeliobusSkill{}
	a._dataUnlockToastMissionID = map[float64]HeliobusSkill{}
	for _, d := range a._data {
		a._dataHeliobusSkillID[d.HeliobusSkillID] = d
		a._dataRelatedEventID[d.RelatedEventID] = d
		a._dataSkillIconPath[d.SkillIconPath] = d
		a._dataSkillUIPosition[d.SkillUIPosition] = d
		a._dataUnlockMissionID[d.UnlockMissionID] = d
		a._dataUnlockToastMissionID[d.UnlockToastMissionID] = d
	}
}

// ByHeliobusSkillID returns the HeliobusSkill uniquely identified by HeliobusSkillID
//
// Error is only non-nil if the source errors out
func (a *HeliobusSkillAccessor) ByHeliobusSkillID(identifier float64) (HeliobusSkill, error) {
	if a._dataHeliobusSkillID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return HeliobusSkill{}, err
			}
		}
		a.GroupData()
	}
	return a._dataHeliobusSkillID[identifier], nil
}

// ByRelatedEventID returns the HeliobusSkill uniquely identified by RelatedEventID
//
// Error is only non-nil if the source errors out
func (a *HeliobusSkillAccessor) ByRelatedEventID(identifier float64) (HeliobusSkill, error) {
	if a._dataRelatedEventID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return HeliobusSkill{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRelatedEventID[identifier], nil
}

// BySkillIconPath returns the HeliobusSkill uniquely identified by SkillIconPath
//
// Error is only non-nil if the source errors out
func (a *HeliobusSkillAccessor) BySkillIconPath(identifier string) (HeliobusSkill, error) {
	if a._dataSkillIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return HeliobusSkill{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSkillIconPath[identifier], nil
}

// BySkillUIPosition returns the HeliobusSkill uniquely identified by SkillUIPosition
//
// Error is only non-nil if the source errors out
func (a *HeliobusSkillAccessor) BySkillUIPosition(identifier float64) (HeliobusSkill, error) {
	if a._dataSkillUIPosition == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return HeliobusSkill{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSkillUIPosition[identifier], nil
}

// ByUnlockMissionID returns the HeliobusSkill uniquely identified by UnlockMissionID
//
// Error is only non-nil if the source errors out
func (a *HeliobusSkillAccessor) ByUnlockMissionID(identifier float64) (HeliobusSkill, error) {
	if a._dataUnlockMissionID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return HeliobusSkill{}, err
			}
		}
		a.GroupData()
	}
	return a._dataUnlockMissionID[identifier], nil
}

// ByUnlockToastMissionID returns the HeliobusSkill uniquely identified by UnlockToastMissionID
//
// Error is only non-nil if the source errors out
func (a *HeliobusSkillAccessor) ByUnlockToastMissionID(identifier float64) (HeliobusSkill, error) {
	if a._dataUnlockToastMissionID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return HeliobusSkill{}, err
			}
		}
		a.GroupData()
	}
	return a._dataUnlockToastMissionID[identifier], nil
}
