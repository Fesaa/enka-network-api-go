package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type MazePuzzleOrigami struct {
	ColonyID           float64  `json:"ColonyID"`
	CreateNpcPropState string   `json:"CreateNpcPropState"`
	FloorID            float64  `json:"FloorID"`
	GroupID            float64  `json:"GroupID"`
	MainPropID         float64  `json:"MainPropID"`
	MainPropStateList  []string `json:"MainPropStateList"`
	NpcGroupID         float64  `json:"NpcGroupID"`
	NpcInstanceID      float64  `json:"NpcInstanceID"`
	SubPropID          float64  `json:"SubPropID"`
}
type MazePuzzleOrigamiAccessor struct {
	_data []MazePuzzleOrigami
}

// LoadData retrieves the data. Must be called before MazePuzzleOrigami.GroupData
func (a *MazePuzzleOrigamiAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MazePuzzleOrigami.json")
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
func (a *MazePuzzleOrigamiAccessor) Raw() ([]MazePuzzleOrigami, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MazePuzzleOrigami{}, err
		}
	}
	return a._data, nil
}
