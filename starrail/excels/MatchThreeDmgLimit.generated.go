package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type MatchThreeDmgLimit struct {
	BasicDamage float64 `json:"BasicDamage"`
	MaxDamage   float64 `json:"MaxDamage"`
	Round       float64 `json:"Round"`
}
type MatchThreeDmgLimitAccessor struct {
	_data      []MatchThreeDmgLimit
	_dataRound map[float64]MatchThreeDmgLimit
}

// LoadData retrieves the data. Must be called before MatchThreeDmgLimit.GroupData
func (a *MatchThreeDmgLimitAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MatchThreeDmgLimit.json")
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
func (a *MatchThreeDmgLimitAccessor) Raw() ([]MatchThreeDmgLimit, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MatchThreeDmgLimit{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MatchThreeDmgLimitAccessor.LoadData to preload everything
func (a *MatchThreeDmgLimitAccessor) GroupData() {
	for _, d := range a._data {
		a._dataRound[d.Round] = d
	}
}

// ByRound returns the MatchThreeDmgLimit uniquely identified by Round
//
// Error is only non-nil if the source errors out
func (a *MatchThreeDmgLimitAccessor) ByRound(identifier float64) (MatchThreeDmgLimit, error) {
	if a._dataRound == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MatchThreeDmgLimit{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRound[identifier], nil
}
