package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type RechargeGiftConfig struct {
	Discount   float64   `json:"Discount"`
	GiftIDList []float64 `json:"GiftIDList"`
	GiftType   float64   `json:"GiftType"`
}
type RechargeGiftConfigAccessor struct {
	_data         []RechargeGiftConfig
	_dataGiftType map[float64]RechargeGiftConfig
	_dataDiscount map[float64]RechargeGiftConfig
}

// LoadData retrieves the data. Must be called before RechargeGiftConfig.GroupData
func (a *RechargeGiftConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RechargeGiftConfig.json")
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
func (a *RechargeGiftConfigAccessor) Raw() ([]RechargeGiftConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RechargeGiftConfig{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RechargeGiftConfigAccessor.LoadData to preload everything
func (a *RechargeGiftConfigAccessor) GroupData() {
	for _, d := range a._data {
		a._dataGiftType[d.GiftType] = d
		a._dataDiscount[d.Discount] = d
	}
}

// ByGiftType returns the RechargeGiftConfig uniquely identified by GiftType
//
// Error is only non-nil if the source errors out
func (a *RechargeGiftConfigAccessor) ByGiftType(identifier float64) (RechargeGiftConfig, error) {
	if a._dataGiftType == nil {
		err := a.LoadData()
		if err != nil {
			return RechargeGiftConfig{}, err
		}
		a.GroupData()
	}
	return a._dataGiftType[identifier], nil
}

// ByDiscount returns the RechargeGiftConfig uniquely identified by Discount
//
// Error is only non-nil if the source errors out
func (a *RechargeGiftConfigAccessor) ByDiscount(identifier float64) (RechargeGiftConfig, error) {
	if a._dataDiscount == nil {
		err := a.LoadData()
		if err != nil {
			return RechargeGiftConfig{}, err
		}
		a.GroupData()
	}
	return a._dataDiscount[identifier], nil
}
