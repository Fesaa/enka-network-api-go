package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type OfferingTypeConfig struct {
	ActivityModuleID float64 `json:"ActivityModuleID"`
	ID               float64 `json:"ID"`
	IsAutoOffer      bool    `json:"IsAutoOffer"`
	ItemID           float64 `json:"ItemID"`
	LongTailLimit    float64 `json:"LongTailLimit"`
	MaxLevel         float64 `json:"MaxLevel"`
	UnlockID         float64 `json:"UnlockID"`
}
type OfferingTypeConfigAccessor struct {
	_data       []OfferingTypeConfig
	_dataID     map[float64]OfferingTypeConfig
	_dataItemID map[float64]OfferingTypeConfig
}

// LoadData retrieves the data. Must be called before OfferingTypeConfig.GroupData
func (a *OfferingTypeConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/OfferingTypeConfig.json")
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
func (a *OfferingTypeConfigAccessor) Raw() ([]OfferingTypeConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []OfferingTypeConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with OfferingTypeConfigAccessor.LoadData to preload everything
func (a *OfferingTypeConfigAccessor) GroupData() {
	for _, d := range a._data {
		a._dataID[d.ID] = d
		a._dataItemID[d.ItemID] = d
	}
}

// ByID returns the OfferingTypeConfig uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *OfferingTypeConfigAccessor) ByID(identifier float64) (OfferingTypeConfig, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return OfferingTypeConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByItemID returns the OfferingTypeConfig uniquely identified by ItemID
//
// Error is only non-nil if the source errors out
func (a *OfferingTypeConfigAccessor) ByItemID(identifier float64) (OfferingTypeConfig, error) {
	if a._dataItemID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return OfferingTypeConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataItemID[identifier], nil
}
