package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type EvolveBuildTagConfig struct {
	ExtraEffectID float64   `json:"ExtraEffectID"`
	ID            float64   `json:"ID"`
	IconPath      string    `json:"IconPath"`
	Name          hash.Hash `json:"Name"`
	Season        string    `json:"Season"`
	ShopSkillID   float64   `json:"ShopSkillID"`
}
type EvolveBuildTagConfigAccessor struct {
	_data              []EvolveBuildTagConfig
	_dataExtraEffectID map[float64]EvolveBuildTagConfig
	_dataID            map[float64]EvolveBuildTagConfig
	_dataIconPath      map[string]EvolveBuildTagConfig
	_dataShopSkillID   map[float64]EvolveBuildTagConfig
}

// LoadData retrieves the data. Must be called before EvolveBuildTagConfig.GroupData
func (a *EvolveBuildTagConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/EvolveBuildTagConfig.json")
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
func (a *EvolveBuildTagConfigAccessor) Raw() ([]EvolveBuildTagConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []EvolveBuildTagConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with EvolveBuildTagConfigAccessor.LoadData to preload everything
func (a *EvolveBuildTagConfigAccessor) GroupData() {
	for _, d := range a._data {
		a._dataExtraEffectID[d.ExtraEffectID] = d
		a._dataID[d.ID] = d
		a._dataIconPath[d.IconPath] = d
		a._dataShopSkillID[d.ShopSkillID] = d
	}
}

// ByExtraEffectID returns the EvolveBuildTagConfig uniquely identified by ExtraEffectID
//
// Error is only non-nil if the source errors out
func (a *EvolveBuildTagConfigAccessor) ByExtraEffectID(identifier float64) (EvolveBuildTagConfig, error) {
	if a._dataExtraEffectID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return EvolveBuildTagConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataExtraEffectID[identifier], nil
}

// ByID returns the EvolveBuildTagConfig uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *EvolveBuildTagConfigAccessor) ByID(identifier float64) (EvolveBuildTagConfig, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return EvolveBuildTagConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByIconPath returns the EvolveBuildTagConfig uniquely identified by IconPath
//
// Error is only non-nil if the source errors out
func (a *EvolveBuildTagConfigAccessor) ByIconPath(identifier string) (EvolveBuildTagConfig, error) {
	if a._dataIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return EvolveBuildTagConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataIconPath[identifier], nil
}

// ByShopSkillID returns the EvolveBuildTagConfig uniquely identified by ShopSkillID
//
// Error is only non-nil if the source errors out
func (a *EvolveBuildTagConfigAccessor) ByShopSkillID(identifier float64) (EvolveBuildTagConfig, error) {
	if a._dataShopSkillID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return EvolveBuildTagConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataShopSkillID[identifier], nil
}
