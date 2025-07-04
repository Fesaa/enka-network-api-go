package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type MatchThreeV2PVPRank struct {
	BigIconPath   string    `json:"BigIconPath"`
	IconPath      string    `json:"IconPath"`
	MaxScore      float64   `json:"MaxScore"`
	Name          hash.Hash `json:"Name"`
	RankID        float64   `json:"RankID"`
	SmallIconPath string    `json:"SmallIconPath"`
}
type MatchThreeV2PVPRankAccessor struct {
	_data              []MatchThreeV2PVPRank
	_dataBigIconPath   map[string]MatchThreeV2PVPRank
	_dataIconPath      map[string]MatchThreeV2PVPRank
	_dataMaxScore      map[float64]MatchThreeV2PVPRank
	_dataRankID        map[float64]MatchThreeV2PVPRank
	_dataSmallIconPath map[string]MatchThreeV2PVPRank
}

// LoadData retrieves the data. Must be called before MatchThreeV2PVPRank.GroupData
func (a *MatchThreeV2PVPRankAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MatchThreeV2PVPRank.json")
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
func (a *MatchThreeV2PVPRankAccessor) Raw() ([]MatchThreeV2PVPRank, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MatchThreeV2PVPRank{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MatchThreeV2PVPRankAccessor.LoadData to preload everything
func (a *MatchThreeV2PVPRankAccessor) GroupData() {
	a._dataBigIconPath = map[string]MatchThreeV2PVPRank{}
	a._dataIconPath = map[string]MatchThreeV2PVPRank{}
	a._dataMaxScore = map[float64]MatchThreeV2PVPRank{}
	a._dataRankID = map[float64]MatchThreeV2PVPRank{}
	a._dataSmallIconPath = map[string]MatchThreeV2PVPRank{}
	for _, d := range a._data {
		a._dataBigIconPath[d.BigIconPath] = d
		a._dataIconPath[d.IconPath] = d
		a._dataMaxScore[d.MaxScore] = d
		a._dataRankID[d.RankID] = d
		a._dataSmallIconPath[d.SmallIconPath] = d
	}
}

// ByBigIconPath returns the MatchThreeV2PVPRank uniquely identified by BigIconPath
//
// Error is only non-nil if the source errors out
func (a *MatchThreeV2PVPRankAccessor) ByBigIconPath(identifier string) (MatchThreeV2PVPRank, error) {
	if a._dataBigIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MatchThreeV2PVPRank{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBigIconPath[identifier], nil
}

// ByIconPath returns the MatchThreeV2PVPRank uniquely identified by IconPath
//
// Error is only non-nil if the source errors out
func (a *MatchThreeV2PVPRankAccessor) ByIconPath(identifier string) (MatchThreeV2PVPRank, error) {
	if a._dataIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MatchThreeV2PVPRank{}, err
			}
		}
		a.GroupData()
	}
	return a._dataIconPath[identifier], nil
}

// ByMaxScore returns the MatchThreeV2PVPRank uniquely identified by MaxScore
//
// Error is only non-nil if the source errors out
func (a *MatchThreeV2PVPRankAccessor) ByMaxScore(identifier float64) (MatchThreeV2PVPRank, error) {
	if a._dataMaxScore == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MatchThreeV2PVPRank{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMaxScore[identifier], nil
}

// ByRankID returns the MatchThreeV2PVPRank uniquely identified by RankID
//
// Error is only non-nil if the source errors out
func (a *MatchThreeV2PVPRankAccessor) ByRankID(identifier float64) (MatchThreeV2PVPRank, error) {
	if a._dataRankID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MatchThreeV2PVPRank{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRankID[identifier], nil
}

// BySmallIconPath returns the MatchThreeV2PVPRank uniquely identified by SmallIconPath
//
// Error is only non-nil if the source errors out
func (a *MatchThreeV2PVPRankAccessor) BySmallIconPath(identifier string) (MatchThreeV2PVPRank, error) {
	if a._dataSmallIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MatchThreeV2PVPRank{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSmallIconPath[identifier], nil
}
