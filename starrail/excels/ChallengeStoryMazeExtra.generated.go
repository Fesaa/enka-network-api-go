package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ChallengeStoryMazeExtra struct {
	BattleTargetID []float64 `json:"BattleTargetID"`
	ClearScore     float64   `json:"ClearScore"`
	ID             float64   `json:"ID"`
	TurnLimit      float64   `json:"TurnLimit"`
}
type ChallengeStoryMazeExtraAccessor struct {
	_data   []ChallengeStoryMazeExtra
	_dataID map[float64]ChallengeStoryMazeExtra
}

// LoadData retrieves the data. Must be called before ChallengeStoryMazeExtra.GroupData
func (a *ChallengeStoryMazeExtraAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ChallengeStoryMazeExtra.json")
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
func (a *ChallengeStoryMazeExtraAccessor) Raw() ([]ChallengeStoryMazeExtra, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ChallengeStoryMazeExtra{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ChallengeStoryMazeExtraAccessor.LoadData to preload everything
func (a *ChallengeStoryMazeExtraAccessor) GroupData() {
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the ChallengeStoryMazeExtra uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *ChallengeStoryMazeExtraAccessor) ByID(identifier float64) (ChallengeStoryMazeExtra, error) {
	if a._dataID == nil {
		err := a.LoadData()
		if err != nil {
			return ChallengeStoryMazeExtra{}, err
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
