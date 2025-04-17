package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type RandomEventChoice struct {
	ChoiceID        json.Number            `json:"ChoiceID"`
	EventBuffDay    json.Number            `json:"EventBuffDay"`
	EventCostOption json.Number            `json:"EventCostOption"`
	EventRewardBuff json.Number            `json:"EventRewardBuff"`
	IsCancel        bool                   `json:"IsCancel"`
	Option          map[string]json.Number `json:"Option"`
}
type RandomEventChoiceAccessor struct {
	_data         []RandomEventChoice
	_dataChoiceID map[json.Number]RandomEventChoice
}

// LoadData retrieves the data. Must be called before RandomEventChoice.GroupData
func (a *RandomEventChoiceAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RandomEventChoice.json")
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
func (a *RandomEventChoiceAccessor) Raw() ([]RandomEventChoice, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RandomEventChoice{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RandomEventChoiceAccessor.LoadData to preload everything
func (a *RandomEventChoiceAccessor) GroupData() {
	a._dataChoiceID = map[json.Number]RandomEventChoice{}
	for _, d := range a._data {
		a._dataChoiceID[d.ChoiceID] = d
	}
}

// ByChoiceID returns the RandomEventChoice uniquely identified by ChoiceID
//
// Error is only non-nil if the source errors out
func (a *RandomEventChoiceAccessor) ByChoiceID(identifier json.Number) (RandomEventChoice, error) {
	if a._dataChoiceID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RandomEventChoice{}, err
			}
		}
		a.GroupData()
	}
	return a._dataChoiceID[identifier], nil
}
