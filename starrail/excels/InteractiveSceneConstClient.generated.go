package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type InteractiveSceneConstClient struct {
	ConstValueName string           `json:"ConstValueName"`
	Value          hash.StringValue `json:"Value"`
}
type InteractiveSceneConstClientAccessor struct {
	_data               []InteractiveSceneConstClient
	_dataConstValueName map[string]InteractiveSceneConstClient
}

// LoadData retrieves the data. Must be called before InteractiveSceneConstClient.GroupData
func (a *InteractiveSceneConstClientAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/InteractiveSceneConstClient.json")
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
func (a *InteractiveSceneConstClientAccessor) Raw() ([]InteractiveSceneConstClient, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []InteractiveSceneConstClient{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with InteractiveSceneConstClientAccessor.LoadData to preload everything
func (a *InteractiveSceneConstClientAccessor) GroupData() {
	a._dataConstValueName = map[string]InteractiveSceneConstClient{}
	for _, d := range a._data {
		a._dataConstValueName[d.ConstValueName] = d
	}
}

// ByConstValueName returns the InteractiveSceneConstClient uniquely identified by ConstValueName
//
// Error is only non-nil if the source errors out
func (a *InteractiveSceneConstClientAccessor) ByConstValueName(identifier string) (InteractiveSceneConstClient, error) {
	if a._dataConstValueName == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return InteractiveSceneConstClient{}, err
			}
		}
		a.GroupData()
	}
	return a._dataConstValueName[identifier], nil
}
