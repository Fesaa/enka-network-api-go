package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type AetherDivideChallengeRank struct {
	ActivityModuleID float64   `json:"ActivityModuleID"`
	ChallengeRank    float64   `json:"ChallengeRank"`
	FunctionUnlockID float64   `json:"FunctionUnlockID"`
	IconPath         string    `json:"IconPath"`
	IsHard           bool      `json:"IsHard"`
	PreRank          float64   `json:"PreRank"`
	TrainerLevel     float64   `json:"TrainerLevel"`
	UnlockText       hash.Hash `json:"UnlockText"`
}
type AetherDivideChallengeRankAccessor struct {
	_data              []AetherDivideChallengeRank
	_dataChallengeRank map[float64]AetherDivideChallengeRank
	_dataIconPath      map[string]AetherDivideChallengeRank
}

// LoadData retrieves the data. Must be called before AetherDivideChallengeRank.GroupData
func (a *AetherDivideChallengeRankAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AetherDivideChallengeRank.json")
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
func (a *AetherDivideChallengeRankAccessor) Raw() ([]AetherDivideChallengeRank, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AetherDivideChallengeRank{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AetherDivideChallengeRankAccessor.LoadData to preload everything
func (a *AetherDivideChallengeRankAccessor) GroupData() {
	a._dataChallengeRank = map[float64]AetherDivideChallengeRank{}
	a._dataIconPath = map[string]AetherDivideChallengeRank{}
	for _, d := range a._data {
		a._dataChallengeRank[d.ChallengeRank] = d
		a._dataIconPath[d.IconPath] = d
	}
}

// ByChallengeRank returns the AetherDivideChallengeRank uniquely identified by ChallengeRank
//
// Error is only non-nil if the source errors out
func (a *AetherDivideChallengeRankAccessor) ByChallengeRank(identifier float64) (AetherDivideChallengeRank, error) {
	if a._dataChallengeRank == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDivideChallengeRank{}, err
			}
		}
		a.GroupData()
	}
	return a._dataChallengeRank[identifier], nil
}

// ByIconPath returns the AetherDivideChallengeRank uniquely identified by IconPath
//
// Error is only non-nil if the source errors out
func (a *AetherDivideChallengeRankAccessor) ByIconPath(identifier string) (AetherDivideChallengeRank, error) {
	if a._dataIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDivideChallengeRank{}, err
			}
		}
		a.GroupData()
	}
	return a._dataIconPath[identifier], nil
}
