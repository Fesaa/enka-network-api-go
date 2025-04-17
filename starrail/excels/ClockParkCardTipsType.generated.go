package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type ClockParkCardTipsType struct {
	CardTips       map[string]json.Number `json:"CardTips"`
	CardTipsTypeID string                 `json:"CardTipsTypeID"`
	CardTipsDetail map[string]json.Number `json:"CardTips_Detail"`
}
type ClockParkCardTipsTypeAccessor struct {
	_data               []ClockParkCardTipsType
	_dataCardTipsTypeID map[string]ClockParkCardTipsType
}

// LoadData retrieves the data. Must be called before ClockParkCardTipsType.GroupData
func (a *ClockParkCardTipsTypeAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ClockParkCardTipsType.json")
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
func (a *ClockParkCardTipsTypeAccessor) Raw() ([]ClockParkCardTipsType, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ClockParkCardTipsType{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ClockParkCardTipsTypeAccessor.LoadData to preload everything
func (a *ClockParkCardTipsTypeAccessor) GroupData() {
	a._dataCardTipsTypeID = map[string]ClockParkCardTipsType{}
	for _, d := range a._data {
		a._dataCardTipsTypeID[d.CardTipsTypeID] = d
	}
}

// ByCardTipsTypeID returns the ClockParkCardTipsType uniquely identified by CardTipsTypeID
//
// Error is only non-nil if the source errors out
func (a *ClockParkCardTipsTypeAccessor) ByCardTipsTypeID(identifier string) (ClockParkCardTipsType, error) {
	if a._dataCardTipsTypeID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ClockParkCardTipsType{}, err
			}
		}
		a.GroupData()
	}
	return a._dataCardTipsTypeID[identifier], nil
}
