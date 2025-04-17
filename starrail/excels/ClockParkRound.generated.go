package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ClockParkRound struct {
	DiceSpecialDisplay bool    `json:"DiceSpecialDisplay"`
	RoundID            float64 `json:"RoundID"`
	RoundType          string  `json:"RoundType"`
}
type ClockParkRoundAccessor struct {
	_data        []ClockParkRound
	_dataRoundID map[float64]ClockParkRound
}

// LoadData retrieves the data. Must be called before ClockParkRound.GroupData
func (a *ClockParkRoundAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ClockParkRound.json")
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
func (a *ClockParkRoundAccessor) Raw() ([]ClockParkRound, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ClockParkRound{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ClockParkRoundAccessor.LoadData to preload everything
func (a *ClockParkRoundAccessor) GroupData() {
	a._dataRoundID = map[float64]ClockParkRound{}
	for _, d := range a._data {
		a._dataRoundID[d.RoundID] = d
	}
}

// ByRoundID returns the ClockParkRound uniquely identified by RoundID
//
// Error is only non-nil if the source errors out
func (a *ClockParkRoundAccessor) ByRoundID(identifier float64) (ClockParkRound, error) {
	if a._dataRoundID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ClockParkRound{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRoundID[identifier], nil
}
