package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type RechargeBenefitConfig struct {
	ActivityModuleID json.Number   `json:"ActivityModuleID"`
	BenefitIDList    []json.Number `json:"BenefitIDList"`
	ID               json.Number   `json:"ID"`
	Type             string        `json:"Type"`
}
type RechargeBenefitConfigAccessor struct {
	_data     []RechargeBenefitConfig
	_dataType map[string]RechargeBenefitConfig
}

// LoadData retrieves the data. Must be called before RechargeBenefitConfig.GroupData
func (a *RechargeBenefitConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RechargeBenefitConfig.json")
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
func (a *RechargeBenefitConfigAccessor) Raw() ([]RechargeBenefitConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RechargeBenefitConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RechargeBenefitConfigAccessor.LoadData to preload everything
func (a *RechargeBenefitConfigAccessor) GroupData() {
	a._dataType = map[string]RechargeBenefitConfig{}
	for _, d := range a._data {
		a._dataType[d.Type] = d
	}
}

// ByType returns the RechargeBenefitConfig uniquely identified by Type
//
// Error is only non-nil if the source errors out
func (a *RechargeBenefitConfigAccessor) ByType(identifier string) (RechargeBenefitConfig, error) {
	if a._dataType == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RechargeBenefitConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataType[identifier], nil
}
