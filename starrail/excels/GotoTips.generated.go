package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type GotoTips struct {
	ID   string                 `json:"ID"`
	Name map[string]json.Number `json:"Name"`
}
type GotoTipsAccessor struct {
	_data   []GotoTips
	_dataID map[string]GotoTips
}

// LoadData retrieves the data. Must be called before GotoTips.GroupData
func (a *GotoTipsAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/GotoTips.json")
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
func (a *GotoTipsAccessor) Raw() ([]GotoTips, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []GotoTips{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with GotoTipsAccessor.LoadData to preload everything
func (a *GotoTipsAccessor) GroupData() {
	a._dataID = map[string]GotoTips{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the GotoTips uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *GotoTipsAccessor) ByID(identifier string) (GotoTips, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return GotoTips{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
