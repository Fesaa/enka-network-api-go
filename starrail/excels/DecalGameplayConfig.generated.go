package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type DecalGameplayConfig struct {
	DecalID        json.Number `json:"DecalID"`
	IconPath       string      `json:"IconPath"`
	TextureMapPath string      `json:"TextureMapPath"`
}
type DecalGameplayConfigAccessor struct {
	_data               []DecalGameplayConfig
	_dataIconPath       map[string]DecalGameplayConfig
	_dataTextureMapPath map[string]DecalGameplayConfig
}

// LoadData retrieves the data. Must be called before DecalGameplayConfig.GroupData
func (a *DecalGameplayConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/DecalGameplayConfig.json")
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
func (a *DecalGameplayConfigAccessor) Raw() ([]DecalGameplayConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []DecalGameplayConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with DecalGameplayConfigAccessor.LoadData to preload everything
func (a *DecalGameplayConfigAccessor) GroupData() {
	a._dataIconPath = map[string]DecalGameplayConfig{}
	a._dataTextureMapPath = map[string]DecalGameplayConfig{}
	for _, d := range a._data {
		a._dataIconPath[d.IconPath] = d
		a._dataTextureMapPath[d.TextureMapPath] = d
	}
}

// ByIconPath returns the DecalGameplayConfig uniquely identified by IconPath
//
// Error is only non-nil if the source errors out
func (a *DecalGameplayConfigAccessor) ByIconPath(identifier string) (DecalGameplayConfig, error) {
	if a._dataIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return DecalGameplayConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataIconPath[identifier], nil
}

// ByTextureMapPath returns the DecalGameplayConfig uniquely identified by TextureMapPath
//
// Error is only non-nil if the source errors out
func (a *DecalGameplayConfigAccessor) ByTextureMapPath(identifier string) (DecalGameplayConfig, error) {
	if a._dataTextureMapPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return DecalGameplayConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataTextureMapPath[identifier], nil
}
