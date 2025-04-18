package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type SpecialChestFindData struct {
	FloorID                 float64 `json:"FloorID"`
	GroupID                 float64 `json:"GroupID"`
	InstanceID              float64 `json:"InstanceID"`
	IsUseSpecialMappinginfo bool    `json:"IsUseSpecialMappinginfo"`
	PlaneID                 float64 `json:"PlaneID"`
	ReplaceGroupID          float64 `json:"ReplaceGroupID"`
	ReplaceInstanceID       float64 `json:"ReplaceInstanceID"`
	ReplaceType             string  `json:"ReplaceType"`
}
type SpecialChestFindDataAccessor struct {
	_data []SpecialChestFindData
}

// LoadData retrieves the data. Must be called before SpecialChestFindData.GroupData
func (a *SpecialChestFindDataAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/SpecialChestFindData.json")
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
func (a *SpecialChestFindDataAccessor) Raw() ([]SpecialChestFindData, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []SpecialChestFindData{}, err
		}
	}
	return a._data, nil
}
