package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ActivityFinishWayPunkLord struct {
	FinishID    float64 `json:"FinishID"`
	FinishPoint float64 `json:"FinishPoint"`
	FinishRare  string  `json:"FinishRare"`
}
type ActivityFinishWayPunkLordAccessor struct {
	_data         []ActivityFinishWayPunkLord
	_dataFinishID map[float64]ActivityFinishWayPunkLord
}

// LoadData retrieves the data. Must be called before ActivityFinishWayPunkLord.GroupData
func (a *ActivityFinishWayPunkLordAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActivityFinishWayPunkLord.json")
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
func (a *ActivityFinishWayPunkLordAccessor) Raw() ([]ActivityFinishWayPunkLord, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActivityFinishWayPunkLord{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ActivityFinishWayPunkLordAccessor.LoadData to preload everything
func (a *ActivityFinishWayPunkLordAccessor) GroupData() {
	a._dataFinishID = map[float64]ActivityFinishWayPunkLord{}
	for _, d := range a._data {
		a._dataFinishID[d.FinishID] = d
	}
}

// ByFinishID returns the ActivityFinishWayPunkLord uniquely identified by FinishID
//
// Error is only non-nil if the source errors out
func (a *ActivityFinishWayPunkLordAccessor) ByFinishID(identifier float64) (ActivityFinishWayPunkLord, error) {
	if a._dataFinishID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityFinishWayPunkLord{}, err
			}
		}
		a.GroupData()
	}
	return a._dataFinishID[identifier], nil
}
