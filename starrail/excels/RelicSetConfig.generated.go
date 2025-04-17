package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type RelicSetConfig struct {
	DisplayItemID        json.Number            `json:"DisplayItemID"`
	DisplayItemIDRarity4 json.Number            `json:"DisplayItemIDRarity4"`
	IsPlanarSuit         bool                   `json:"IsPlanarSuit"`
	Release              bool                   `json:"Release"`
	ReleaseVersion       string                 `json:"ReleaseVersion"`
	SetID                json.Number            `json:"SetID"`
	SetIconFigurePath    string                 `json:"SetIconFigurePath"`
	SetIconPath          string                 `json:"SetIconPath"`
	SetName              map[string]json.Number `json:"SetName"`
	SetSkillList         []json.Number          `json:"SetSkillList"`
}
type RelicSetConfigAccessor struct {
	_data                     []RelicSetConfig
	_dataDisplayItemID        map[json.Number]RelicSetConfig
	_dataDisplayItemIDRarity4 map[json.Number]RelicSetConfig
	_dataSetID                map[json.Number]RelicSetConfig
	_dataSetIconFigurePath    map[string]RelicSetConfig
	_dataSetIconPath          map[string]RelicSetConfig
}

// LoadData retrieves the data. Must be called before RelicSetConfig.GroupData
func (a *RelicSetConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RelicSetConfig.json")
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
func (a *RelicSetConfigAccessor) Raw() ([]RelicSetConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RelicSetConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RelicSetConfigAccessor.LoadData to preload everything
func (a *RelicSetConfigAccessor) GroupData() {
	a._dataDisplayItemID = map[json.Number]RelicSetConfig{}
	a._dataDisplayItemIDRarity4 = map[json.Number]RelicSetConfig{}
	a._dataSetID = map[json.Number]RelicSetConfig{}
	a._dataSetIconFigurePath = map[string]RelicSetConfig{}
	a._dataSetIconPath = map[string]RelicSetConfig{}
	for _, d := range a._data {
		a._dataDisplayItemID[d.DisplayItemID] = d
		a._dataDisplayItemIDRarity4[d.DisplayItemIDRarity4] = d
		a._dataSetID[d.SetID] = d
		a._dataSetIconFigurePath[d.SetIconFigurePath] = d
		a._dataSetIconPath[d.SetIconPath] = d
	}
}

// ByDisplayItemID returns the RelicSetConfig uniquely identified by DisplayItemID
//
// Error is only non-nil if the source errors out
func (a *RelicSetConfigAccessor) ByDisplayItemID(identifier json.Number) (RelicSetConfig, error) {
	if a._dataDisplayItemID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RelicSetConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataDisplayItemID[identifier], nil
}

// ByDisplayItemIDRarity4 returns the RelicSetConfig uniquely identified by DisplayItemIDRarity4
//
// Error is only non-nil if the source errors out
func (a *RelicSetConfigAccessor) ByDisplayItemIDRarity4(identifier json.Number) (RelicSetConfig, error) {
	if a._dataDisplayItemIDRarity4 == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RelicSetConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataDisplayItemIDRarity4[identifier], nil
}

// BySetID returns the RelicSetConfig uniquely identified by SetID
//
// Error is only non-nil if the source errors out
func (a *RelicSetConfigAccessor) BySetID(identifier json.Number) (RelicSetConfig, error) {
	if a._dataSetID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RelicSetConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSetID[identifier], nil
}

// BySetIconFigurePath returns the RelicSetConfig uniquely identified by SetIconFigurePath
//
// Error is only non-nil if the source errors out
func (a *RelicSetConfigAccessor) BySetIconFigurePath(identifier string) (RelicSetConfig, error) {
	if a._dataSetIconFigurePath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RelicSetConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSetIconFigurePath[identifier], nil
}

// BySetIconPath returns the RelicSetConfig uniquely identified by SetIconPath
//
// Error is only non-nil if the source errors out
func (a *RelicSetConfigAccessor) BySetIconPath(identifier string) (RelicSetConfig, error) {
	if a._dataSetIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RelicSetConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSetIconPath[identifier], nil
}
