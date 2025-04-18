package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type RelicSetBonusValue struct {
	BonusValue float64    `json:"BonusValue"`
	Property   string     `json:"Property"`
	SetID      float64    `json:"SetID"`
	Threshold  hash.Value `json:"Threshold"`
}
type RelicSetBonusValueAccessor struct {
	_data      []RelicSetBonusValue
	_dataSetID map[float64]RelicSetBonusValue
}

// LoadData retrieves the data. Must be called before RelicSetBonusValue.GroupData
func (a *RelicSetBonusValueAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RelicSetBonusValue.json")
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
func (a *RelicSetBonusValueAccessor) Raw() ([]RelicSetBonusValue, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RelicSetBonusValue{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RelicSetBonusValueAccessor.LoadData to preload everything
func (a *RelicSetBonusValueAccessor) GroupData() {
	a._dataSetID = map[float64]RelicSetBonusValue{}
	for _, d := range a._data {
		a._dataSetID[d.SetID] = d
	}
}

// BySetID returns the RelicSetBonusValue uniquely identified by SetID
//
// Error is only non-nil if the source errors out
func (a *RelicSetBonusValueAccessor) BySetID(identifier float64) (RelicSetBonusValue, error) {
	if a._dataSetID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RelicSetBonusValue{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSetID[identifier], nil
}
