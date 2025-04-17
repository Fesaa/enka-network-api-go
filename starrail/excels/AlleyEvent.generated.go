package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type AlleyEvent struct {
	EventFinishTitle   hash.Hash                    `json:"EventFinishTitle"`
	EventID            float64                      `json:"EventID"`
	EventIcon          string                       `json:"EventIcon"`
	EventNewOrderTips  hash.Hash                    `json:"EventNewOrderTips"`
	EventPic           string                       `json:"EventPic"`
	EventPriority      float64                      `json:"EventPriority"`
	EventShopContent   hash.Hash                    `json:"EventShopContent"`
	EventShopFinish    hash.Hash                    `json:"EventShopFinish"`
	EventShopOrder     hash.Hash                    `json:"EventShopOrder"`
	EventShopTitle     hash.Hash                    `json:"EventShopTitle"`
	EventTitle         hash.Hash                    `json:"EventTitle"`
	EventType          string                       `json:"EventType"`
	MapEntranceID      float64                      `json:"MapEntranceID"`
	MappingInfoID      float64                      `json:"MappingInfoID"`
	RewardID           float64                      `json:"RewardID"`
	StartMissionIDList []float64                    `json:"StartMissionIDList"`
	UnlockConditions   []AlleyEventUnlockConditions `json:"UnlockConditions"`
}
type AlleyEventUnlockConditions struct {
	CAOAPDCCPCA float64 `json:"CAOAPDCCPCA"`
	PICHIHHCOCB string  `json:"PICHIHHCOCB"`
}
type AlleyEventAccessor struct {
	_data              []AlleyEvent
	_dataEventPriority map[float64]AlleyEvent
	_dataEventID       map[float64]AlleyEvent
}

// LoadData retrieves the data. Must be called before AlleyEvent.GroupData
func (a *AlleyEventAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AlleyEvent.json")
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
func (a *AlleyEventAccessor) Raw() ([]AlleyEvent, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AlleyEvent{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AlleyEventAccessor.LoadData to preload everything
func (a *AlleyEventAccessor) GroupData() {
	for _, d := range a._data {
		a._dataEventPriority[d.EventPriority] = d
		a._dataEventID[d.EventID] = d
	}
}

// ByEventPriority returns the AlleyEvent uniquely identified by EventPriority
//
// Error is only non-nil if the source errors out
func (a *AlleyEventAccessor) ByEventPriority(identifier float64) (AlleyEvent, error) {
	if a._dataEventPriority == nil {
		err := a.LoadData()
		if err != nil {
			return AlleyEvent{}, err
		}
		a.GroupData()
	}
	return a._dataEventPriority[identifier], nil
}

// ByEventID returns the AlleyEvent uniquely identified by EventID
//
// Error is only non-nil if the source errors out
func (a *AlleyEventAccessor) ByEventID(identifier float64) (AlleyEvent, error) {
	if a._dataEventID == nil {
		err := a.LoadData()
		if err != nil {
			return AlleyEvent{}, err
		}
		a.GroupData()
	}
	return a._dataEventID[identifier], nil
}
