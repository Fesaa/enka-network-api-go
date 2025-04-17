package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type AvatarStatusConfig struct {
	CanDispel              bool                   `json:"CanDispel"`
	ModifierName           string                 `json:"ModifierName"`
	ReadParamList          []string               `json:"ReadParamList"`
	StatusDesc             map[string]json.Number `json:"StatusDesc"`
	StatusEffect           map[string]json.Number `json:"StatusEffect"`
	StatusID               json.Number            `json:"StatusID"`
	StatusIconPath         string                 `json:"StatusIconPath"`
	StatusIconPathHighSize string                 `json:"StatusIconPathHighSize"`
	StatusName             map[string]json.Number `json:"StatusName"`
	StatusType             string                 `json:"StatusType"`
	TagList                []interface{}          `json:"TagList"`
}
type AvatarStatusConfigAccessor struct {
	_data             []AvatarStatusConfig
	_dataModifierName map[string]AvatarStatusConfig
	_dataStatusID     map[json.Number]AvatarStatusConfig
}

// LoadData retrieves the data. Must be called before AvatarStatusConfig.GroupData
func (a *AvatarStatusConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AvatarStatusConfig.json")
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
func (a *AvatarStatusConfigAccessor) Raw() ([]AvatarStatusConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AvatarStatusConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AvatarStatusConfigAccessor.LoadData to preload everything
func (a *AvatarStatusConfigAccessor) GroupData() {
	a._dataModifierName = map[string]AvatarStatusConfig{}
	a._dataStatusID = map[json.Number]AvatarStatusConfig{}
	for _, d := range a._data {
		a._dataModifierName[d.ModifierName] = d
		a._dataStatusID[d.StatusID] = d
	}
}

// ByModifierName returns the AvatarStatusConfig uniquely identified by ModifierName
//
// Error is only non-nil if the source errors out
func (a *AvatarStatusConfigAccessor) ByModifierName(identifier string) (AvatarStatusConfig, error) {
	if a._dataModifierName == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarStatusConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataModifierName[identifier], nil
}

// ByStatusID returns the AvatarStatusConfig uniquely identified by StatusID
//
// Error is only non-nil if the source errors out
func (a *AvatarStatusConfigAccessor) ByStatusID(identifier json.Number) (AvatarStatusConfig, error) {
	if a._dataStatusID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarStatusConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataStatusID[identifier], nil
}
