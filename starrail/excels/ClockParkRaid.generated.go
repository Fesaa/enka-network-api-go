package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ClockParkRaid struct {
	RaidID             float64 `json:"RaidID"`
	RaidMapinfo        float64 `json:"RaidMapinfo"`
	RaidUnlockProgress float64 `json:"RaidUnlockProgress"`
}
type ClockParkRaidAccessor struct {
	_data                   []ClockParkRaid
	_dataRaidID             map[float64]ClockParkRaid
	_dataRaidMapinfo        map[float64]ClockParkRaid
	_dataRaidUnlockProgress map[float64]ClockParkRaid
}

// LoadData retrieves the data. Must be called before ClockParkRaid.GroupData
func (a *ClockParkRaidAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ClockParkRaid.json")
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
func (a *ClockParkRaidAccessor) Raw() ([]ClockParkRaid, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ClockParkRaid{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ClockParkRaidAccessor.LoadData to preload everything
func (a *ClockParkRaidAccessor) GroupData() {
	a._dataRaidID = map[float64]ClockParkRaid{}
	a._dataRaidMapinfo = map[float64]ClockParkRaid{}
	a._dataRaidUnlockProgress = map[float64]ClockParkRaid{}
	for _, d := range a._data {
		a._dataRaidID[d.RaidID] = d
		a._dataRaidMapinfo[d.RaidMapinfo] = d
		a._dataRaidUnlockProgress[d.RaidUnlockProgress] = d
	}
}

// ByRaidID returns the ClockParkRaid uniquely identified by RaidID
//
// Error is only non-nil if the source errors out
func (a *ClockParkRaidAccessor) ByRaidID(identifier float64) (ClockParkRaid, error) {
	if a._dataRaidID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ClockParkRaid{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRaidID[identifier], nil
}

// ByRaidMapinfo returns the ClockParkRaid uniquely identified by RaidMapinfo
//
// Error is only non-nil if the source errors out
func (a *ClockParkRaidAccessor) ByRaidMapinfo(identifier float64) (ClockParkRaid, error) {
	if a._dataRaidMapinfo == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ClockParkRaid{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRaidMapinfo[identifier], nil
}

// ByRaidUnlockProgress returns the ClockParkRaid uniquely identified by RaidUnlockProgress
//
// Error is only non-nil if the source errors out
func (a *ClockParkRaidAccessor) ByRaidUnlockProgress(identifier float64) (ClockParkRaid, error) {
	if a._dataRaidUnlockProgress == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ClockParkRaid{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRaidUnlockProgress[identifier], nil
}
