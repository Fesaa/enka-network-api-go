package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type TrainPartyConstValueCommon struct {
	ConstValueName string           `json:"ConstValueName"`
	Value          hash.StringValue `json:"Value"`
}
type TrainPartyConstValueCommonAccessor struct {
	_data               []TrainPartyConstValueCommon
	_dataConstValueName map[string]TrainPartyConstValueCommon
}

// LoadData retrieves the data. Must be called before TrainPartyConstValueCommon.GroupData
func (a *TrainPartyConstValueCommonAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TrainPartyConstValueCommon.json")
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
func (a *TrainPartyConstValueCommonAccessor) Raw() ([]TrainPartyConstValueCommon, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TrainPartyConstValueCommon{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with TrainPartyConstValueCommonAccessor.LoadData to preload everything
func (a *TrainPartyConstValueCommonAccessor) GroupData() {
	for _, d := range a._data {
		a._dataConstValueName[d.ConstValueName] = d
	}
}

// ByConstValueName returns the TrainPartyConstValueCommon uniquely identified by ConstValueName
//
// Error is only non-nil if the source errors out
func (a *TrainPartyConstValueCommonAccessor) ByConstValueName(identifier string) (TrainPartyConstValueCommon, error) {
	if a._dataConstValueName == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TrainPartyConstValueCommon{}, err
			}
		}
		a.GroupData()
	}
	return a._dataConstValueName[identifier], nil
}
