package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type AlleyMapReward struct {
	LayerID  string  `json:"LayerID"`
	MapScore float64 `json:"MapScore"`
	RewardID float64 `json:"RewardID"`
	ScoreID  float64 `json:"ScoreID"`
}
type AlleyMapRewardAccessor struct {
	_data        []AlleyMapReward
	_dataScoreID map[float64]AlleyMapReward
}

// LoadData retrieves the data. Must be called before AlleyMapReward.GroupData
func (a *AlleyMapRewardAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AlleyMapReward.json")
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
func (a *AlleyMapRewardAccessor) Raw() ([]AlleyMapReward, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AlleyMapReward{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AlleyMapRewardAccessor.LoadData to preload everything
func (a *AlleyMapRewardAccessor) GroupData() {
	a._dataScoreID = map[float64]AlleyMapReward{}
	for _, d := range a._data {
		a._dataScoreID[d.ScoreID] = d
	}
}

// ByScoreID returns the AlleyMapReward uniquely identified by ScoreID
//
// Error is only non-nil if the source errors out
func (a *AlleyMapRewardAccessor) ByScoreID(identifier float64) (AlleyMapReward, error) {
	if a._dataScoreID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AlleyMapReward{}, err
			}
		}
		a.GroupData()
	}
	return a._dataScoreID[identifier], nil
}
