package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type ConstValueChallengeClient struct {
	ConstValueName string        `json:"ConstValueName"`
	Value          hash.IntValue `json:"Value"`
}
type ConstValueChallengeClientAccessor struct {
	_data               []ConstValueChallengeClient
	_dataConstValueName map[string]ConstValueChallengeClient
}

// LoadData retrieves the data. Must be called before ConstValueChallengeClient.GroupData
func (a *ConstValueChallengeClientAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ConstValueChallengeClient.json")
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
func (a *ConstValueChallengeClientAccessor) Raw() ([]ConstValueChallengeClient, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ConstValueChallengeClient{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ConstValueChallengeClientAccessor.LoadData to preload everything
func (a *ConstValueChallengeClientAccessor) GroupData() {
	for _, d := range a._data {
		a._dataConstValueName[d.ConstValueName] = d
	}
}

// ByConstValueName returns the ConstValueChallengeClient uniquely identified by ConstValueName
//
// Error is only non-nil if the source errors out
func (a *ConstValueChallengeClientAccessor) ByConstValueName(identifier string) (ConstValueChallengeClient, error) {
	if a._dataConstValueName == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ConstValueChallengeClient{}, err
			}
		}
		a.GroupData()
	}
	return a._dataConstValueName[identifier], nil
}
