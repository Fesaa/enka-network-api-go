package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type MonopolyConstValue struct {
	ConstValueName string        `json:"ConstValueName"`
	Value          hash.IntValue `json:"Value"`
}
type MonopolyConstValueAccessor struct {
	_data               []MonopolyConstValue
	_dataConstValueName map[string]MonopolyConstValue
}

// LoadData retrieves the data. Must be called before MonopolyConstValue.GroupData
func (a *MonopolyConstValueAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MonopolyConstValue.json")
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
func (a *MonopolyConstValueAccessor) Raw() ([]MonopolyConstValue, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MonopolyConstValue{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MonopolyConstValueAccessor.LoadData to preload everything
func (a *MonopolyConstValueAccessor) GroupData() {
	a._dataConstValueName = map[string]MonopolyConstValue{}
	for _, d := range a._data {
		a._dataConstValueName[d.ConstValueName] = d
	}
}

// ByConstValueName returns the MonopolyConstValue uniquely identified by ConstValueName
//
// Error is only non-nil if the source errors out
func (a *MonopolyConstValueAccessor) ByConstValueName(identifier string) (MonopolyConstValue, error) {
	if a._dataConstValueName == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MonopolyConstValue{}, err
			}
		}
		a.GroupData()
	}
	return a._dataConstValueName[identifier], nil
}
