package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type RogueTournHandBookEvent struct {
	EventHandbookID         float64                                          `json:"EventHandbookID"`
	EventTitle              hash.Hash                                        `json:"EventTitle"`
	ImageID                 float64                                          `json:"ImageID"`
	IsUsed                  bool                                             `json:"IsUsed"`
	Priority                float64                                          `json:"Priority"`
	TypeDisplayID           float64                                          `json:"TypeDisplayID"`
	UnlockDisplayID         float64                                          `json:"UnlockDisplayID"`
	UnlockNPCProgressIDList []RogueTournHandBookEventUnlockNPCProgressIDList `json:"UnlockNPCProgressIDList"`
}
type RogueTournHandBookEventUnlockNPCProgressIDList struct {
	MBNKLBEBOHB float64 `json:"MBNKLBEBOHB"`
	NNDEOKKKKPE float64 `json:"NNDEOKKKKPE"`
}
type RogueTournHandBookEventAccessor struct {
	_data                []RogueTournHandBookEvent
	_dataPriority        map[float64]RogueTournHandBookEvent
	_dataEventHandbookID map[float64]RogueTournHandBookEvent
}

// LoadData retrieves the data. Must be called before RogueTournHandBookEvent.GroupData
func (a *RogueTournHandBookEventAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueTournHandBookEvent.json")
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
func (a *RogueTournHandBookEventAccessor) Raw() ([]RogueTournHandBookEvent, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueTournHandBookEvent{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueTournHandBookEventAccessor.LoadData to preload everything
func (a *RogueTournHandBookEventAccessor) GroupData() {
	for _, d := range a._data {
		a._dataPriority[d.Priority] = d
		a._dataEventHandbookID[d.EventHandbookID] = d
	}
}

// ByPriority returns the RogueTournHandBookEvent uniquely identified by Priority
//
// Error is only non-nil if the source errors out
func (a *RogueTournHandBookEventAccessor) ByPriority(identifier float64) (RogueTournHandBookEvent, error) {
	if a._dataPriority == nil {
		err := a.LoadData()
		if err != nil {
			return RogueTournHandBookEvent{}, err
		}
		a.GroupData()
	}
	return a._dataPriority[identifier], nil
}

// ByEventHandbookID returns the RogueTournHandBookEvent uniquely identified by EventHandbookID
//
// Error is only non-nil if the source errors out
func (a *RogueTournHandBookEventAccessor) ByEventHandbookID(identifier float64) (RogueTournHandBookEvent, error) {
	if a._dataEventHandbookID == nil {
		err := a.LoadData()
		if err != nil {
			return RogueTournHandBookEvent{}, err
		}
		a.GroupData()
	}
	return a._dataEventHandbookID[identifier], nil
}
