package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ChallengeSkipConfig struct {
	ChallengeGroupType string  `json:"ChallengeGroupType"`
	MaxSkipLevel       float64 `json:"MaxSkipLevel"`
}
type ChallengeSkipConfigAccessor struct {
	_data                   []ChallengeSkipConfig
	_dataChallengeGroupType map[string]ChallengeSkipConfig
}

// LoadData retrieves the data. Must be called before ChallengeSkipConfig.GroupData
func (a *ChallengeSkipConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ChallengeSkipConfig.json")
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
func (a *ChallengeSkipConfigAccessor) Raw() ([]ChallengeSkipConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ChallengeSkipConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ChallengeSkipConfigAccessor.LoadData to preload everything
func (a *ChallengeSkipConfigAccessor) GroupData() {
	a._dataChallengeGroupType = map[string]ChallengeSkipConfig{}
	for _, d := range a._data {
		a._dataChallengeGroupType[d.ChallengeGroupType] = d
	}
}

// ByChallengeGroupType returns the ChallengeSkipConfig uniquely identified by ChallengeGroupType
//
// Error is only non-nil if the source errors out
func (a *ChallengeSkipConfigAccessor) ByChallengeGroupType(identifier string) (ChallengeSkipConfig, error) {
	if a._dataChallengeGroupType == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ChallengeSkipConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataChallengeGroupType[identifier], nil
}
