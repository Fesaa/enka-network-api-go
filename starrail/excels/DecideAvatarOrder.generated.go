package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type DecideAvatarOrder struct {
	ItemID float64 `json:"ItemID"`
	Order  float64 `json:"Order"`
}
type DecideAvatarOrderAccessor struct {
	_data       []DecideAvatarOrder
	_dataItemID map[float64]DecideAvatarOrder
	_dataOrder  map[float64]DecideAvatarOrder
}

// LoadData retrieves the data. Must be called before DecideAvatarOrder.GroupData
func (a *DecideAvatarOrderAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/DecideAvatarOrder.json")
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
func (a *DecideAvatarOrderAccessor) Raw() ([]DecideAvatarOrder, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []DecideAvatarOrder{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with DecideAvatarOrderAccessor.LoadData to preload everything
func (a *DecideAvatarOrderAccessor) GroupData() {
	for _, d := range a._data {
		a._dataItemID[d.ItemID] = d
		a._dataOrder[d.Order] = d
	}
}

// ByItemID returns the DecideAvatarOrder uniquely identified by ItemID
//
// Error is only non-nil if the source errors out
func (a *DecideAvatarOrderAccessor) ByItemID(identifier float64) (DecideAvatarOrder, error) {
	if a._dataItemID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return DecideAvatarOrder{}, err
			}
		}
		a.GroupData()
	}
	return a._dataItemID[identifier], nil
}

// ByOrder returns the DecideAvatarOrder uniquely identified by Order
//
// Error is only non-nil if the source errors out
func (a *DecideAvatarOrderAccessor) ByOrder(identifier float64) (DecideAvatarOrder, error) {
	if a._dataOrder == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return DecideAvatarOrder{}, err
			}
		}
		a.GroupData()
	}
	return a._dataOrder[identifier], nil
}
