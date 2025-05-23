package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type PerformanceRecallData struct {
	CategoryID      float64                                `json:"CategoryID"`
	ID              float64                                `json:"ID"`
	ImgHeightSize   float64                                `json:"ImgHeightSize"`
	ImgPath         string                                 `json:"ImgPath"`
	ImgPathWall     string                                 `json:"ImgPathWall"`
	ImgPathWallF    string                                 `json:"ImgPathWall_F"`
	ImgPathF        string                                 `json:"ImgPath_F"`
	Name            hash.Hash                              `json:"Name"`
	PerformanceID   float64                                `json:"PerformanceID"`
	UnlockCondition []PerformanceRecallDataUnlockCondition `json:"UnlockCondition"`
	IsVideo         bool                                   `json:"isVideo"`
}
type PerformanceRecallDataUnlockCondition struct {
	Param string `json:"Param"`
	Type  string `json:"Type"`
}
type PerformanceRecallDataAccessor struct {
	_data   []PerformanceRecallData
	_dataID map[float64]PerformanceRecallData
}

// LoadData retrieves the data. Must be called before PerformanceRecallData.GroupData
func (a *PerformanceRecallDataAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PerformanceRecallData.json")
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
func (a *PerformanceRecallDataAccessor) Raw() ([]PerformanceRecallData, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PerformanceRecallData{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PerformanceRecallDataAccessor.LoadData to preload everything
func (a *PerformanceRecallDataAccessor) GroupData() {
	a._dataID = map[float64]PerformanceRecallData{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the PerformanceRecallData uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *PerformanceRecallDataAccessor) ByID(identifier float64) (PerformanceRecallData, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PerformanceRecallData{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
