package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type MazePuzzleDollyZoomTeleport struct {
	FloorID            float64 `json:"FloorID"`
	GroupID            float64 `json:"GroupID"`
	ID                 float64 `json:"ID"`
	InstanceIDA        float64 `json:"InstanceIDA"`
	InstanceIDB        float64 `json:"InstanceIDB"`
	OverrideInitFOV    float64 `json:"OverrideInitFOV"`
	OverridePosA       float64 `json:"OverridePosA"`
	OverridePosB       float64 `json:"OverridePosB"`
	OverrideTargetPosA float64 `json:"OverrideTargetPosA"`
	OverrideTargetPosB float64 `json:"OverrideTargetPosB"`
	PuzzlePrefab       string  `json:"PuzzlePrefab"`
}
type MazePuzzleDollyZoomTeleportAccessor struct {
	_data   []MazePuzzleDollyZoomTeleport
	_dataID map[float64]MazePuzzleDollyZoomTeleport
}

// LoadData retrieves the data. Must be called before MazePuzzleDollyZoomTeleport.GroupData
func (a *MazePuzzleDollyZoomTeleportAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MazePuzzleDollyZoomTeleport.json")
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
func (a *MazePuzzleDollyZoomTeleportAccessor) Raw() ([]MazePuzzleDollyZoomTeleport, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MazePuzzleDollyZoomTeleport{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MazePuzzleDollyZoomTeleportAccessor.LoadData to preload everything
func (a *MazePuzzleDollyZoomTeleportAccessor) GroupData() {
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the MazePuzzleDollyZoomTeleport uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *MazePuzzleDollyZoomTeleportAccessor) ByID(identifier float64) (MazePuzzleDollyZoomTeleport, error) {
	if a._dataID == nil {
		err := a.LoadData()
		if err != nil {
			return MazePuzzleDollyZoomTeleport{}, err
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
