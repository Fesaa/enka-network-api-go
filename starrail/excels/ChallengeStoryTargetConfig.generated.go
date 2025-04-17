package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type ChallengeStoryTargetConfig struct {
	ChallengeTargetName   hash.Hash `json:"ChallengeTargetName"`
	ChallengeTargetParam1 float64   `json:"ChallengeTargetParam1"`
	ChallengeTargetType   string    `json:"ChallengeTargetType"`
	ID                    float64   `json:"ID"`
}
type ChallengeStoryTargetConfigAccessor struct {
	_data                      []ChallengeStoryTargetConfig
	_dataChallengeTargetParam1 map[float64]ChallengeStoryTargetConfig
	_dataID                    map[float64]ChallengeStoryTargetConfig
}

// LoadData retrieves the data. Must be called before ChallengeStoryTargetConfig.GroupData
func (a *ChallengeStoryTargetConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ChallengeStoryTargetConfig.json")
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
func (a *ChallengeStoryTargetConfigAccessor) Raw() ([]ChallengeStoryTargetConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ChallengeStoryTargetConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ChallengeStoryTargetConfigAccessor.LoadData to preload everything
func (a *ChallengeStoryTargetConfigAccessor) GroupData() {
	for _, d := range a._data {
		a._dataChallengeTargetParam1[d.ChallengeTargetParam1] = d
		a._dataID[d.ID] = d
	}
}

// ByChallengeTargetParam1 returns the ChallengeStoryTargetConfig uniquely identified by ChallengeTargetParam1
//
// Error is only non-nil if the source errors out
func (a *ChallengeStoryTargetConfigAccessor) ByChallengeTargetParam1(identifier float64) (ChallengeStoryTargetConfig, error) {
	if a._dataChallengeTargetParam1 == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ChallengeStoryTargetConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataChallengeTargetParam1[identifier], nil
}

// ByID returns the ChallengeStoryTargetConfig uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *ChallengeStoryTargetConfigAccessor) ByID(identifier float64) (ChallengeStoryTargetConfig, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ChallengeStoryTargetConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
