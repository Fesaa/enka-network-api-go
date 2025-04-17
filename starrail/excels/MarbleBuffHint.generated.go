package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type MarbleBuffHint struct {
	HintText map[string]json.Number `json:"HintText"`
	ID       json.Number            `json:"ID"`
}
type MarbleBuffHintAccessor struct {
	_data   []MarbleBuffHint
	_dataID map[json.Number]MarbleBuffHint
}

// LoadData retrieves the data. Must be called before MarbleBuffHint.GroupData
func (a *MarbleBuffHintAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MarbleBuffHint.json")
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
func (a *MarbleBuffHintAccessor) Raw() ([]MarbleBuffHint, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MarbleBuffHint{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MarbleBuffHintAccessor.LoadData to preload everything
func (a *MarbleBuffHintAccessor) GroupData() {
	a._dataID = map[json.Number]MarbleBuffHint{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the MarbleBuffHint uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *MarbleBuffHintAccessor) ByID(identifier json.Number) (MarbleBuffHint, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MarbleBuffHint{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
