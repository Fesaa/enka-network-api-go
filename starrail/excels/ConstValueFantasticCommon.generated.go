package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type ConstValueFantasticCommon struct {
	ConstValueName string        `json:"ConstValueName"`
	Value          hash.IntValue `json:"Value"`
}
type ConstValueFantasticCommonAccessor struct {
	_data               []ConstValueFantasticCommon
	_dataConstValueName map[string]ConstValueFantasticCommon
}

// LoadData retrieves the data. Must be called before ConstValueFantasticCommon.GroupData
func (a *ConstValueFantasticCommonAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ConstValueFantasticCommon.json")
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
func (a *ConstValueFantasticCommonAccessor) Raw() ([]ConstValueFantasticCommon, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ConstValueFantasticCommon{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ConstValueFantasticCommonAccessor.LoadData to preload everything
func (a *ConstValueFantasticCommonAccessor) GroupData() {
	for _, d := range a._data {
		a._dataConstValueName[d.ConstValueName] = d
	}
}

// ByConstValueName returns the ConstValueFantasticCommon uniquely identified by ConstValueName
//
// Error is only non-nil if the source errors out
func (a *ConstValueFantasticCommonAccessor) ByConstValueName(identifier string) (ConstValueFantasticCommon, error) {
	if a._dataConstValueName == nil {
		err := a.LoadData()
		if err != nil {
			return ConstValueFantasticCommon{}, err
		}
		a.GroupData()
	}
	return a._dataConstValueName[identifier], nil
}
