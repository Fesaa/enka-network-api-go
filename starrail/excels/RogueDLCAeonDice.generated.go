package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type RogueDLCAeonDice struct {
	AeonDiceID           float64         `json:"AeonDiceID"`
	DescParam            []hash.IntValue `json:"DescParam"`
	DiceIcon             string          `json:"DiceIcon"`
	DiceModel            string          `json:"DiceModel"`
	DiceShortDesc        hash.Hash       `json:"DiceShortDesc"`
	DiceStartEffectDesc  hash.Hash       `json:"DiceStartEffectDesc"`
	ExtraEffect          []float64       `json:"ExtraEffect"`
	SoundReRoll          string          `json:"SoundReRoll"`
	SoundRoll            string          `json:"SoundRoll"`
	SoundSuspensionStart string          `json:"SoundSuspensionStart"`
	SoundSuspensionStop  string          `json:"SoundSuspensionStop"`
	StartDescParam       []float64       `json:"StartDescParam"`
}
type RogueDLCAeonDiceAccessor struct {
	_data                     []RogueDLCAeonDice
	_dataAeonDiceID           map[float64]RogueDLCAeonDice
	_dataDiceIcon             map[string]RogueDLCAeonDice
	_dataDiceModel            map[string]RogueDLCAeonDice
	_dataSoundReRoll          map[string]RogueDLCAeonDice
	_dataSoundRoll            map[string]RogueDLCAeonDice
	_dataSoundSuspensionStart map[string]RogueDLCAeonDice
	_dataSoundSuspensionStop  map[string]RogueDLCAeonDice
}

// LoadData retrieves the data. Must be called before RogueDLCAeonDice.GroupData
func (a *RogueDLCAeonDiceAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueDLCAeonDice.json")
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
func (a *RogueDLCAeonDiceAccessor) Raw() ([]RogueDLCAeonDice, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueDLCAeonDice{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueDLCAeonDiceAccessor.LoadData to preload everything
func (a *RogueDLCAeonDiceAccessor) GroupData() {
	a._dataAeonDiceID = map[float64]RogueDLCAeonDice{}
	a._dataDiceIcon = map[string]RogueDLCAeonDice{}
	a._dataDiceModel = map[string]RogueDLCAeonDice{}
	a._dataSoundReRoll = map[string]RogueDLCAeonDice{}
	a._dataSoundRoll = map[string]RogueDLCAeonDice{}
	a._dataSoundSuspensionStart = map[string]RogueDLCAeonDice{}
	a._dataSoundSuspensionStop = map[string]RogueDLCAeonDice{}
	for _, d := range a._data {
		a._dataAeonDiceID[d.AeonDiceID] = d
		a._dataDiceIcon[d.DiceIcon] = d
		a._dataDiceModel[d.DiceModel] = d
		a._dataSoundReRoll[d.SoundReRoll] = d
		a._dataSoundRoll[d.SoundRoll] = d
		a._dataSoundSuspensionStart[d.SoundSuspensionStart] = d
		a._dataSoundSuspensionStop[d.SoundSuspensionStop] = d
	}
}

// ByAeonDiceID returns the RogueDLCAeonDice uniquely identified by AeonDiceID
//
// Error is only non-nil if the source errors out
func (a *RogueDLCAeonDiceAccessor) ByAeonDiceID(identifier float64) (RogueDLCAeonDice, error) {
	if a._dataAeonDiceID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueDLCAeonDice{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAeonDiceID[identifier], nil
}

// ByDiceIcon returns the RogueDLCAeonDice uniquely identified by DiceIcon
//
// Error is only non-nil if the source errors out
func (a *RogueDLCAeonDiceAccessor) ByDiceIcon(identifier string) (RogueDLCAeonDice, error) {
	if a._dataDiceIcon == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueDLCAeonDice{}, err
			}
		}
		a.GroupData()
	}
	return a._dataDiceIcon[identifier], nil
}

// ByDiceModel returns the RogueDLCAeonDice uniquely identified by DiceModel
//
// Error is only non-nil if the source errors out
func (a *RogueDLCAeonDiceAccessor) ByDiceModel(identifier string) (RogueDLCAeonDice, error) {
	if a._dataDiceModel == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueDLCAeonDice{}, err
			}
		}
		a.GroupData()
	}
	return a._dataDiceModel[identifier], nil
}

// BySoundReRoll returns the RogueDLCAeonDice uniquely identified by SoundReRoll
//
// Error is only non-nil if the source errors out
func (a *RogueDLCAeonDiceAccessor) BySoundReRoll(identifier string) (RogueDLCAeonDice, error) {
	if a._dataSoundReRoll == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueDLCAeonDice{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSoundReRoll[identifier], nil
}

// BySoundRoll returns the RogueDLCAeonDice uniquely identified by SoundRoll
//
// Error is only non-nil if the source errors out
func (a *RogueDLCAeonDiceAccessor) BySoundRoll(identifier string) (RogueDLCAeonDice, error) {
	if a._dataSoundRoll == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueDLCAeonDice{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSoundRoll[identifier], nil
}

// BySoundSuspensionStart returns the RogueDLCAeonDice uniquely identified by SoundSuspensionStart
//
// Error is only non-nil if the source errors out
func (a *RogueDLCAeonDiceAccessor) BySoundSuspensionStart(identifier string) (RogueDLCAeonDice, error) {
	if a._dataSoundSuspensionStart == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueDLCAeonDice{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSoundSuspensionStart[identifier], nil
}

// BySoundSuspensionStop returns the RogueDLCAeonDice uniquely identified by SoundSuspensionStop
//
// Error is only non-nil if the source errors out
func (a *RogueDLCAeonDiceAccessor) BySoundSuspensionStop(identifier string) (RogueDLCAeonDice, error) {
	if a._dataSoundSuspensionStop == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueDLCAeonDice{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSoundSuspensionStop[identifier], nil
}
