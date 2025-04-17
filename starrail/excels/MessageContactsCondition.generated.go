package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type MessageContactsCondition struct {
	FakeContactID         float64 `json:"FakeContactID"`
	ID                    float64 `json:"ID"`
	TruthMissionCondition float64 `json:"TruthMissionCondition"`
}
type MessageContactsConditionAccessor struct {
	_data              []MessageContactsCondition
	_dataFakeContactID map[float64]MessageContactsCondition
	_dataID            map[float64]MessageContactsCondition
}

// LoadData retrieves the data. Must be called before MessageContactsCondition.GroupData
func (a *MessageContactsConditionAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MessageContactsCondition.json")
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
func (a *MessageContactsConditionAccessor) Raw() ([]MessageContactsCondition, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MessageContactsCondition{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MessageContactsConditionAccessor.LoadData to preload everything
func (a *MessageContactsConditionAccessor) GroupData() {
	a._dataFakeContactID = map[float64]MessageContactsCondition{}
	a._dataID = map[float64]MessageContactsCondition{}
	for _, d := range a._data {
		a._dataFakeContactID[d.FakeContactID] = d
		a._dataID[d.ID] = d
	}
}

// ByFakeContactID returns the MessageContactsCondition uniquely identified by FakeContactID
//
// Error is only non-nil if the source errors out
func (a *MessageContactsConditionAccessor) ByFakeContactID(identifier float64) (MessageContactsCondition, error) {
	if a._dataFakeContactID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MessageContactsCondition{}, err
			}
		}
		a.GroupData()
	}
	return a._dataFakeContactID[identifier], nil
}

// ByID returns the MessageContactsCondition uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *MessageContactsConditionAccessor) ByID(identifier float64) (MessageContactsCondition, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MessageContactsCondition{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
