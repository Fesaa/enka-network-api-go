package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type EvolveBuildCardConfig struct {
	CardSelectablePeriod []interface{} `json:"CardSelectablePeriod"`
	ID                   float64       `json:"ID"`
	InfluenceScope       string        `json:"InfluenceScope"`
	ItemIcon             string        `json:"ItemIcon"`
	ItemMiniIcon         string        `json:"ItemMiniIcon"`
	LvID                 float64       `json:"LvID"`
	ParamList            []hash.Value  `json:"ParamList"`
	Season               string        `json:"Season"`
	Type                 string        `json:"Type"`
}
type EvolveBuildCardConfigAccessor struct {
	_data             []EvolveBuildCardConfig
	_dataID           map[float64]EvolveBuildCardConfig
	_dataItemIcon     map[string]EvolveBuildCardConfig
	_dataItemMiniIcon map[string]EvolveBuildCardConfig
	_dataLvID         map[float64]EvolveBuildCardConfig
}

// LoadData retrieves the data. Must be called before EvolveBuildCardConfig.GroupData
func (a *EvolveBuildCardConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/EvolveBuildCardConfig.json")
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
func (a *EvolveBuildCardConfigAccessor) Raw() ([]EvolveBuildCardConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []EvolveBuildCardConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with EvolveBuildCardConfigAccessor.LoadData to preload everything
func (a *EvolveBuildCardConfigAccessor) GroupData() {
	a._dataID = map[float64]EvolveBuildCardConfig{}
	a._dataItemIcon = map[string]EvolveBuildCardConfig{}
	a._dataItemMiniIcon = map[string]EvolveBuildCardConfig{}
	a._dataLvID = map[float64]EvolveBuildCardConfig{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
		a._dataItemIcon[d.ItemIcon] = d
		a._dataItemMiniIcon[d.ItemMiniIcon] = d
		a._dataLvID[d.LvID] = d
	}
}

// ByID returns the EvolveBuildCardConfig uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *EvolveBuildCardConfigAccessor) ByID(identifier float64) (EvolveBuildCardConfig, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return EvolveBuildCardConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByItemIcon returns the EvolveBuildCardConfig uniquely identified by ItemIcon
//
// Error is only non-nil if the source errors out
func (a *EvolveBuildCardConfigAccessor) ByItemIcon(identifier string) (EvolveBuildCardConfig, error) {
	if a._dataItemIcon == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return EvolveBuildCardConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataItemIcon[identifier], nil
}

// ByItemMiniIcon returns the EvolveBuildCardConfig uniquely identified by ItemMiniIcon
//
// Error is only non-nil if the source errors out
func (a *EvolveBuildCardConfigAccessor) ByItemMiniIcon(identifier string) (EvolveBuildCardConfig, error) {
	if a._dataItemMiniIcon == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return EvolveBuildCardConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataItemMiniIcon[identifier], nil
}

// ByLvID returns the EvolveBuildCardConfig uniquely identified by LvID
//
// Error is only non-nil if the source errors out
func (a *EvolveBuildCardConfigAccessor) ByLvID(identifier float64) (EvolveBuildCardConfig, error) {
	if a._dataLvID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return EvolveBuildCardConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataLvID[identifier], nil
}
