package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type RogueHandBookEvent struct {
	EventHandbookID         json.Number              `json:"EventHandbookID"`
	EventReward             json.Number              `json:"EventReward"`
	EventTitle              map[string]json.Number   `json:"EventTitle"`
	EventType               map[string]json.Number   `json:"EventType"`
	EventTypeList           []json.Number            `json:"EventTypeList"`
	ImageID                 json.Number              `json:"ImageID"`
	Order                   json.Number              `json:"Order"`
	UnlockHintDesc          map[string]json.Number   `json:"UnlockHintDesc"`
	UnlockNPCProgressIDList []map[string]json.Number `json:"UnlockNPCProgressIDList"`
}
type RogueHandBookEventAccessor struct {
	_data                []RogueHandBookEvent
	_dataEventHandbookID map[json.Number]RogueHandBookEvent
	_dataOrder           map[json.Number]RogueHandBookEvent
}

// LoadData retrieves the data. Must be called before RogueHandBookEvent.GroupData
func (a *RogueHandBookEventAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueHandBookEvent.json")
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
func (a *RogueHandBookEventAccessor) Raw() ([]RogueHandBookEvent, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueHandBookEvent{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueHandBookEventAccessor.LoadData to preload everything
func (a *RogueHandBookEventAccessor) GroupData() {
	a._dataEventHandbookID = map[json.Number]RogueHandBookEvent{}
	a._dataOrder = map[json.Number]RogueHandBookEvent{}
	for _, d := range a._data {
		a._dataEventHandbookID[d.EventHandbookID] = d
		a._dataOrder[d.Order] = d
	}
}

// ByEventHandbookID returns the RogueHandBookEvent uniquely identified by EventHandbookID
//
// Error is only non-nil if the source errors out
func (a *RogueHandBookEventAccessor) ByEventHandbookID(identifier json.Number) (RogueHandBookEvent, error) {
	if a._dataEventHandbookID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueHandBookEvent{}, err
			}
		}
		a.GroupData()
	}
	return a._dataEventHandbookID[identifier], nil
}

// ByOrder returns the RogueHandBookEvent uniquely identified by Order
//
// Error is only non-nil if the source errors out
func (a *RogueHandBookEventAccessor) ByOrder(identifier json.Number) (RogueHandBookEvent, error) {
	if a._dataOrder == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueHandBookEvent{}, err
			}
		}
		a.GroupData()
	}
	return a._dataOrder[identifier], nil
}
