package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type TrainPartyMTRank struct {
	Rank           float64   `json:"Rank"`
	RankName       hash.Hash `json:"RankName"`
	RankNum        float64   `json:"RankNum"`
	RankPrefabPath string    `json:"RankPrefabPath"`
	RankScore      float64   `json:"RankScore"`
}
type TrainPartyMTRankAccessor struct {
	_data     []TrainPartyMTRank
	_dataRank map[float64]TrainPartyMTRank
}

// LoadData retrieves the data. Must be called before TrainPartyMTRank.GroupData
func (a *TrainPartyMTRankAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TrainPartyMTRank.json")
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
func (a *TrainPartyMTRankAccessor) Raw() ([]TrainPartyMTRank, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TrainPartyMTRank{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with TrainPartyMTRankAccessor.LoadData to preload everything
func (a *TrainPartyMTRankAccessor) GroupData() {
	a._dataRank = map[float64]TrainPartyMTRank{}
	for _, d := range a._data {
		a._dataRank[d.Rank] = d
	}
}

// ByRank returns the TrainPartyMTRank uniquely identified by Rank
//
// Error is only non-nil if the source errors out
func (a *TrainPartyMTRankAccessor) ByRank(identifier float64) (TrainPartyMTRank, error) {
	if a._dataRank == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TrainPartyMTRank{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRank[identifier], nil
}
