package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type PsActivity struct {
	ActivityID   float64   `json:"ActivityID"`
	ObjectIDList []float64 `json:"ObjectIDList"`
	Description  hash.Hash `json:"description"`
	Name         hash.Hash `json:"name"`
	Task         hash.Hash `json:"task"`
}
type PsActivityAccessor struct {
	_data           []PsActivity
	_dataActivityID map[float64]PsActivity
}

// LoadData retrieves the data. Must be called before PsActivity.GroupData
func (a *PsActivityAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PsActivity.json")
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
func (a *PsActivityAccessor) Raw() ([]PsActivity, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PsActivity{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PsActivityAccessor.LoadData to preload everything
func (a *PsActivityAccessor) GroupData() {
	a._dataActivityID = map[float64]PsActivity{}
	for _, d := range a._data {
		a._dataActivityID[d.ActivityID] = d
	}
}

// ByActivityID returns the PsActivity uniquely identified by ActivityID
//
// Error is only non-nil if the source errors out
func (a *PsActivityAccessor) ByActivityID(identifier float64) (PsActivity, error) {
	if a._dataActivityID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PsActivity{}, err
			}
		}
		a.GroupData()
	}
	return a._dataActivityID[identifier], nil
}
