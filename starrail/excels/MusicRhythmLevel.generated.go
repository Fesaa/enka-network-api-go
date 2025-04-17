package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type MusicRhythmLevel struct {
	Difficulty       float64   `json:"Difficulty"`
	EnterType        float64   `json:"EnterType"`
	FeverComboCount  float64   `json:"FeverComboCount"`
	Group            float64   `json:"Group"`
	ID               float64   `json:"ID"`
	InputScore       []float64 `json:"InputScore"`
	StarRewardIDList []float64 `json:"StarRewardIDList"`
	StarScoreList    []float64 `json:"StarScoreList"`
}
type MusicRhythmLevelAccessor struct {
	_data   []MusicRhythmLevel
	_dataID map[float64]MusicRhythmLevel
}

// LoadData retrieves the data. Must be called before MusicRhythmLevel.GroupData
func (a *MusicRhythmLevelAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MusicRhythmLevel.json")
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
func (a *MusicRhythmLevelAccessor) Raw() ([]MusicRhythmLevel, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MusicRhythmLevel{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MusicRhythmLevelAccessor.LoadData to preload everything
func (a *MusicRhythmLevelAccessor) GroupData() {
	a._dataID = map[float64]MusicRhythmLevel{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the MusicRhythmLevel uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *MusicRhythmLevelAccessor) ByID(identifier float64) (MusicRhythmLevel, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MusicRhythmLevel{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
