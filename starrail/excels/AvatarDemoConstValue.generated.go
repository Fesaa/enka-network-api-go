package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type AvatarDemoConstValue struct {
	ConstValueName string           `json:"ConstValueName"`
	Value          hash.StringValue `json:"Value"`
}
type AvatarDemoConstValueAccessor struct {
	_data               []AvatarDemoConstValue
	_dataConstValueName map[string]AvatarDemoConstValue
}

// LoadData retrieves the data. Must be called before AvatarDemoConstValue.GroupData
func (a *AvatarDemoConstValueAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AvatarDemoConstValue.json")
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
func (a *AvatarDemoConstValueAccessor) Raw() ([]AvatarDemoConstValue, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AvatarDemoConstValue{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AvatarDemoConstValueAccessor.LoadData to preload everything
func (a *AvatarDemoConstValueAccessor) GroupData() {
	a._dataConstValueName = map[string]AvatarDemoConstValue{}
	for _, d := range a._data {
		a._dataConstValueName[d.ConstValueName] = d
	}
}

// ByConstValueName returns the AvatarDemoConstValue uniquely identified by ConstValueName
//
// Error is only non-nil if the source errors out
func (a *AvatarDemoConstValueAccessor) ByConstValueName(identifier string) (AvatarDemoConstValue, error) {
	if a._dataConstValueName == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarDemoConstValue{}, err
			}
		}
		a.GroupData()
	}
	return a._dataConstValueName[identifier], nil
}
