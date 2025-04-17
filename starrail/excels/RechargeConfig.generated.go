package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type RechargeConfig struct {
	FirstCharge  json.Number            `json:"FirstCharge"`
	GiftImage    string                 `json:"GiftImage"`
	GiftName     map[string]json.Number `json:"GiftName"`
	GiftType     json.Number            `json:"GiftType"`
	ListOrder    json.Number            `json:"ListOrder"`
	NormalCharge json.Number            `json:"NormalCharge"`
	ProductID    string                 `json:"ProductID"`
	TierID       string                 `json:"TierID"`
}
type RechargeConfigAccessor struct {
	_data          []RechargeConfig
	_dataProductID map[string]RechargeConfig
}

// LoadData retrieves the data. Must be called before RechargeConfig.GroupData
func (a *RechargeConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RechargeConfig.json")
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
func (a *RechargeConfigAccessor) Raw() ([]RechargeConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RechargeConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RechargeConfigAccessor.LoadData to preload everything
func (a *RechargeConfigAccessor) GroupData() {
	a._dataProductID = map[string]RechargeConfig{}
	for _, d := range a._data {
		a._dataProductID[d.ProductID] = d
	}
}

// ByProductID returns the RechargeConfig uniquely identified by ProductID
//
// Error is only non-nil if the source errors out
func (a *RechargeConfigAccessor) ByProductID(identifier string) (RechargeConfig, error) {
	if a._dataProductID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RechargeConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataProductID[identifier], nil
}
