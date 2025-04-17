package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type ChallengeTargetConfig struct {
	ChallengeTargetName   hash.Hash `json:"ChallengeTargetName"`
	ChallengeTargetParam1 float64   `json:"ChallengeTargetParam1"`
	ChallengeTargetType   string    `json:"ChallengeTargetType"`
	ID                    float64   `json:"ID"`
	RewardID              float64   `json:"RewardID"`
}
type ChallengeTargetConfigAccessor struct {
	_data   []ChallengeTargetConfig
	_dataID map[float64]ChallengeTargetConfig
}

// LoadData retrieves the data. Must be called before ChallengeTargetConfig.GroupData
func (a *ChallengeTargetConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ChallengeTargetConfig.json")
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
func (a *ChallengeTargetConfigAccessor) Raw() ([]ChallengeTargetConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ChallengeTargetConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ChallengeTargetConfigAccessor.LoadData to preload everything
func (a *ChallengeTargetConfigAccessor) GroupData() {
	a._dataID = map[float64]ChallengeTargetConfig{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the ChallengeTargetConfig uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *ChallengeTargetConfigAccessor) ByID(identifier float64) (ChallengeTargetConfig, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ChallengeTargetConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
