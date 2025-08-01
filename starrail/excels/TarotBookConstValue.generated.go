package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type TarotBookConstValue struct {
	ConstValueName string        `json:"ConstValueName"`
	Value          hash.IntValue `json:"Value"`
}
type TarotBookConstValueAccessor struct {
	_data               []TarotBookConstValue
	_dataConstValueName map[string]TarotBookConstValue
}

// LoadData retrieves the data. Must be called before TarotBookConstValue.GroupData
func (a *TarotBookConstValueAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TarotBookConstValue.json")
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
func (a *TarotBookConstValueAccessor) Raw() ([]TarotBookConstValue, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TarotBookConstValue{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with TarotBookConstValueAccessor.LoadData to preload everything
func (a *TarotBookConstValueAccessor) GroupData() {
	a._dataConstValueName = map[string]TarotBookConstValue{}
	for _, d := range a._data {
		a._dataConstValueName[d.ConstValueName] = d
	}
}

// ByConstValueName returns the TarotBookConstValue uniquely identified by ConstValueName
//
// Error is only non-nil if the source errors out
func (a *TarotBookConstValueAccessor) ByConstValueName(identifier string) (TarotBookConstValue, error) {
	if a._dataConstValueName == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TarotBookConstValue{}, err
			}
		}
		a.GroupData()
	}
	return a._dataConstValueName[identifier], nil
}
