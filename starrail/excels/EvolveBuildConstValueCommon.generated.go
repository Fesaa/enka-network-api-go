package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"github.com/Fesaa/enka-network-api-go/hash"
	"io"
	"net/http"
)

type EvolveBuildConstValueCommon struct {
	ConstValueName string                           `json:"ConstValueName"`
	Value          EvolveBuildConstValueCommonValue `json:"Value"`
}
type EvolveBuildConstValueCommonValue struct {
	ArrayValue []hash.StringValue `json:"ArrayValue"`
	FloatValue json.Number        `json:"FloatValue"`
	IntValue   json.Number        `json:"IntValue"`
}
type EvolveBuildConstValueCommonAccessor struct {
	_data               []EvolveBuildConstValueCommon
	_dataConstValueName map[string]EvolveBuildConstValueCommon
}

// LoadData retrieves the data. Must be called before EvolveBuildConstValueCommon.GroupData
func (a *EvolveBuildConstValueCommonAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/EvolveBuildConstValueCommon.json")
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
func (a *EvolveBuildConstValueCommonAccessor) Raw() ([]EvolveBuildConstValueCommon, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []EvolveBuildConstValueCommon{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with EvolveBuildConstValueCommonAccessor.LoadData to preload everything
func (a *EvolveBuildConstValueCommonAccessor) GroupData() {
	a._dataConstValueName = map[string]EvolveBuildConstValueCommon{}
	for _, d := range a._data {
		a._dataConstValueName[d.ConstValueName] = d
	}
}

// ByConstValueName returns the EvolveBuildConstValueCommon uniquely identified by ConstValueName
//
// Error is only non-nil if the source errors out
func (a *EvolveBuildConstValueCommonAccessor) ByConstValueName(identifier string) (EvolveBuildConstValueCommon, error) {
	if a._dataConstValueName == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return EvolveBuildConstValueCommon{}, err
			}
		}
		a.GroupData()
	}
	return a._dataConstValueName[identifier], nil
}
