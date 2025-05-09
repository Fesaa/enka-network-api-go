package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ItemGotoData struct {
	GotoID float64 `json:"GotoID"`
	ID     float64 `json:"ID"`
}
type ItemGotoDataAccessor struct {
	_data       []ItemGotoData
	_dataGotoID map[float64]ItemGotoData
	_dataID     map[float64]ItemGotoData
}

// LoadData retrieves the data. Must be called before ItemGotoData.GroupData
func (a *ItemGotoDataAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ItemGotoData.json")
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
func (a *ItemGotoDataAccessor) Raw() ([]ItemGotoData, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ItemGotoData{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ItemGotoDataAccessor.LoadData to preload everything
func (a *ItemGotoDataAccessor) GroupData() {
	a._dataGotoID = map[float64]ItemGotoData{}
	a._dataID = map[float64]ItemGotoData{}
	for _, d := range a._data {
		a._dataGotoID[d.GotoID] = d
		a._dataID[d.ID] = d
	}
}

// ByGotoID returns the ItemGotoData uniquely identified by GotoID
//
// Error is only non-nil if the source errors out
func (a *ItemGotoDataAccessor) ByGotoID(identifier float64) (ItemGotoData, error) {
	if a._dataGotoID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ItemGotoData{}, err
			}
		}
		a.GroupData()
	}
	return a._dataGotoID[identifier], nil
}

// ByID returns the ItemGotoData uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *ItemGotoDataAccessor) ByID(identifier float64) (ItemGotoData, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ItemGotoData{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
