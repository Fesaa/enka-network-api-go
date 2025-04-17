package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type RogueDLCAeonCabinet struct {
	CabinetDesc                  map[string]json.Number   `json:"CabinetDesc"`
	CabinetID                    json.Number              `json:"CabinetID"`
	CabinetIcon                  string                   `json:"CabinetIcon"`
	CabinetMissionDesc           map[string]json.Number   `json:"CabinetMissionDesc"`
	CabinetName                  map[string]json.Number   `json:"CabinetName"`
	CabinetType                  string                   `json:"CabinetType"`
	DescParam                    []json.Number            `json:"DescParam"`
	FinishAeonDimensionPointList []map[string]json.Number `json:"FinishAeonDimensionPointList"`
	QuestID                      json.Number              `json:"QuestID"`
	Sort                         json.Number              `json:"Sort"`
	UnlockCabinetID              []json.Number            `json:"UnlockCabinetID"`
}
type RogueDLCAeonCabinetAccessor struct {
	_data            []RogueDLCAeonCabinet
	_dataCabinetIcon map[string]RogueDLCAeonCabinet
}

// LoadData retrieves the data. Must be called before RogueDLCAeonCabinet.GroupData
func (a *RogueDLCAeonCabinetAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueDLCAeonCabinet.json")
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
func (a *RogueDLCAeonCabinetAccessor) Raw() ([]RogueDLCAeonCabinet, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueDLCAeonCabinet{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueDLCAeonCabinetAccessor.LoadData to preload everything
func (a *RogueDLCAeonCabinetAccessor) GroupData() {
	a._dataCabinetIcon = map[string]RogueDLCAeonCabinet{}
	for _, d := range a._data {
		a._dataCabinetIcon[d.CabinetIcon] = d
	}
}

// ByCabinetIcon returns the RogueDLCAeonCabinet uniquely identified by CabinetIcon
//
// Error is only non-nil if the source errors out
func (a *RogueDLCAeonCabinetAccessor) ByCabinetIcon(identifier string) (RogueDLCAeonCabinet, error) {
	if a._dataCabinetIcon == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueDLCAeonCabinet{}, err
			}
		}
		a.GroupData()
	}
	return a._dataCabinetIcon[identifier], nil
}
