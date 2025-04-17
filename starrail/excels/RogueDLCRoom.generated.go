package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type RogueDLCRoom struct {
	RogueRoomID       float64   `json:"RogueRoomID"`
	RogueRoomSections []float64 `json:"RogueRoomSections"`
	RogueSubMode      string    `json:"RogueSubMode"`
}
type RogueDLCRoomAccessor struct {
	_data            []RogueDLCRoom
	_dataRogueRoomID map[float64]RogueDLCRoom
}

// LoadData retrieves the data. Must be called before RogueDLCRoom.GroupData
func (a *RogueDLCRoomAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueDLCRoom.json")
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
func (a *RogueDLCRoomAccessor) Raw() ([]RogueDLCRoom, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueDLCRoom{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueDLCRoomAccessor.LoadData to preload everything
func (a *RogueDLCRoomAccessor) GroupData() {
	for _, d := range a._data {
		a._dataRogueRoomID[d.RogueRoomID] = d
	}
}

// ByRogueRoomID returns the RogueDLCRoom uniquely identified by RogueRoomID
//
// Error is only non-nil if the source errors out
func (a *RogueDLCRoomAccessor) ByRogueRoomID(identifier float64) (RogueDLCRoom, error) {
	if a._dataRogueRoomID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueDLCRoom{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRogueRoomID[identifier], nil
}
