package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type ActivityResidentPanel struct {
	DisplayItemList       []ActivityResidentPanelDisplayItemList  `json:"DisplayItemList"`
	DisplayItemManualSort bool                                    `json:"DisplayItemManualSort"`
	FinishConditions      []ActivityResidentPanelFinishConditions `json:"FinishConditions"`
	IntroDesc             hash.Hash                               `json:"IntroDesc"`
	PanelDesc             hash.Hash                               `json:"PanelDesc"`
	PanelID               float64                                 `json:"PanelID"`
	SortWeight            float64                                 `json:"SortWeight"`
}
type ActivityResidentPanelDisplayItemList struct {
	ItemID float64 `json:"ItemID"`
}
type ActivityResidentPanelFinishConditions struct {
	Param string `json:"Param"`
	Type  string `json:"Type"`
}
type ActivityResidentPanelAccessor struct {
	_data           []ActivityResidentPanel
	_dataPanelID    map[float64]ActivityResidentPanel
	_dataSortWeight map[float64]ActivityResidentPanel
}

// LoadData retrieves the data. Must be called before ActivityResidentPanel.GroupData
func (a *ActivityResidentPanelAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActivityResidentPanel.json")
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
func (a *ActivityResidentPanelAccessor) Raw() ([]ActivityResidentPanel, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActivityResidentPanel{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ActivityResidentPanelAccessor.LoadData to preload everything
func (a *ActivityResidentPanelAccessor) GroupData() {
	a._dataPanelID = map[float64]ActivityResidentPanel{}
	a._dataSortWeight = map[float64]ActivityResidentPanel{}
	for _, d := range a._data {
		a._dataPanelID[d.PanelID] = d
		a._dataSortWeight[d.SortWeight] = d
	}
}

// ByPanelID returns the ActivityResidentPanel uniquely identified by PanelID
//
// Error is only non-nil if the source errors out
func (a *ActivityResidentPanelAccessor) ByPanelID(identifier float64) (ActivityResidentPanel, error) {
	if a._dataPanelID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityResidentPanel{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPanelID[identifier], nil
}

// BySortWeight returns the ActivityResidentPanel uniquely identified by SortWeight
//
// Error is only non-nil if the source errors out
func (a *ActivityResidentPanelAccessor) BySortWeight(identifier float64) (ActivityResidentPanel, error) {
	if a._dataSortWeight == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityResidentPanel{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSortWeight[identifier], nil
}
