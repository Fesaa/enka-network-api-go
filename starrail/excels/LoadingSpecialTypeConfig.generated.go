package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type LoadingSpecialTypeConfig struct {
	AvailableEntranceIDList []float64 `json:"AvailableEntranceIDList"`
	LoadingType             float64   `json:"LoadingType"`
	Priority                float64   `json:"Priority"`
	SubMissionID            float64   `json:"SubMissionID"`
}
type LoadingSpecialTypeConfigAccessor struct {
	_data             []LoadingSpecialTypeConfig
	_dataSubMissionID map[float64]LoadingSpecialTypeConfig
}

// LoadData retrieves the data. Must be called before LoadingSpecialTypeConfig.GroupData
func (a *LoadingSpecialTypeConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/LoadingSpecialTypeConfig.json")
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
func (a *LoadingSpecialTypeConfigAccessor) Raw() ([]LoadingSpecialTypeConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []LoadingSpecialTypeConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with LoadingSpecialTypeConfigAccessor.LoadData to preload everything
func (a *LoadingSpecialTypeConfigAccessor) GroupData() {
	a._dataSubMissionID = map[float64]LoadingSpecialTypeConfig{}
	for _, d := range a._data {
		a._dataSubMissionID[d.SubMissionID] = d
	}
}

// BySubMissionID returns the LoadingSpecialTypeConfig uniquely identified by SubMissionID
//
// Error is only non-nil if the source errors out
func (a *LoadingSpecialTypeConfigAccessor) BySubMissionID(identifier float64) (LoadingSpecialTypeConfig, error) {
	if a._dataSubMissionID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return LoadingSpecialTypeConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSubMissionID[identifier], nil
}
