package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type TrainPartyStatusRank struct {
	Rank           float64 `json:"Rank"`
	RankRequireExp float64 `json:"RankRequireExp"`
	RequireValue   float64 `json:"RequireValue"`
}
type TrainPartyStatusRankAccessor struct {
	_data     []TrainPartyStatusRank
	_dataRank map[float64]TrainPartyStatusRank
}

// LoadData retrieves the data. Must be called before TrainPartyStatusRank.GroupData
func (a *TrainPartyStatusRankAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TrainPartyStatusRank.json")
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
func (a *TrainPartyStatusRankAccessor) Raw() ([]TrainPartyStatusRank, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TrainPartyStatusRank{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with TrainPartyStatusRankAccessor.LoadData to preload everything
func (a *TrainPartyStatusRankAccessor) GroupData() {
	for _, d := range a._data {
		a._dataRank[d.Rank] = d
	}
}

// ByRank returns the TrainPartyStatusRank uniquely identified by Rank
//
// Error is only non-nil if the source errors out
func (a *TrainPartyStatusRankAccessor) ByRank(identifier float64) (TrainPartyStatusRank, error) {
	if a._dataRank == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TrainPartyStatusRank{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRank[identifier], nil
}
