package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type RogueMagicRoomMark struct {
	MarkType       string                 `json:"MarkType"`
	RoomIconEffect string                 `json:"RoomIconEffect"`
	RoomType       string                 `json:"RoomType"`
	RoomTypeIcon   string                 `json:"RoomTypeIcon"`
	RoomTypeName   map[string]json.Number `json:"RoomTypeName"`
	ToastIcon      string                 `json:"ToastIcon"`
}
type RogueMagicRoomMarkAccessor struct {
	_data []RogueMagicRoomMark
}

// LoadData retrieves the data. Must be called before RogueMagicRoomMark.GroupData
func (a *RogueMagicRoomMarkAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueMagicRoomMark.json")
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
func (a *RogueMagicRoomMarkAccessor) Raw() ([]RogueMagicRoomMark, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueMagicRoomMark{}, err
		}
	}
	return a._data, nil
}
