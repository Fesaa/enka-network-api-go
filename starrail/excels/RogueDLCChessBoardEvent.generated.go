package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type RogueDLCChessBoardEvent struct {
	ChessBoardEventDesc map[string]json.Number `json:"ChessBoardEventDesc"`
	ChessBoardEventID   json.Number            `json:"ChessBoardEventID"`
	ChessBoardEventName map[string]json.Number `json:"ChessBoardEventName"`
}
type RogueDLCChessBoardEventAccessor struct {
	_data []RogueDLCChessBoardEvent
}

// LoadData retrieves the data. Must be called before RogueDLCChessBoardEvent.GroupData
func (a *RogueDLCChessBoardEventAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueDLCChessBoardEvent.json")
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
func (a *RogueDLCChessBoardEventAccessor) Raw() ([]RogueDLCChessBoardEvent, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueDLCChessBoardEvent{}, err
		}
	}
	return a._data, nil
}
