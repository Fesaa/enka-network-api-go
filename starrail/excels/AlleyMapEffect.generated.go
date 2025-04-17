package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type AlleyMapEffect struct {
	BuffOrDebuff     string    `json:"BuffOrDebuff"`
	MapEffectID      float64   `json:"MapEffectID"`
	MapEffectSubType string    `json:"MapEffectSubType"`
	MapEffectTitle   hash.Hash `json:"MapEffectTitle"`
	Param1           float64   `json:"Param1"`
	Param2           float64   `json:"Param2"`
}
type AlleyMapEffectAccessor struct {
	_data            []AlleyMapEffect
	_dataMapEffectID map[float64]AlleyMapEffect
}

// LoadData retrieves the data. Must be called before AlleyMapEffect.GroupData
func (a *AlleyMapEffectAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AlleyMapEffect.json")
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
func (a *AlleyMapEffectAccessor) Raw() ([]AlleyMapEffect, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AlleyMapEffect{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AlleyMapEffectAccessor.LoadData to preload everything
func (a *AlleyMapEffectAccessor) GroupData() {
	for _, d := range a._data {
		a._dataMapEffectID[d.MapEffectID] = d
	}
}

// ByMapEffectID returns the AlleyMapEffect uniquely identified by MapEffectID
//
// Error is only non-nil if the source errors out
func (a *AlleyMapEffectAccessor) ByMapEffectID(identifier float64) (AlleyMapEffect, error) {
	if a._dataMapEffectID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AlleyMapEffect{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMapEffectID[identifier], nil
}
