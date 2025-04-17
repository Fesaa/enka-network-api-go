package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type MaterialSubmitter struct {
	ActivityModuleID float64                         `json:"ActivityModuleID"`
	ID               float64                         `json:"ID"`
	MaterialList     []MaterialSubmitterMaterialList `json:"MaterialList"`
	MissionID        float64                         `json:"MissionID"`
	ParamList        []float64                       `json:"ParamList"`
	RewardID         float64                         `json:"RewardID"`
}
type MaterialSubmitterMaterialList struct {
	ItemID  float64 `json:"ItemID"`
	ItemNum float64 `json:"ItemNum"`
}
type MaterialSubmitterAccessor struct {
	_data                 []MaterialSubmitter
	_dataActivityModuleID map[float64]MaterialSubmitter
	_dataID               map[float64]MaterialSubmitter
	_dataMissionID        map[float64]MaterialSubmitter
	_dataRewardID         map[float64]MaterialSubmitter
}

// LoadData retrieves the data. Must be called before MaterialSubmitter.GroupData
func (a *MaterialSubmitterAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MaterialSubmitter.json")
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
func (a *MaterialSubmitterAccessor) Raw() ([]MaterialSubmitter, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MaterialSubmitter{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MaterialSubmitterAccessor.LoadData to preload everything
func (a *MaterialSubmitterAccessor) GroupData() {
	for _, d := range a._data {
		a._dataActivityModuleID[d.ActivityModuleID] = d
		a._dataID[d.ID] = d
		a._dataMissionID[d.MissionID] = d
		a._dataRewardID[d.RewardID] = d
	}
}

// ByActivityModuleID returns the MaterialSubmitter uniquely identified by ActivityModuleID
//
// Error is only non-nil if the source errors out
func (a *MaterialSubmitterAccessor) ByActivityModuleID(identifier float64) (MaterialSubmitter, error) {
	if a._dataActivityModuleID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MaterialSubmitter{}, err
			}
		}
		a.GroupData()
	}
	return a._dataActivityModuleID[identifier], nil
}

// ByID returns the MaterialSubmitter uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *MaterialSubmitterAccessor) ByID(identifier float64) (MaterialSubmitter, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MaterialSubmitter{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByMissionID returns the MaterialSubmitter uniquely identified by MissionID
//
// Error is only non-nil if the source errors out
func (a *MaterialSubmitterAccessor) ByMissionID(identifier float64) (MaterialSubmitter, error) {
	if a._dataMissionID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MaterialSubmitter{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMissionID[identifier], nil
}

// ByRewardID returns the MaterialSubmitter uniquely identified by RewardID
//
// Error is only non-nil if the source errors out
func (a *MaterialSubmitterAccessor) ByRewardID(identifier float64) (MaterialSubmitter, error) {
	if a._dataRewardID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MaterialSubmitter{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRewardID[identifier], nil
}
