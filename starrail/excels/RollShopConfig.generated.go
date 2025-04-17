package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type RollShopConfig struct {
	CostItemID       float64                          `json:"CostItemID"`
	CostItemNum      float64                          `json:"CostItemNum"`
	IntroduceID      float64                          `json:"IntroduceID"`
	RollShopID       float64                          `json:"RollShopID"`
	RollShopType     string                           `json:"RollShopType"`
	SecretGroupID    float64                          `json:"SecretGroupID"`
	ShopName         hash.Hash                        `json:"ShopName"`
	SpecialGroupList []RollShopConfigSpecialGroupList `json:"SpecialGroupList"`
	T1GroupID        float64                          `json:"T1GroupID"`
	T2GroupID        float64                          `json:"T2GroupID"`
	T3GroupID        float64                          `json:"T3GroupID"`
	T4GroupID        float64                          `json:"T4GroupID"`
}
type RollShopConfigSpecialGroupList struct {
	CMKJLKBIDLH float64 `json:"CMKJLKBIDLH"`
}
type RollShopConfigAccessor struct {
	_data            []RollShopConfig
	_dataIntroduceID map[float64]RollShopConfig
	_dataRollShopID  map[float64]RollShopConfig
	_dataT2GroupID   map[float64]RollShopConfig
	_dataT3GroupID   map[float64]RollShopConfig
	_dataT4GroupID   map[float64]RollShopConfig
}

// LoadData retrieves the data. Must be called before RollShopConfig.GroupData
func (a *RollShopConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RollShopConfig.json")
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
func (a *RollShopConfigAccessor) Raw() ([]RollShopConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RollShopConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RollShopConfigAccessor.LoadData to preload everything
func (a *RollShopConfigAccessor) GroupData() {
	a._dataIntroduceID = map[float64]RollShopConfig{}
	a._dataRollShopID = map[float64]RollShopConfig{}
	a._dataT2GroupID = map[float64]RollShopConfig{}
	a._dataT3GroupID = map[float64]RollShopConfig{}
	a._dataT4GroupID = map[float64]RollShopConfig{}
	for _, d := range a._data {
		a._dataIntroduceID[d.IntroduceID] = d
		a._dataRollShopID[d.RollShopID] = d
		a._dataT2GroupID[d.T2GroupID] = d
		a._dataT3GroupID[d.T3GroupID] = d
		a._dataT4GroupID[d.T4GroupID] = d
	}
}

// ByIntroduceID returns the RollShopConfig uniquely identified by IntroduceID
//
// Error is only non-nil if the source errors out
func (a *RollShopConfigAccessor) ByIntroduceID(identifier float64) (RollShopConfig, error) {
	if a._dataIntroduceID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RollShopConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataIntroduceID[identifier], nil
}

// ByRollShopID returns the RollShopConfig uniquely identified by RollShopID
//
// Error is only non-nil if the source errors out
func (a *RollShopConfigAccessor) ByRollShopID(identifier float64) (RollShopConfig, error) {
	if a._dataRollShopID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RollShopConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRollShopID[identifier], nil
}

// ByT2GroupID returns the RollShopConfig uniquely identified by T2GroupID
//
// Error is only non-nil if the source errors out
func (a *RollShopConfigAccessor) ByT2GroupID(identifier float64) (RollShopConfig, error) {
	if a._dataT2GroupID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RollShopConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataT2GroupID[identifier], nil
}

// ByT3GroupID returns the RollShopConfig uniquely identified by T3GroupID
//
// Error is only non-nil if the source errors out
func (a *RollShopConfigAccessor) ByT3GroupID(identifier float64) (RollShopConfig, error) {
	if a._dataT3GroupID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RollShopConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataT3GroupID[identifier], nil
}

// ByT4GroupID returns the RollShopConfig uniquely identified by T4GroupID
//
// Error is only non-nil if the source errors out
func (a *RollShopConfigAccessor) ByT4GroupID(identifier float64) (RollShopConfig, error) {
	if a._dataT4GroupID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RollShopConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataT4GroupID[identifier], nil
}
