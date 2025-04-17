package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type RogueMagicScore struct {
	LayerNum    float64 `json:"LayerNum"`
	RoomNum     float64 `json:"RoomNum"`
	WeeklyScore float64 `json:"WeeklyScore"`
	WorldLevel  float64 `json:"WorldLevel"`
}
type RogueMagicScoreAccessor struct {
	_data []RogueMagicScore
}

// LoadData retrieves the data. Must be called before RogueMagicScore.GroupData
func (a *RogueMagicScoreAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueMagicScore.json")
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
func (a *RogueMagicScoreAccessor) Raw() ([]RogueMagicScore, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueMagicScore{}, err
		}
	}
	return a._data, nil
}
