package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ActivityFeverTimeUnderline struct {
	AvailableBuffID float64 `json:"AvailableBuffID"`
	ExtraEffectID   float64 `json:"ExtraEffectID"`
}
type ActivityFeverTimeUnderlineAccessor struct {
	_data                []ActivityFeverTimeUnderline
	_dataAvailableBuffID map[float64]ActivityFeverTimeUnderline
}

// LoadData retrieves the data. Must be called before ActivityFeverTimeUnderline.GroupData
func (a *ActivityFeverTimeUnderlineAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActivityFeverTimeUnderline.json")
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
func (a *ActivityFeverTimeUnderlineAccessor) Raw() ([]ActivityFeverTimeUnderline, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActivityFeverTimeUnderline{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ActivityFeverTimeUnderlineAccessor.LoadData to preload everything
func (a *ActivityFeverTimeUnderlineAccessor) GroupData() {
	a._dataAvailableBuffID = map[float64]ActivityFeverTimeUnderline{}
	for _, d := range a._data {
		a._dataAvailableBuffID[d.AvailableBuffID] = d
	}
}

// ByAvailableBuffID returns the ActivityFeverTimeUnderline uniquely identified by AvailableBuffID
//
// Error is only non-nil if the source errors out
func (a *ActivityFeverTimeUnderlineAccessor) ByAvailableBuffID(identifier float64) (ActivityFeverTimeUnderline, error) {
	if a._dataAvailableBuffID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityFeverTimeUnderline{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAvailableBuffID[identifier], nil
}
