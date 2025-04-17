package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type ActivityExpeditionConst struct {
	ConstValueName string        `json:"ConstValueName"`
	Value          hash.IntValue `json:"Value"`
}
type ActivityExpeditionConstAccessor struct {
	_data               []ActivityExpeditionConst
	_dataConstValueName map[string]ActivityExpeditionConst
}

// LoadData retrieves the data. Must be called before ActivityExpeditionConst.GroupData
func (a *ActivityExpeditionConstAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActivityExpeditionConst.json")
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
func (a *ActivityExpeditionConstAccessor) Raw() ([]ActivityExpeditionConst, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActivityExpeditionConst{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ActivityExpeditionConstAccessor.LoadData to preload everything
func (a *ActivityExpeditionConstAccessor) GroupData() {
	for _, d := range a._data {
		a._dataConstValueName[d.ConstValueName] = d
	}
}

// ByConstValueName returns the ActivityExpeditionConst uniquely identified by ConstValueName
//
// Error is only non-nil if the source errors out
func (a *ActivityExpeditionConstAccessor) ByConstValueName(identifier string) (ActivityExpeditionConst, error) {
	if a._dataConstValueName == nil {
		err := a.LoadData()
		if err != nil {
			return ActivityExpeditionConst{}, err
		}
		a.GroupData()
	}
	return a._dataConstValueName[identifier], nil
}
