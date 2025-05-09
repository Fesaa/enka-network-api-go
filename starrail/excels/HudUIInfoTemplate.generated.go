package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type HudUIInfoTemplate struct {
	ActionOperationSetID float64  `json:"ActionOperationSetID"`
	HideHudUINodeList    []string `json:"HideHudUINodeList"`
	ID                   float64  `json:"ID"`
	LockGotoTypeList     []string `json:"LockGotoTypeList"`
	LockInputActionName  []string `json:"LockInputActionName"`
}
type HudUIInfoTemplateAccessor struct {
	_data   []HudUIInfoTemplate
	_dataID map[float64]HudUIInfoTemplate
}

// LoadData retrieves the data. Must be called before HudUIInfoTemplate.GroupData
func (a *HudUIInfoTemplateAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/HudUIInfoTemplate.json")
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
func (a *HudUIInfoTemplateAccessor) Raw() ([]HudUIInfoTemplate, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []HudUIInfoTemplate{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with HudUIInfoTemplateAccessor.LoadData to preload everything
func (a *HudUIInfoTemplateAccessor) GroupData() {
	a._dataID = map[float64]HudUIInfoTemplate{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the HudUIInfoTemplate uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *HudUIInfoTemplateAccessor) ByID(identifier float64) (HudUIInfoTemplate, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return HudUIInfoTemplate{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
