package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type ClockParkLottery struct {
	LotteryAttributeGain map[string]json.Number `json:"LotteryAttributeGain"`
	LotteryID            json.Number            `json:"LotteryID"`
	LotteryType          json.Number            `json:"LotteryType"`
	Weight               json.Number            `json:"Weight"`
}
type ClockParkLotteryAccessor struct {
	_data          []ClockParkLottery
	_dataLotteryID map[json.Number]ClockParkLottery
}

// LoadData retrieves the data. Must be called before ClockParkLottery.GroupData
func (a *ClockParkLotteryAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ClockParkLottery.json")
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
func (a *ClockParkLotteryAccessor) Raw() ([]ClockParkLottery, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ClockParkLottery{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ClockParkLotteryAccessor.LoadData to preload everything
func (a *ClockParkLotteryAccessor) GroupData() {
	a._dataLotteryID = map[json.Number]ClockParkLottery{}
	for _, d := range a._data {
		a._dataLotteryID[d.LotteryID] = d
	}
}

// ByLotteryID returns the ClockParkLottery uniquely identified by LotteryID
//
// Error is only non-nil if the source errors out
func (a *ClockParkLotteryAccessor) ByLotteryID(identifier json.Number) (ClockParkLottery, error) {
	if a._dataLotteryID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ClockParkLottery{}, err
			}
		}
		a.GroupData()
	}
	return a._dataLotteryID[identifier], nil
}
