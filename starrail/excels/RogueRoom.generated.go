package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type RogueRoom struct {
	GroupID           json.Number            `json:"GroupID"`
	GroupWithContent  map[string]json.Number `json:"GroupWithContent"`
	MapEntrance       json.Number            `json:"MapEntrance"`
	RogueRoomID       json.Number            `json:"RogueRoomID"`
	RogueRoomSections []json.Number          `json:"RogueRoomSections"`
	RogueRoomType     json.Number            `json:"RogueRoomType"`
}
type RogueRoomAccessor struct {
	_data            []RogueRoom
	_dataRogueRoomID map[json.Number]RogueRoom
}

// LoadData retrieves the data. Must be called before RogueRoom.GroupData
func (a *RogueRoomAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueRoom.json")
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
func (a *RogueRoomAccessor) Raw() ([]RogueRoom, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueRoom{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueRoomAccessor.LoadData to preload everything
func (a *RogueRoomAccessor) GroupData() {
	a._dataRogueRoomID = map[json.Number]RogueRoom{}
	for _, d := range a._data {
		a._dataRogueRoomID[d.RogueRoomID] = d
	}
}

// ByRogueRoomID returns the RogueRoom uniquely identified by RogueRoomID
//
// Error is only non-nil if the source errors out
func (a *RogueRoomAccessor) ByRogueRoomID(identifier json.Number) (RogueRoom, error) {
	if a._dataRogueRoomID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueRoom{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRogueRoomID[identifier], nil
}
