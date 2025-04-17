package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type RogueBonus struct {
	BonusDesc  hash.Hash `json:"BonusDesc"`
	BonusEvent float64   `json:"BonusEvent"`
	BonusID    float64   `json:"BonusID"`
	BonusIcon  string    `json:"BonusIcon"`
	BonusTag   hash.Hash `json:"BonusTag"`
	BonusTitle hash.Hash `json:"BonusTitle"`
}
type RogueBonusAccessor struct {
	_data           []RogueBonus
	_dataBonusEvent map[float64]RogueBonus
	_dataBonusID    map[float64]RogueBonus
}

// LoadData retrieves the data. Must be called before RogueBonus.GroupData
func (a *RogueBonusAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueBonus.json")
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
func (a *RogueBonusAccessor) Raw() ([]RogueBonus, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueBonus{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueBonusAccessor.LoadData to preload everything
func (a *RogueBonusAccessor) GroupData() {
	for _, d := range a._data {
		a._dataBonusEvent[d.BonusEvent] = d
		a._dataBonusID[d.BonusID] = d
	}
}

// ByBonusEvent returns the RogueBonus uniquely identified by BonusEvent
//
// Error is only non-nil if the source errors out
func (a *RogueBonusAccessor) ByBonusEvent(identifier float64) (RogueBonus, error) {
	if a._dataBonusEvent == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueBonus{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBonusEvent[identifier], nil
}

// ByBonusID returns the RogueBonus uniquely identified by BonusID
//
// Error is only non-nil if the source errors out
func (a *RogueBonusAccessor) ByBonusID(identifier float64) (RogueBonus, error) {
	if a._dataBonusID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueBonus{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBonusID[identifier], nil
}
