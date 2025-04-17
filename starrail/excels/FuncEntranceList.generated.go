package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type FuncEntranceList struct {
	BottomFuncEntranceIDList  []json.Number `json:"BottomFuncEntranceIDList"`
	FuncEntranceIDList        []json.Number `json:"FuncEntranceIDList"`
	HudFuncEntranceIDList     []json.Number `json:"HudFuncEntranceIDList"`
	ID                        json.Number   `json:"ID"`
	LeftHudFuncEntranceIDList []json.Number `json:"LeftHudFuncEntranceIDList"`
	UnlockGotoTypeList        []json.Number `json:"UnlockGotoTypeList"`
	WheelSupport              bool          `json:"WheelSupport"`
}
type FuncEntranceListAccessor struct {
	_data   []FuncEntranceList
	_dataID map[json.Number]FuncEntranceList
}

// LoadData retrieves the data. Must be called before FuncEntranceList.GroupData
func (a *FuncEntranceListAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/FuncEntranceList.json")
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
func (a *FuncEntranceListAccessor) Raw() ([]FuncEntranceList, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []FuncEntranceList{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with FuncEntranceListAccessor.LoadData to preload everything
func (a *FuncEntranceListAccessor) GroupData() {
	a._dataID = map[json.Number]FuncEntranceList{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the FuncEntranceList uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *FuncEntranceListAccessor) ByID(identifier json.Number) (FuncEntranceList, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return FuncEntranceList{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
