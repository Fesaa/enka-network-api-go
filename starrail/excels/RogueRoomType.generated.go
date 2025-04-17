package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type RogueRoomType struct {
	IsSuper                bool                   `json:"IsSuper"`
	MapShowType            bool                   `json:"MapShowType"`
	RogueRoomType          json.Number            `json:"RogueRoomType"`
	RogueRoomTypeIcon      string                 `json:"RogueRoomTypeIcon"`
	RogueRoomTypeTextmapID map[string]json.Number `json:"RogueRoomTypeTextmapID"`
	RoomIconEffect         string                 `json:"RoomIconEffect"`
	RoomTypeDescTextmapID  map[string]json.Number `json:"RoomTypeDescTextmapID"`
	RoomTypeDescTextmapID2 map[string]json.Number `json:"RoomTypeDescTextmapID2"`
}
type RogueRoomTypeAccessor struct {
	_data []RogueRoomType
}

// LoadData retrieves the data. Must be called before RogueRoomType.GroupData
func (a *RogueRoomTypeAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueRoomType.json")
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
func (a *RogueRoomTypeAccessor) Raw() ([]RogueRoomType, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueRoomType{}, err
		}
	}
	return a._data, nil
}
