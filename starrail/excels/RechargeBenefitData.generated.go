package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type RechargeBenefitData struct {
	BenefitID  float64   `json:"BenefitID"`
	ConsumeNum float64   `json:"ConsumeNum"`
	GiftName   hash.Hash `json:"GiftName"`
	Reward     float64   `json:"Reward"`
}
type RechargeBenefitDataAccessor struct {
	_data          []RechargeBenefitData
	_dataBenefitID map[float64]RechargeBenefitData
	_dataReward    map[float64]RechargeBenefitData
}

// LoadData retrieves the data. Must be called before RechargeBenefitData.GroupData
func (a *RechargeBenefitDataAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RechargeBenefitData.json")
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
func (a *RechargeBenefitDataAccessor) Raw() ([]RechargeBenefitData, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RechargeBenefitData{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RechargeBenefitDataAccessor.LoadData to preload everything
func (a *RechargeBenefitDataAccessor) GroupData() {
	for _, d := range a._data {
		a._dataBenefitID[d.BenefitID] = d
		a._dataReward[d.Reward] = d
	}
}

// ByBenefitID returns the RechargeBenefitData uniquely identified by BenefitID
//
// Error is only non-nil if the source errors out
func (a *RechargeBenefitDataAccessor) ByBenefitID(identifier float64) (RechargeBenefitData, error) {
	if a._dataBenefitID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RechargeBenefitData{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBenefitID[identifier], nil
}

// ByReward returns the RechargeBenefitData uniquely identified by Reward
//
// Error is only non-nil if the source errors out
func (a *RechargeBenefitDataAccessor) ByReward(identifier float64) (RechargeBenefitData, error) {
	if a._dataReward == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RechargeBenefitData{}, err
			}
		}
		a.GroupData()
	}
	return a._dataReward[identifier], nil
}
