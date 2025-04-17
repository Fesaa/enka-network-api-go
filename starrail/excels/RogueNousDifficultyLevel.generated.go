package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type RogueNousDifficultyLevel struct {
	DifficultyDesc hash.Hash    `json:"DifficultyDesc"`
	DifficultyID   float64      `json:"DifficultyID"`
	DifficultyType string       `json:"DifficultyType"`
	ParamList      []hash.Value `json:"ParamList"`
	Sort           float64      `json:"Sort"`
	Tag            float64      `json:"Tag"`
}
type RogueNousDifficultyLevelAccessor struct {
	_data             []RogueNousDifficultyLevel
	_dataDifficultyID map[float64]RogueNousDifficultyLevel
}

// LoadData retrieves the data. Must be called before RogueNousDifficultyLevel.GroupData
func (a *RogueNousDifficultyLevelAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueNousDifficultyLevel.json")
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
func (a *RogueNousDifficultyLevelAccessor) Raw() ([]RogueNousDifficultyLevel, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueNousDifficultyLevel{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueNousDifficultyLevelAccessor.LoadData to preload everything
func (a *RogueNousDifficultyLevelAccessor) GroupData() {
	a._dataDifficultyID = map[float64]RogueNousDifficultyLevel{}
	for _, d := range a._data {
		a._dataDifficultyID[d.DifficultyID] = d
	}
}

// ByDifficultyID returns the RogueNousDifficultyLevel uniquely identified by DifficultyID
//
// Error is only non-nil if the source errors out
func (a *RogueNousDifficultyLevelAccessor) ByDifficultyID(identifier float64) (RogueNousDifficultyLevel, error) {
	if a._dataDifficultyID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueNousDifficultyLevel{}, err
			}
		}
		a.GroupData()
	}
	return a._dataDifficultyID[identifier], nil
}
