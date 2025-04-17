package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type RogueTournConstClient struct {
	ConstValueName string           `json:"ConstValueName"`
	Value          hash.StringValue `json:"Value"`
}
type RogueTournConstClientAccessor struct {
	_data               []RogueTournConstClient
	_dataConstValueName map[string]RogueTournConstClient
}

// LoadData retrieves the data. Must be called before RogueTournConstClient.GroupData
func (a *RogueTournConstClientAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueTournConstClient.json")
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
func (a *RogueTournConstClientAccessor) Raw() ([]RogueTournConstClient, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueTournConstClient{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueTournConstClientAccessor.LoadData to preload everything
func (a *RogueTournConstClientAccessor) GroupData() {
	for _, d := range a._data {
		a._dataConstValueName[d.ConstValueName] = d
	}
}

// ByConstValueName returns the RogueTournConstClient uniquely identified by ConstValueName
//
// Error is only non-nil if the source errors out
func (a *RogueTournConstClientAccessor) ByConstValueName(identifier string) (RogueTournConstClient, error) {
	if a._dataConstValueName == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueTournConstClient{}, err
			}
		}
		a.GroupData()
	}
	return a._dataConstValueName[identifier], nil
}
