package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type ConstValueChallengeCommon struct {
	ConstValueName string        `json:"ConstValueName"`
	Value          hash.IntValue `json:"Value"`
}
type ConstValueChallengeCommonAccessor struct {
	_data               []ConstValueChallengeCommon
	_dataConstValueName map[string]ConstValueChallengeCommon
}

// LoadData retrieves the data. Must be called before ConstValueChallengeCommon.GroupData
func (a *ConstValueChallengeCommonAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ConstValueChallengeCommon.json")
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
func (a *ConstValueChallengeCommonAccessor) Raw() ([]ConstValueChallengeCommon, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ConstValueChallengeCommon{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ConstValueChallengeCommonAccessor.LoadData to preload everything
func (a *ConstValueChallengeCommonAccessor) GroupData() {
	for _, d := range a._data {
		a._dataConstValueName[d.ConstValueName] = d
	}
}

// ByConstValueName returns the ConstValueChallengeCommon uniquely identified by ConstValueName
//
// Error is only non-nil if the source errors out
func (a *ConstValueChallengeCommonAccessor) ByConstValueName(identifier string) (ConstValueChallengeCommon, error) {
	if a._dataConstValueName == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ConstValueChallengeCommon{}, err
			}
		}
		a.GroupData()
	}
	return a._dataConstValueName[identifier], nil
}
