package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type BattleActionEventConfig struct {
	AbilityName      string          `json:"AbilityName"`
	ActiveDefault    bool            `json:"ActiveDefault"`
	BriefDescription hash.Hash       `json:"BriefDescription"`
	EventID          float64         `json:"EventID"`
	EventName        hash.Hash       `json:"EventName"`
	FullDescription  hash.Hash       `json:"FullDescription"`
	IconPath         string          `json:"IconPath"`
	InitialInterval  float64         `json:"InitialInterval"`
	Interval         float64         `json:"Interval"`
	ParamList        []hash.IntValue `json:"ParamList"`
}
type BattleActionEventConfigAccessor struct {
	_data        []BattleActionEventConfig
	_dataEventID map[float64]BattleActionEventConfig
}

// LoadData retrieves the data. Must be called before BattleActionEventConfig.GroupData
func (a *BattleActionEventConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/BattleActionEventConfig.json")
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
func (a *BattleActionEventConfigAccessor) Raw() ([]BattleActionEventConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []BattleActionEventConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with BattleActionEventConfigAccessor.LoadData to preload everything
func (a *BattleActionEventConfigAccessor) GroupData() {
	a._dataEventID = map[float64]BattleActionEventConfig{}
	for _, d := range a._data {
		a._dataEventID[d.EventID] = d
	}
}

// ByEventID returns the BattleActionEventConfig uniquely identified by EventID
//
// Error is only non-nil if the source errors out
func (a *BattleActionEventConfigAccessor) ByEventID(identifier float64) (BattleActionEventConfig, error) {
	if a._dataEventID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return BattleActionEventConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataEventID[identifier], nil
}
