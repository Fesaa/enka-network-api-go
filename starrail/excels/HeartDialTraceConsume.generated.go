package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type HeartDialTraceConsume struct {
	FloorID          float64                             `json:"FloorID"`
	HeartDialEmotion string                              `json:"HeartDialEmotion"`
	HeartDialTraceID float64                             `json:"HeartDialTraceID"`
	MapInfoID        float64                             `json:"MapInfoID"`
	MaterialCost     []HeartDialTraceConsumeMaterialCost `json:"MaterialCost"`
	MiniMapID        float64                             `json:"MiniMapID"`
}
type HeartDialTraceConsumeMaterialCost struct {
	ItemID  float64 `json:"ItemID"`
	ItemNum float64 `json:"ItemNum"`
}
type HeartDialTraceConsumeAccessor struct {
	_data                 []HeartDialTraceConsume
	_dataHeartDialTraceID map[float64]HeartDialTraceConsume
}

// LoadData retrieves the data. Must be called before HeartDialTraceConsume.GroupData
func (a *HeartDialTraceConsumeAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/HeartDialTraceConsume.json")
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
func (a *HeartDialTraceConsumeAccessor) Raw() ([]HeartDialTraceConsume, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []HeartDialTraceConsume{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with HeartDialTraceConsumeAccessor.LoadData to preload everything
func (a *HeartDialTraceConsumeAccessor) GroupData() {
	for _, d := range a._data {
		a._dataHeartDialTraceID[d.HeartDialTraceID] = d
	}
}

// ByHeartDialTraceID returns the HeartDialTraceConsume uniquely identified by HeartDialTraceID
//
// Error is only non-nil if the source errors out
func (a *HeartDialTraceConsumeAccessor) ByHeartDialTraceID(identifier float64) (HeartDialTraceConsume, error) {
	if a._dataHeartDialTraceID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return HeartDialTraceConsume{}, err
			}
		}
		a.GroupData()
	}
	return a._dataHeartDialTraceID[identifier], nil
}
