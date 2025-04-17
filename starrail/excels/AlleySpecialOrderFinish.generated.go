package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type AlleySpecialOrderFinish struct {
	Param1                 float64   `json:"Param1"`
	Param2                 float64   `json:"Param2"`
	SpecialOrderFinishDesc hash.Hash `json:"SpecialOrderFinishDesc"`
	SpecialOrderFinishID   float64   `json:"SpecialOrderFinishID"`
	SpecialOrderFinishType string    `json:"SpecialOrderFinishType"`
}
type AlleySpecialOrderFinishAccessor struct {
	_data                     []AlleySpecialOrderFinish
	_dataSpecialOrderFinishID map[float64]AlleySpecialOrderFinish
}

// LoadData retrieves the data. Must be called before AlleySpecialOrderFinish.GroupData
func (a *AlleySpecialOrderFinishAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AlleySpecialOrderFinish.json")
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
func (a *AlleySpecialOrderFinishAccessor) Raw() ([]AlleySpecialOrderFinish, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AlleySpecialOrderFinish{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AlleySpecialOrderFinishAccessor.LoadData to preload everything
func (a *AlleySpecialOrderFinishAccessor) GroupData() {
	for _, d := range a._data {
		a._dataSpecialOrderFinishID[d.SpecialOrderFinishID] = d
	}
}

// BySpecialOrderFinishID returns the AlleySpecialOrderFinish uniquely identified by SpecialOrderFinishID
//
// Error is only non-nil if the source errors out
func (a *AlleySpecialOrderFinishAccessor) BySpecialOrderFinishID(identifier float64) (AlleySpecialOrderFinish, error) {
	if a._dataSpecialOrderFinishID == nil {
		err := a.LoadData()
		if err != nil {
			return AlleySpecialOrderFinish{}, err
		}
		a.GroupData()
	}
	return a._dataSpecialOrderFinishID[identifier], nil
}
