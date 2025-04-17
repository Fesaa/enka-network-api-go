package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type FantasticStoryBuffSlotID struct {
	BuffSlotID      float64 `json:"BuffSlotID"`
	UnlockChapterID float64 `json:"UnlockChapterID"`
}
type FantasticStoryBuffSlotIDAccessor struct {
	_data           []FantasticStoryBuffSlotID
	_dataBuffSlotID map[float64]FantasticStoryBuffSlotID
}

// LoadData retrieves the data. Must be called before FantasticStoryBuffSlotID.GroupData
func (a *FantasticStoryBuffSlotIDAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/FantasticStoryBuffSlotID.json")
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
func (a *FantasticStoryBuffSlotIDAccessor) Raw() ([]FantasticStoryBuffSlotID, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []FantasticStoryBuffSlotID{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with FantasticStoryBuffSlotIDAccessor.LoadData to preload everything
func (a *FantasticStoryBuffSlotIDAccessor) GroupData() {
	for _, d := range a._data {
		a._dataBuffSlotID[d.BuffSlotID] = d
	}
}

// ByBuffSlotID returns the FantasticStoryBuffSlotID uniquely identified by BuffSlotID
//
// Error is only non-nil if the source errors out
func (a *FantasticStoryBuffSlotIDAccessor) ByBuffSlotID(identifier float64) (FantasticStoryBuffSlotID, error) {
	if a._dataBuffSlotID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return FantasticStoryBuffSlotID{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBuffSlotID[identifier], nil
}
