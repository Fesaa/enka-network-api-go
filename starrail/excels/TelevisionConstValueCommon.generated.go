package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
)

type TelevisionConstValueCommon struct {
	ConstValueName string                          `json:"ConstValueName"`
	Value          TelevisionConstValueCommonValue `json:"Value"`
}
type TelevisionConstValueCommonValue struct {
	ArrayValue []hash.StringValue `json:"ArrayValue"`
	IntValue   json.Number        `json:"IntValue"`
}
type TelevisionConstValueCommonAccessor struct {
	_data               []TelevisionConstValueCommon
	_dataConstValueName map[string]TelevisionConstValueCommon
}

// LoadData retrieves the data. Must be called before TelevisionConstValueCommon.GroupData
func (a *TelevisionConstValueCommonAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TelevisionConstValueCommon.json")
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
func (a *TelevisionConstValueCommonAccessor) Raw() ([]TelevisionConstValueCommon, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TelevisionConstValueCommon{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with TelevisionConstValueCommonAccessor.LoadData to preload everything
func (a *TelevisionConstValueCommonAccessor) GroupData() {
	a._dataConstValueName = map[string]TelevisionConstValueCommon{}
	for _, d := range a._data {
		a._dataConstValueName[d.ConstValueName] = d
	}
}

// ByConstValueName returns the TelevisionConstValueCommon uniquely identified by ConstValueName
//
// Error is only non-nil if the source errors out
func (a *TelevisionConstValueCommonAccessor) ByConstValueName(identifier string) (TelevisionConstValueCommon, error) {
	if a._dataConstValueName == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TelevisionConstValueCommon{}, err
			}
		}
		a.GroupData()
	}
	return a._dataConstValueName[identifier], nil
}
