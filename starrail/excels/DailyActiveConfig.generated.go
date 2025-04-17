package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type DailyActiveConfig struct {
	DailyActivePoint  float64 `json:"DailyActivePoint"`
	DailyActiveReward float64 `json:"DailyActiveReward"`
	Level             float64 `json:"Level"`
	WorldLevel        float64 `json:"WorldLevel"`
}
type DailyActiveConfigAccessor struct {
	_data                  []DailyActiveConfig
	_dataDailyActiveReward map[float64]DailyActiveConfig
}

// LoadData retrieves the data. Must be called before DailyActiveConfig.GroupData
func (a *DailyActiveConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/DailyActiveConfig.json")
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
func (a *DailyActiveConfigAccessor) Raw() ([]DailyActiveConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []DailyActiveConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with DailyActiveConfigAccessor.LoadData to preload everything
func (a *DailyActiveConfigAccessor) GroupData() {
	a._dataDailyActiveReward = map[float64]DailyActiveConfig{}
	for _, d := range a._data {
		a._dataDailyActiveReward[d.DailyActiveReward] = d
	}
}

// ByDailyActiveReward returns the DailyActiveConfig uniquely identified by DailyActiveReward
//
// Error is only non-nil if the source errors out
func (a *DailyActiveConfigAccessor) ByDailyActiveReward(identifier float64) (DailyActiveConfig, error) {
	if a._dataDailyActiveReward == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return DailyActiveConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataDailyActiveReward[identifier], nil
}
