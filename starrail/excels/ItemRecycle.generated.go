package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type ItemRecycle struct {
	ItemID       json.Number            `json:"ItemID"`
	RecycleTime  string                 `json:"RecycleTime"`
	ShowTimeType bool                   `json:"ShowTimeType"`
	ShowType     string                 `json:"ShowType"`
	Tips         map[string]json.Number `json:"Tips"`
}
type ItemRecycleAccessor struct {
	_data       []ItemRecycle
	_dataItemID map[json.Number]ItemRecycle
}

// LoadData retrieves the data. Must be called before ItemRecycle.GroupData
func (a *ItemRecycleAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ItemRecycle.json")
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
func (a *ItemRecycleAccessor) Raw() ([]ItemRecycle, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ItemRecycle{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ItemRecycleAccessor.LoadData to preload everything
func (a *ItemRecycleAccessor) GroupData() {
	a._dataItemID = map[json.Number]ItemRecycle{}
	for _, d := range a._data {
		a._dataItemID[d.ItemID] = d
	}
}

// ByItemID returns the ItemRecycle uniquely identified by ItemID
//
// Error is only non-nil if the source errors out
func (a *ItemRecycleAccessor) ByItemID(identifier json.Number) (ItemRecycle, error) {
	if a._dataItemID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ItemRecycle{}, err
			}
		}
		a.GroupData()
	}
	return a._dataItemID[identifier], nil
}
