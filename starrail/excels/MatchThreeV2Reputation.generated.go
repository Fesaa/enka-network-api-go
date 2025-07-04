package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type MatchThreeV2Reputation struct {
	BgPath         string    `json:"BgPath"`
	ChallengerList []float64 `json:"ChallengerList"`
	ImagePath      string    `json:"ImagePath"`
	LevelUpDesc    hash.Hash `json:"LevelUpDesc"`
	LevelUpReward  float64   `json:"LevelUpReward"`
	Reputation     float64   `json:"Reputation"`
	TabName        hash.Hash `json:"TabName"`
	Title          hash.Hash `json:"Title"`
}
type MatchThreeV2ReputationAccessor struct {
	_data           []MatchThreeV2Reputation
	_dataBgPath     map[string]MatchThreeV2Reputation
	_dataImagePath  map[string]MatchThreeV2Reputation
	_dataReputation map[float64]MatchThreeV2Reputation
}

// LoadData retrieves the data. Must be called before MatchThreeV2Reputation.GroupData
func (a *MatchThreeV2ReputationAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MatchThreeV2Reputation.json")
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
func (a *MatchThreeV2ReputationAccessor) Raw() ([]MatchThreeV2Reputation, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MatchThreeV2Reputation{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MatchThreeV2ReputationAccessor.LoadData to preload everything
func (a *MatchThreeV2ReputationAccessor) GroupData() {
	a._dataBgPath = map[string]MatchThreeV2Reputation{}
	a._dataImagePath = map[string]MatchThreeV2Reputation{}
	a._dataReputation = map[float64]MatchThreeV2Reputation{}
	for _, d := range a._data {
		a._dataBgPath[d.BgPath] = d
		a._dataImagePath[d.ImagePath] = d
		a._dataReputation[d.Reputation] = d
	}
}

// ByBgPath returns the MatchThreeV2Reputation uniquely identified by BgPath
//
// Error is only non-nil if the source errors out
func (a *MatchThreeV2ReputationAccessor) ByBgPath(identifier string) (MatchThreeV2Reputation, error) {
	if a._dataBgPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MatchThreeV2Reputation{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBgPath[identifier], nil
}

// ByImagePath returns the MatchThreeV2Reputation uniquely identified by ImagePath
//
// Error is only non-nil if the source errors out
func (a *MatchThreeV2ReputationAccessor) ByImagePath(identifier string) (MatchThreeV2Reputation, error) {
	if a._dataImagePath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MatchThreeV2Reputation{}, err
			}
		}
		a.GroupData()
	}
	return a._dataImagePath[identifier], nil
}

// ByReputation returns the MatchThreeV2Reputation uniquely identified by Reputation
//
// Error is only non-nil if the source errors out
func (a *MatchThreeV2ReputationAccessor) ByReputation(identifier float64) (MatchThreeV2Reputation, error) {
	if a._dataReputation == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MatchThreeV2Reputation{}, err
			}
		}
		a.GroupData()
	}
	return a._dataReputation[identifier], nil
}
