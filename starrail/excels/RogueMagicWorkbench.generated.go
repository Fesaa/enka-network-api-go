package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type RogueMagicWorkbench struct {
	FuncList    []float64 `json:"FuncList"`
	WorkbenchID float64   `json:"WorkbenchID"`
}
type RogueMagicWorkbenchAccessor struct {
	_data            []RogueMagicWorkbench
	_dataWorkbenchID map[float64]RogueMagicWorkbench
}

// LoadData retrieves the data. Must be called before RogueMagicWorkbench.GroupData
func (a *RogueMagicWorkbenchAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueMagicWorkbench.json")
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
func (a *RogueMagicWorkbenchAccessor) Raw() ([]RogueMagicWorkbench, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueMagicWorkbench{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueMagicWorkbenchAccessor.LoadData to preload everything
func (a *RogueMagicWorkbenchAccessor) GroupData() {
	a._dataWorkbenchID = map[float64]RogueMagicWorkbench{}
	for _, d := range a._data {
		a._dataWorkbenchID[d.WorkbenchID] = d
	}
}

// ByWorkbenchID returns the RogueMagicWorkbench uniquely identified by WorkbenchID
//
// Error is only non-nil if the source errors out
func (a *RogueMagicWorkbenchAccessor) ByWorkbenchID(identifier float64) (RogueMagicWorkbench, error) {
	if a._dataWorkbenchID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueMagicWorkbench{}, err
			}
		}
		a.GroupData()
	}
	return a._dataWorkbenchID[identifier], nil
}
