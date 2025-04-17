package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type DialogueCondition struct {
	ID     float64 `json:"ID"`
	Param1 float64 `json:"Param1"`
	Param2 float64 `json:"Param2"`
	Type   string  `json:"Type"`
}
type DialogueConditionAccessor struct {
	_data   []DialogueCondition
	_dataID map[float64]DialogueCondition
}

// LoadData retrieves the data. Must be called before DialogueCondition.GroupData
func (a *DialogueConditionAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/DialogueCondition.json")
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
func (a *DialogueConditionAccessor) Raw() ([]DialogueCondition, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []DialogueCondition{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with DialogueConditionAccessor.LoadData to preload everything
func (a *DialogueConditionAccessor) GroupData() {
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the DialogueCondition uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *DialogueConditionAccessor) ByID(identifier float64) (DialogueCondition, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return DialogueCondition{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
