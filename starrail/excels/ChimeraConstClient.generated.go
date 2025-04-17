package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type ChimeraConstClient struct {
	ConstValueName string                  `json:"ConstValueName"`
	Value          ChimeraConstClientValue `json:"Value"`
}
type ChimeraConstClientValue struct {
	ArrayValue  []hash.IntValue             `json:"ArrayValue"`
	FloatValue  float64                     `json:"FloatValue"`
	IntValue    float64                     `json:"IntValue"`
	MapValue    map[string]hash.StringValue `json:"MapValue"`
	StringValue string                      `json:"StringValue"`
}
type ChimeraConstClientAccessor struct {
	_data               []ChimeraConstClient
	_dataConstValueName map[string]ChimeraConstClient
}

// LoadData retrieves the data. Must be called before ChimeraConstClient.GroupData
func (a *ChimeraConstClientAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ChimeraConstClient.json")
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
func (a *ChimeraConstClientAccessor) Raw() ([]ChimeraConstClient, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ChimeraConstClient{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ChimeraConstClientAccessor.LoadData to preload everything
func (a *ChimeraConstClientAccessor) GroupData() {
	for _, d := range a._data {
		a._dataConstValueName[d.ConstValueName] = d
	}
}

// ByConstValueName returns the ChimeraConstClient uniquely identified by ConstValueName
//
// Error is only non-nil if the source errors out
func (a *ChimeraConstClientAccessor) ByConstValueName(identifier string) (ChimeraConstClient, error) {
	if a._dataConstValueName == nil {
		err := a.LoadData()
		if err != nil {
			return ChimeraConstClient{}, err
		}
		a.GroupData()
	}
	return a._dataConstValueName[identifier], nil
}
