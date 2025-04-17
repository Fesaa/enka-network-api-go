package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type MapPropConditionConfig struct {
	ActivityModuleID float64                                  `json:"ActivityModuleID"`
	ID               float64                                  `json:"ID"`
	MappingInfoID    float64                                  `json:"MappingInfoID"`
	MiniMapIconID    float64                                  `json:"MiniMapIconID"`
	Priority         float64                                  `json:"Priority"`
	UnloadConditions []interface{}                            `json:"UnloadConditions"`
	UnlockConditions []MapPropConditionConfigUnlockConditions `json:"UnlockConditions"`
}
type MapPropConditionConfigUnlockConditions struct {
	Param string `json:"Param"`
	Type  string `json:"Type"`
}
type MapPropConditionConfigAccessor struct {
	_data                 []MapPropConditionConfig
	_dataActivityModuleID map[float64]MapPropConditionConfig
	_dataID               map[float64]MapPropConditionConfig
	_dataMappingInfoID    map[float64]MapPropConditionConfig
}

// LoadData retrieves the data. Must be called before MapPropConditionConfig.GroupData
func (a *MapPropConditionConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MapPropConditionConfig.json")
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
func (a *MapPropConditionConfigAccessor) Raw() ([]MapPropConditionConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MapPropConditionConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MapPropConditionConfigAccessor.LoadData to preload everything
func (a *MapPropConditionConfigAccessor) GroupData() {
	a._dataActivityModuleID = map[float64]MapPropConditionConfig{}
	a._dataID = map[float64]MapPropConditionConfig{}
	a._dataMappingInfoID = map[float64]MapPropConditionConfig{}
	for _, d := range a._data {
		a._dataActivityModuleID[d.ActivityModuleID] = d
		a._dataID[d.ID] = d
		a._dataMappingInfoID[d.MappingInfoID] = d
	}
}

// ByActivityModuleID returns the MapPropConditionConfig uniquely identified by ActivityModuleID
//
// Error is only non-nil if the source errors out
func (a *MapPropConditionConfigAccessor) ByActivityModuleID(identifier float64) (MapPropConditionConfig, error) {
	if a._dataActivityModuleID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MapPropConditionConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataActivityModuleID[identifier], nil
}

// ByID returns the MapPropConditionConfig uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *MapPropConditionConfigAccessor) ByID(identifier float64) (MapPropConditionConfig, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MapPropConditionConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByMappingInfoID returns the MapPropConditionConfig uniquely identified by MappingInfoID
//
// Error is only non-nil if the source errors out
func (a *MapPropConditionConfigAccessor) ByMappingInfoID(identifier float64) (MapPropConditionConfig, error) {
	if a._dataMappingInfoID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MapPropConditionConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMappingInfoID[identifier], nil
}
