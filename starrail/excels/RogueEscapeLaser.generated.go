package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type RogueEscapeLaser struct {
	GameTimeperRound float64   `json:"GameTimeperRound"`
	ParamGroupID     float64   `json:"ParamGroupID"`
	PrepareTime      float64   `json:"PrepareTime"`
	ScoreRange       []float64 `json:"ScoreRange"`
	ScoreperRound    []float64 `json:"ScoreperRound"`
	ScoreperWave     float64   `json:"ScoreperWave"`
	TotalRounds      float64   `json:"TotalRounds"`
}
type RogueEscapeLaserAccessor struct {
	_data             []RogueEscapeLaser
	_dataParamGroupID map[float64]RogueEscapeLaser
}

// LoadData retrieves the data. Must be called before RogueEscapeLaser.GroupData
func (a *RogueEscapeLaserAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueEscapeLaser.json")
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
func (a *RogueEscapeLaserAccessor) Raw() ([]RogueEscapeLaser, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueEscapeLaser{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueEscapeLaserAccessor.LoadData to preload everything
func (a *RogueEscapeLaserAccessor) GroupData() {
	for _, d := range a._data {
		a._dataParamGroupID[d.ParamGroupID] = d
	}
}

// ByParamGroupID returns the RogueEscapeLaser uniquely identified by ParamGroupID
//
// Error is only non-nil if the source errors out
func (a *RogueEscapeLaserAccessor) ByParamGroupID(identifier float64) (RogueEscapeLaser, error) {
	if a._dataParamGroupID == nil {
		err := a.LoadData()
		if err != nil {
			return RogueEscapeLaser{}, err
		}
		a.GroupData()
	}
	return a._dataParamGroupID[identifier], nil
}
