package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type RogueAdventureRoom struct {
	AdventureType string      `json:"AdventureType"`
	ParamGroupID  json.Number `json:"ParamGroupID"`
	RoomID        json.Number `json:"RoomID"`
}
type RogueAdventureRoomAccessor struct {
	_data             []RogueAdventureRoom
	_dataParamGroupID map[json.Number]RogueAdventureRoom
	_dataRoomID       map[json.Number]RogueAdventureRoom
}

// LoadData retrieves the data. Must be called before RogueAdventureRoom.GroupData
func (a *RogueAdventureRoomAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueAdventureRoom.json")
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
func (a *RogueAdventureRoomAccessor) Raw() ([]RogueAdventureRoom, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueAdventureRoom{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueAdventureRoomAccessor.LoadData to preload everything
func (a *RogueAdventureRoomAccessor) GroupData() {
	a._dataParamGroupID = map[json.Number]RogueAdventureRoom{}
	a._dataRoomID = map[json.Number]RogueAdventureRoom{}
	for _, d := range a._data {
		a._dataParamGroupID[d.ParamGroupID] = d
		a._dataRoomID[d.RoomID] = d
	}
}

// ByParamGroupID returns the RogueAdventureRoom uniquely identified by ParamGroupID
//
// Error is only non-nil if the source errors out
func (a *RogueAdventureRoomAccessor) ByParamGroupID(identifier json.Number) (RogueAdventureRoom, error) {
	if a._dataParamGroupID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueAdventureRoom{}, err
			}
		}
		a.GroupData()
	}
	return a._dataParamGroupID[identifier], nil
}

// ByRoomID returns the RogueAdventureRoom uniquely identified by RoomID
//
// Error is only non-nil if the source errors out
func (a *RogueAdventureRoomAccessor) ByRoomID(identifier json.Number) (RogueAdventureRoom, error) {
	if a._dataRoomID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueAdventureRoom{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRoomID[identifier], nil
}
