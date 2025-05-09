package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type BattlePassConstValue struct {
	ConstValueName string                    `json:"ConstValueName"`
	Value          BattlePassConstValueValue `json:"Value"`
}
type BattlePassConstValueValue struct {
	ArrayValue []hash.IntValue `json:"ArrayValue"`
	FloatValue float64         `json:"FloatValue"`
	IntValue   float64         `json:"IntValue"`
}
type BattlePassConstValueAccessor struct {
	_data               []BattlePassConstValue
	_dataConstValueName map[string]BattlePassConstValue
}

// LoadData retrieves the data. Must be called before BattlePassConstValue.GroupData
func (a *BattlePassConstValueAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/BattlePassConstValue.json")
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
func (a *BattlePassConstValueAccessor) Raw() ([]BattlePassConstValue, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []BattlePassConstValue{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with BattlePassConstValueAccessor.LoadData to preload everything
func (a *BattlePassConstValueAccessor) GroupData() {
	a._dataConstValueName = map[string]BattlePassConstValue{}
	for _, d := range a._data {
		a._dataConstValueName[d.ConstValueName] = d
	}
}

// ByConstValueName returns the BattlePassConstValue uniquely identified by ConstValueName
//
// Error is only non-nil if the source errors out
func (a *BattlePassConstValueAccessor) ByConstValueName(identifier string) (BattlePassConstValue, error) {
	if a._dataConstValueName == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return BattlePassConstValue{}, err
			}
		}
		a.GroupData()
	}
	return a._dataConstValueName[identifier], nil
}
