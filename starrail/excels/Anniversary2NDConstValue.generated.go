package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type Anniversary2NDConstValue struct {
	ConstValueName string        `json:"ConstValueName"`
	Value          hash.IntValue `json:"Value"`
}
type Anniversary2NDConstValueAccessor struct {
	_data               []Anniversary2NDConstValue
	_dataConstValueName map[string]Anniversary2NDConstValue
}

// LoadData retrieves the data. Must be called before Anniversary2NDConstValue.GroupData
func (a *Anniversary2NDConstValueAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/Anniversary2NDConstValue.json")
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
func (a *Anniversary2NDConstValueAccessor) Raw() ([]Anniversary2NDConstValue, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []Anniversary2NDConstValue{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with Anniversary2NDConstValueAccessor.LoadData to preload everything
func (a *Anniversary2NDConstValueAccessor) GroupData() {
	a._dataConstValueName = map[string]Anniversary2NDConstValue{}
	for _, d := range a._data {
		a._dataConstValueName[d.ConstValueName] = d
	}
}

// ByConstValueName returns the Anniversary2NDConstValue uniquely identified by ConstValueName
//
// Error is only non-nil if the source errors out
func (a *Anniversary2NDConstValueAccessor) ByConstValueName(identifier string) (Anniversary2NDConstValue, error) {
	if a._dataConstValueName == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return Anniversary2NDConstValue{}, err
			}
		}
		a.GroupData()
	}
	return a._dataConstValueName[identifier], nil
}
