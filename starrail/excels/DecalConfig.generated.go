package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type DecalConfig struct {
	BgPath         string                 `json:"BgPath"`
	Comment        string                 `json:"Comment"`
	DecalID        json.Number            `json:"DecalID"`
	Desc           map[string]json.Number `json:"Desc"`
	FigurePath     string                 `json:"FigurePath"`
	IconPath       string                 `json:"IconPath"`
	Name           map[string]json.Number `json:"Name"`
	TextureMapPath string                 `json:"TextureMapPath"`
	UnlockMission  json.Number            `json:"UnlockMission"`
}
type DecalConfigAccessor struct {
	_data               []DecalConfig
	_dataComment        map[string]DecalConfig
	_dataDecalID        map[json.Number]DecalConfig
	_dataTextureMapPath map[string]DecalConfig
}

// LoadData retrieves the data. Must be called before DecalConfig.GroupData
func (a *DecalConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/DecalConfig.json")
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
func (a *DecalConfigAccessor) Raw() ([]DecalConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []DecalConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with DecalConfigAccessor.LoadData to preload everything
func (a *DecalConfigAccessor) GroupData() {
	a._dataComment = map[string]DecalConfig{}
	a._dataDecalID = map[json.Number]DecalConfig{}
	a._dataTextureMapPath = map[string]DecalConfig{}
	for _, d := range a._data {
		a._dataComment[d.Comment] = d
		a._dataDecalID[d.DecalID] = d
		a._dataTextureMapPath[d.TextureMapPath] = d
	}
}

// ByComment returns the DecalConfig uniquely identified by Comment
//
// Error is only non-nil if the source errors out
func (a *DecalConfigAccessor) ByComment(identifier string) (DecalConfig, error) {
	if a._dataComment == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return DecalConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataComment[identifier], nil
}

// ByDecalID returns the DecalConfig uniquely identified by DecalID
//
// Error is only non-nil if the source errors out
func (a *DecalConfigAccessor) ByDecalID(identifier json.Number) (DecalConfig, error) {
	if a._dataDecalID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return DecalConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataDecalID[identifier], nil
}

// ByTextureMapPath returns the DecalConfig uniquely identified by TextureMapPath
//
// Error is only non-nil if the source errors out
func (a *DecalConfigAccessor) ByTextureMapPath(identifier string) (DecalConfig, error) {
	if a._dataTextureMapPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return DecalConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataTextureMapPath[identifier], nil
}
