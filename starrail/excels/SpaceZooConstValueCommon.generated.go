package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type SpaceZooConstValueCommon struct {
	ConstValueName string                        `json:"ConstValueName"`
	Value          SpaceZooConstValueCommonValue `json:"Value"`
}
type SpaceZooConstValueCommonValue struct {
	ArrayValue  []hash.IntValue `json:"ArrayValue"`
	IntValue    float64         `json:"IntValue"`
	StringValue string          `json:"StringValue"`
}
type SpaceZooConstValueCommonAccessor struct {
	_data               []SpaceZooConstValueCommon
	_dataConstValueName map[string]SpaceZooConstValueCommon
}

// LoadData retrieves the data. Must be called before SpaceZooConstValueCommon.GroupData
func (a *SpaceZooConstValueCommonAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/SpaceZooConstValueCommon.json")
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
func (a *SpaceZooConstValueCommonAccessor) Raw() ([]SpaceZooConstValueCommon, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []SpaceZooConstValueCommon{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with SpaceZooConstValueCommonAccessor.LoadData to preload everything
func (a *SpaceZooConstValueCommonAccessor) GroupData() {
	a._dataConstValueName = map[string]SpaceZooConstValueCommon{}
	for _, d := range a._data {
		a._dataConstValueName[d.ConstValueName] = d
	}
}

// ByConstValueName returns the SpaceZooConstValueCommon uniquely identified by ConstValueName
//
// Error is only non-nil if the source errors out
func (a *SpaceZooConstValueCommonAccessor) ByConstValueName(identifier string) (SpaceZooConstValueCommon, error) {
	if a._dataConstValueName == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return SpaceZooConstValueCommon{}, err
			}
		}
		a.GroupData()
	}
	return a._dataConstValueName[identifier], nil
}
