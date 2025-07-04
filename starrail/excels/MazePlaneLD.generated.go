package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type MazePlaneLD struct {
	FloorIDList  []float64 `json:"FloorIDList"`
	MazePoolType float64   `json:"MazePoolType"`
	PlaneID      float64   `json:"PlaneID"`
	PlaneName    hash.Hash `json:"PlaneName"`
	PlaneType    string    `json:"PlaneType"`
	StartFloorID float64   `json:"StartFloorID"`
	SubType      float64   `json:"SubType"`
	WorldID      float64   `json:"WorldID"`
}
type MazePlaneLDAccessor struct {
	_data             []MazePlaneLD
	_dataPlaneID      map[float64]MazePlaneLD
	_dataStartFloorID map[float64]MazePlaneLD
}

// LoadData retrieves the data. Must be called before MazePlaneLD.GroupData
func (a *MazePlaneLDAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MazePlaneLD.json")
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
func (a *MazePlaneLDAccessor) Raw() ([]MazePlaneLD, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MazePlaneLD{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MazePlaneLDAccessor.LoadData to preload everything
func (a *MazePlaneLDAccessor) GroupData() {
	a._dataPlaneID = map[float64]MazePlaneLD{}
	a._dataStartFloorID = map[float64]MazePlaneLD{}
	for _, d := range a._data {
		a._dataPlaneID[d.PlaneID] = d
		a._dataStartFloorID[d.StartFloorID] = d
	}
}

// ByPlaneID returns the MazePlaneLD uniquely identified by PlaneID
//
// Error is only non-nil if the source errors out
func (a *MazePlaneLDAccessor) ByPlaneID(identifier float64) (MazePlaneLD, error) {
	if a._dataPlaneID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MazePlaneLD{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPlaneID[identifier], nil
}

// ByStartFloorID returns the MazePlaneLD uniquely identified by StartFloorID
//
// Error is only non-nil if the source errors out
func (a *MazePlaneLDAccessor) ByStartFloorID(identifier float64) (MazePlaneLD, error) {
	if a._dataStartFloorID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MazePlaneLD{}, err
			}
		}
		a.GroupData()
	}
	return a._dataStartFloorID[identifier], nil
}
