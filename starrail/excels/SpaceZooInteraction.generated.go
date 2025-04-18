package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type SpaceZooInteraction struct {
	Case          string                   `json:"Case"`
	ID            float64                  `json:"ID"`
	Param         SpaceZooInteractionParam `json:"Param"`
	PerformanceID float64                  `json:"PerformanceID"`
	Priority      float64                  `json:"Priority"`
	RoomID        float64                  `json:"RoomID"`
}
type SpaceZooInteractionParam struct {
	ArrayValue  []hash.StringValue `json:"ArrayValue"`
	IntValue    float64            `json:"IntValue"`
	StringValue string             `json:"StringValue"`
}
type SpaceZooInteractionAccessor struct {
	_data   []SpaceZooInteraction
	_dataID map[float64]SpaceZooInteraction
}

// LoadData retrieves the data. Must be called before SpaceZooInteraction.GroupData
func (a *SpaceZooInteractionAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/SpaceZooInteraction.json")
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
func (a *SpaceZooInteractionAccessor) Raw() ([]SpaceZooInteraction, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []SpaceZooInteraction{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with SpaceZooInteractionAccessor.LoadData to preload everything
func (a *SpaceZooInteractionAccessor) GroupData() {
	a._dataID = map[float64]SpaceZooInteraction{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the SpaceZooInteraction uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *SpaceZooInteractionAccessor) ByID(identifier float64) (SpaceZooInteraction, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return SpaceZooInteraction{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
