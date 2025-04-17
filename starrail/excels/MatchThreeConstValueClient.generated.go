package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type MatchThreeConstValueClient struct {
	ConstValueName string                          `json:"ConstValueName"`
	Value          MatchThreeConstValueClientValue `json:"Value"`
}
type MatchThreeConstValueClientValue struct {
	ArrayValue  []map[string]json.Number `json:"ArrayValue"`
	IntValue    json.Number              `json:"IntValue"`
	StringValue string                   `json:"StringValue"`
}
type MatchThreeConstValueClientAccessor struct {
	_data               []MatchThreeConstValueClient
	_dataConstValueName map[string]MatchThreeConstValueClient
}

// LoadData retrieves the data. Must be called before MatchThreeConstValueClient.GroupData
func (a *MatchThreeConstValueClientAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MatchThreeConstValueClient.json")
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
func (a *MatchThreeConstValueClientAccessor) Raw() ([]MatchThreeConstValueClient, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MatchThreeConstValueClient{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MatchThreeConstValueClientAccessor.LoadData to preload everything
func (a *MatchThreeConstValueClientAccessor) GroupData() {
	a._dataConstValueName = map[string]MatchThreeConstValueClient{}
	for _, d := range a._data {
		a._dataConstValueName[d.ConstValueName] = d
	}
}

// ByConstValueName returns the MatchThreeConstValueClient uniquely identified by ConstValueName
//
// Error is only non-nil if the source errors out
func (a *MatchThreeConstValueClientAccessor) ByConstValueName(identifier string) (MatchThreeConstValueClient, error) {
	if a._dataConstValueName == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MatchThreeConstValueClient{}, err
			}
		}
		a.GroupData()
	}
	return a._dataConstValueName[identifier], nil
}
