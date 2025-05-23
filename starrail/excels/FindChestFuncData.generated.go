package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type FindChestFuncData struct {
	ChestTypeList      []string  `json:"ChestTypeList"`
	FindNum            float64   `json:"FindNum"`
	FuncID             float64   `json:"FuncID"`
	GameModeList       []string  `json:"GameModeList"`
	MapIconID          float64   `json:"MapIconID"`
	MappingInfoID      float64   `json:"MappingInfoID"`
	SpecialMappinginfo float64   `json:"SpecialMappinginfo"`
	TriggerParamList   []float64 `json:"TriggerParamList"`
	TriggerType        string    `json:"TriggerType"`
	WorldIDList        []float64 `json:"WorldIDList"`
}
type FindChestFuncDataAccessor struct {
	_data                   []FindChestFuncData
	_dataFuncID             map[float64]FindChestFuncData
	_dataMappingInfoID      map[float64]FindChestFuncData
	_dataSpecialMappinginfo map[float64]FindChestFuncData
}

// LoadData retrieves the data. Must be called before FindChestFuncData.GroupData
func (a *FindChestFuncDataAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/FindChestFuncData.json")
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
func (a *FindChestFuncDataAccessor) Raw() ([]FindChestFuncData, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []FindChestFuncData{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with FindChestFuncDataAccessor.LoadData to preload everything
func (a *FindChestFuncDataAccessor) GroupData() {
	a._dataFuncID = map[float64]FindChestFuncData{}
	a._dataMappingInfoID = map[float64]FindChestFuncData{}
	a._dataSpecialMappinginfo = map[float64]FindChestFuncData{}
	for _, d := range a._data {
		a._dataFuncID[d.FuncID] = d
		a._dataMappingInfoID[d.MappingInfoID] = d
		a._dataSpecialMappinginfo[d.SpecialMappinginfo] = d
	}
}

// ByFuncID returns the FindChestFuncData uniquely identified by FuncID
//
// Error is only non-nil if the source errors out
func (a *FindChestFuncDataAccessor) ByFuncID(identifier float64) (FindChestFuncData, error) {
	if a._dataFuncID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return FindChestFuncData{}, err
			}
		}
		a.GroupData()
	}
	return a._dataFuncID[identifier], nil
}

// ByMappingInfoID returns the FindChestFuncData uniquely identified by MappingInfoID
//
// Error is only non-nil if the source errors out
func (a *FindChestFuncDataAccessor) ByMappingInfoID(identifier float64) (FindChestFuncData, error) {
	if a._dataMappingInfoID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return FindChestFuncData{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMappingInfoID[identifier], nil
}

// BySpecialMappinginfo returns the FindChestFuncData uniquely identified by SpecialMappinginfo
//
// Error is only non-nil if the source errors out
func (a *FindChestFuncDataAccessor) BySpecialMappinginfo(identifier float64) (FindChestFuncData, error) {
	if a._dataSpecialMappinginfo == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return FindChestFuncData{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSpecialMappinginfo[identifier], nil
}
