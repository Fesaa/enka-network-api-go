package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type MazeFloorConnectivity struct {
	FromFloorID      float64 `json:"FromFloorID"`
	LockAreaMapID    float64 `json:"LockAreaMapID"`
	ToFloorID        float64 `json:"ToFloorID"`
	WayPointEntityID float64 `json:"WayPointEntityID"`
	WayPointGroupID  float64 `json:"WayPointGroupID"`
}
type MazeFloorConnectivityAccessor struct {
	_data []MazeFloorConnectivity
}

// LoadData retrieves the data. Must be called before MazeFloorConnectivity.GroupData
func (a *MazeFloorConnectivityAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MazeFloorConnectivity.json")
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
func (a *MazeFloorConnectivityAccessor) Raw() ([]MazeFloorConnectivity, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MazeFloorConnectivity{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}
