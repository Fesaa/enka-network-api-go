package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type StrongChallengeBuffConfig struct {
	BuffCost              float64 `json:"BuffCost"`
	StrongChallengeBuffID float64 `json:"StrongChallengeBuffID"`
}
type StrongChallengeBuffConfigAccessor struct {
	_data                      []StrongChallengeBuffConfig
	_dataStrongChallengeBuffID map[float64]StrongChallengeBuffConfig
}

// LoadData retrieves the data. Must be called before StrongChallengeBuffConfig.GroupData
func (a *StrongChallengeBuffConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/StrongChallengeBuffConfig.json")
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
func (a *StrongChallengeBuffConfigAccessor) Raw() ([]StrongChallengeBuffConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []StrongChallengeBuffConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with StrongChallengeBuffConfigAccessor.LoadData to preload everything
func (a *StrongChallengeBuffConfigAccessor) GroupData() {
	a._dataStrongChallengeBuffID = map[float64]StrongChallengeBuffConfig{}
	for _, d := range a._data {
		a._dataStrongChallengeBuffID[d.StrongChallengeBuffID] = d
	}
}

// ByStrongChallengeBuffID returns the StrongChallengeBuffConfig uniquely identified by StrongChallengeBuffID
//
// Error is only non-nil if the source errors out
func (a *StrongChallengeBuffConfigAccessor) ByStrongChallengeBuffID(identifier float64) (StrongChallengeBuffConfig, error) {
	if a._dataStrongChallengeBuffID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return StrongChallengeBuffConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataStrongChallengeBuffID[identifier], nil
}
