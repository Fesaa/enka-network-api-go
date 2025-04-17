package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type RogueDLCConstValueClient struct {
	ConstValueName string                        `json:"ConstValueName"`
	Value          RogueDLCConstValueClientValue `json:"Value"`
}
type RogueDLCConstValueClientValue struct {
	ArrayValue []map[string]json.Number `json:"ArrayValue"`
	IntValue   json.Number              `json:"IntValue"`
}
type RogueDLCConstValueClientAccessor struct {
	_data               []RogueDLCConstValueClient
	_dataConstValueName map[string]RogueDLCConstValueClient
}

// LoadData retrieves the data. Must be called before RogueDLCConstValueClient.GroupData
func (a *RogueDLCConstValueClientAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueDLCConstValueClient.json")
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
func (a *RogueDLCConstValueClientAccessor) Raw() ([]RogueDLCConstValueClient, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueDLCConstValueClient{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueDLCConstValueClientAccessor.LoadData to preload everything
func (a *RogueDLCConstValueClientAccessor) GroupData() {
	a._dataConstValueName = map[string]RogueDLCConstValueClient{}
	for _, d := range a._data {
		a._dataConstValueName[d.ConstValueName] = d
	}
}

// ByConstValueName returns the RogueDLCConstValueClient uniquely identified by ConstValueName
//
// Error is only non-nil if the source errors out
func (a *RogueDLCConstValueClientAccessor) ByConstValueName(identifier string) (RogueDLCConstValueClient, error) {
	if a._dataConstValueName == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueDLCConstValueClient{}, err
			}
		}
		a.GroupData()
	}
	return a._dataConstValueName[identifier], nil
}
