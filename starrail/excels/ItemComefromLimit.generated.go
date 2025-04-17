package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ItemComefromLimit struct {
	ComefromID   float64   `json:"ComefromID"`
	ID           float64   `json:"ID"`
	OpenParaList []float64 `json:"OpenParaList"`
	OpenType     string    `json:"OpenType"`
}
type ItemComefromLimitAccessor struct {
	_data           []ItemComefromLimit
	_dataComefromID map[float64]ItemComefromLimit
}

// LoadData retrieves the data. Must be called before ItemComefromLimit.GroupData
func (a *ItemComefromLimitAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ItemComefromLimit.json")
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
func (a *ItemComefromLimitAccessor) Raw() ([]ItemComefromLimit, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ItemComefromLimit{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ItemComefromLimitAccessor.LoadData to preload everything
func (a *ItemComefromLimitAccessor) GroupData() {
	for _, d := range a._data {
		a._dataComefromID[d.ComefromID] = d
	}
}

// ByComefromID returns the ItemComefromLimit uniquely identified by ComefromID
//
// Error is only non-nil if the source errors out
func (a *ItemComefromLimitAccessor) ByComefromID(identifier float64) (ItemComefromLimit, error) {
	if a._dataComefromID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ItemComefromLimit{}, err
			}
		}
		a.GroupData()
	}
	return a._dataComefromID[identifier], nil
}
