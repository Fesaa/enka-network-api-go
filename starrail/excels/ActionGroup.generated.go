package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type ActionGroup struct {
	ActionGroupName          string                 `json:"ActionGroupName"`
	ActionGroupTextmapID     map[string]json.Number `json:"ActionGroupTextmapID"`
	ActionListForAnd         []string               `json:"ActionListForAnd"`
	ActionListForOr          []string               `json:"ActionListForOr"`
	ActionName               string                 `json:"ActionName"`
	FranceKeyMouseImagePath  string                 `json:"FranceKeyMouseImagePath"`
	GermanyKeyMouseImagePath string                 `json:"GermanyKeyMouseImagePath"`
	KeyMouseImagePath        string                 `json:"KeyMouseImagePath"`
	PsImagePath              string                 `json:"PsImagePath"`
	XboxImagePath            string                 `json:"XboxImagePath"`
}
type ActionGroupAccessor struct {
	_data                []ActionGroup
	_dataActionGroupName map[string]ActionGroup
}

// LoadData retrieves the data. Must be called before ActionGroup.GroupData
func (a *ActionGroupAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActionGroup.json")
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
func (a *ActionGroupAccessor) Raw() ([]ActionGroup, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActionGroup{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ActionGroupAccessor.LoadData to preload everything
func (a *ActionGroupAccessor) GroupData() {
	a._dataActionGroupName = map[string]ActionGroup{}
	for _, d := range a._data {
		a._dataActionGroupName[d.ActionGroupName] = d
	}
}

// ByActionGroupName returns the ActionGroup uniquely identified by ActionGroupName
//
// Error is only non-nil if the source errors out
func (a *ActionGroupAccessor) ByActionGroupName(identifier string) (ActionGroup, error) {
	if a._dataActionGroupName == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActionGroup{}, err
			}
		}
		a.GroupData()
	}
	return a._dataActionGroupName[identifier], nil
}
