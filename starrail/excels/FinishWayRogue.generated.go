package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type FinishWayRogue struct {
	FinishType    string        `json:"FinishType"`
	ID            float64       `json:"ID"`
	IsBackTrack   bool          `json:"IsBackTrack"`
	ParamInt1     float64       `json:"ParamInt1"`
	ParamInt2     float64       `json:"ParamInt2"`
	ParamInt3     float64       `json:"ParamInt3"`
	ParamIntList  []float64     `json:"ParamIntList"`
	ParamItemList []interface{} `json:"ParamItemList"`
	ParamStr1     string        `json:"ParamStr1"`
	ParamType     string        `json:"ParamType"`
	Progress      float64       `json:"Progress"`
}
type FinishWayRogueAccessor struct {
	_data   []FinishWayRogue
	_dataID map[float64]FinishWayRogue
}

// LoadData retrieves the data. Must be called before FinishWayRogue.GroupData
func (a *FinishWayRogueAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/FinishWayRogue.json")
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
func (a *FinishWayRogueAccessor) Raw() ([]FinishWayRogue, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []FinishWayRogue{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with FinishWayRogueAccessor.LoadData to preload everything
func (a *FinishWayRogueAccessor) GroupData() {
	a._dataID = map[float64]FinishWayRogue{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the FinishWayRogue uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *FinishWayRogueAccessor) ByID(identifier float64) (FinishWayRogue, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return FinishWayRogue{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
