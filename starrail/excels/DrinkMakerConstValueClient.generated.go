package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type DrinkMakerConstValueClient struct {
	ConstValueName string                          `json:"ConstValueName"`
	Value          DrinkMakerConstValueClientValue `json:"Value"`
}
type DrinkMakerConstValueClientValue struct {
	ArrayValue  []hash.IntValue `json:"ArrayValue"`
	FloatValue  float64         `json:"FloatValue"`
	IntValue    float64         `json:"IntValue"`
	StringValue string          `json:"StringValue"`
}
type DrinkMakerConstValueClientAccessor struct {
	_data               []DrinkMakerConstValueClient
	_dataConstValueName map[string]DrinkMakerConstValueClient
}

// LoadData retrieves the data. Must be called before DrinkMakerConstValueClient.GroupData
func (a *DrinkMakerConstValueClientAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/DrinkMakerConstValueClient.json")
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
func (a *DrinkMakerConstValueClientAccessor) Raw() ([]DrinkMakerConstValueClient, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []DrinkMakerConstValueClient{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with DrinkMakerConstValueClientAccessor.LoadData to preload everything
func (a *DrinkMakerConstValueClientAccessor) GroupData() {
	a._dataConstValueName = map[string]DrinkMakerConstValueClient{}
	for _, d := range a._data {
		a._dataConstValueName[d.ConstValueName] = d
	}
}

// ByConstValueName returns the DrinkMakerConstValueClient uniquely identified by ConstValueName
//
// Error is only non-nil if the source errors out
func (a *DrinkMakerConstValueClientAccessor) ByConstValueName(identifier string) (DrinkMakerConstValueClient, error) {
	if a._dataConstValueName == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return DrinkMakerConstValueClient{}, err
			}
		}
		a.GroupData()
	}
	return a._dataConstValueName[identifier], nil
}
