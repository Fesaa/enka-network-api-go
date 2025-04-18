package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type RogueMagicAdventureRoom struct {
	AdventureType string  `json:"AdventureType"`
	ParamGroupID  float64 `json:"ParamGroupID"`
	RoomID        float64 `json:"RoomID"`
}
type RogueMagicAdventureRoomAccessor struct {
	_data       []RogueMagicAdventureRoom
	_dataRoomID map[float64]RogueMagicAdventureRoom
}

// LoadData retrieves the data. Must be called before RogueMagicAdventureRoom.GroupData
func (a *RogueMagicAdventureRoomAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueMagicAdventureRoom.json")
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
func (a *RogueMagicAdventureRoomAccessor) Raw() ([]RogueMagicAdventureRoom, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueMagicAdventureRoom{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueMagicAdventureRoomAccessor.LoadData to preload everything
func (a *RogueMagicAdventureRoomAccessor) GroupData() {
	a._dataRoomID = map[float64]RogueMagicAdventureRoom{}
	for _, d := range a._data {
		a._dataRoomID[d.RoomID] = d
	}
}

// ByRoomID returns the RogueMagicAdventureRoom uniquely identified by RoomID
//
// Error is only non-nil if the source errors out
func (a *RogueMagicAdventureRoomAccessor) ByRoomID(identifier float64) (RogueMagicAdventureRoom, error) {
	if a._dataRoomID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueMagicAdventureRoom{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRoomID[identifier], nil
}
