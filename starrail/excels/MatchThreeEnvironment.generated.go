package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type MatchThreeEnvironment struct {
	Desc          map[string]json.Number `json:"Desc"`
	EnvironmentID json.Number            `json:"EnvironmentID"`
	IconPath      string                 `json:"IconPath"`
	ImagePath     string                 `json:"ImagePath"`
	Name          map[string]json.Number `json:"Name"`
	ParamList     []string               `json:"ParamList"`
}
type MatchThreeEnvironmentAccessor struct {
	_data          []MatchThreeEnvironment
	_dataIconPath  map[string]MatchThreeEnvironment
	_dataImagePath map[string]MatchThreeEnvironment
}

// LoadData retrieves the data. Must be called before MatchThreeEnvironment.GroupData
func (a *MatchThreeEnvironmentAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MatchThreeEnvironment.json")
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
func (a *MatchThreeEnvironmentAccessor) Raw() ([]MatchThreeEnvironment, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MatchThreeEnvironment{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MatchThreeEnvironmentAccessor.LoadData to preload everything
func (a *MatchThreeEnvironmentAccessor) GroupData() {
	a._dataIconPath = map[string]MatchThreeEnvironment{}
	a._dataImagePath = map[string]MatchThreeEnvironment{}
	for _, d := range a._data {
		a._dataIconPath[d.IconPath] = d
		a._dataImagePath[d.ImagePath] = d
	}
}

// ByIconPath returns the MatchThreeEnvironment uniquely identified by IconPath
//
// Error is only non-nil if the source errors out
func (a *MatchThreeEnvironmentAccessor) ByIconPath(identifier string) (MatchThreeEnvironment, error) {
	if a._dataIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MatchThreeEnvironment{}, err
			}
		}
		a.GroupData()
	}
	return a._dataIconPath[identifier], nil
}

// ByImagePath returns the MatchThreeEnvironment uniquely identified by ImagePath
//
// Error is only non-nil if the source errors out
func (a *MatchThreeEnvironmentAccessor) ByImagePath(identifier string) (MatchThreeEnvironment, error) {
	if a._dataImagePath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MatchThreeEnvironment{}, err
			}
		}
		a.GroupData()
	}
	return a._dataImagePath[identifier], nil
}
