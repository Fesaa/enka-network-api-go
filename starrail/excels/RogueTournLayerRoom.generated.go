package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type RogueTournLayerRoom struct {
	Door1     map[string]json.Number `json:"Door1"`
	Door2     map[string]json.Number `json:"Door2"`
	Door3     map[string]json.Number `json:"Door3"`
	LayerID   json.Number            `json:"LayerID"`
	RoomIndex json.Number            `json:"RoomIndex"`
}
type RogueTournLayerRoomAccessor struct {
	_data []RogueTournLayerRoom
}

// LoadData retrieves the data. Must be called before RogueTournLayerRoom.GroupData
func (a *RogueTournLayerRoomAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueTournLayerRoom.json")
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
func (a *RogueTournLayerRoomAccessor) Raw() ([]RogueTournLayerRoom, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueTournLayerRoom{}, err
		}
	}
	return a._data, nil
}
