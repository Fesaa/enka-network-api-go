package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type RogueMagicLayerEffect struct {
	DescParamList   []hash.IntValue `json:"DescParamList"`
	LayerEffectDesc hash.Hash       `json:"LayerEffectDesc"`
	LayerEffectID   float64         `json:"LayerEffectID"`
	LayerEffectName hash.Hash       `json:"LayerEffectName"`
}
type RogueMagicLayerEffectAccessor struct {
	_data              []RogueMagicLayerEffect
	_dataLayerEffectID map[float64]RogueMagicLayerEffect
}

// LoadData retrieves the data. Must be called before RogueMagicLayerEffect.GroupData
func (a *RogueMagicLayerEffectAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueMagicLayerEffect.json")
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
func (a *RogueMagicLayerEffectAccessor) Raw() ([]RogueMagicLayerEffect, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueMagicLayerEffect{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueMagicLayerEffectAccessor.LoadData to preload everything
func (a *RogueMagicLayerEffectAccessor) GroupData() {
	a._dataLayerEffectID = map[float64]RogueMagicLayerEffect{}
	for _, d := range a._data {
		a._dataLayerEffectID[d.LayerEffectID] = d
	}
}

// ByLayerEffectID returns the RogueMagicLayerEffect uniquely identified by LayerEffectID
//
// Error is only non-nil if the source errors out
func (a *RogueMagicLayerEffectAccessor) ByLayerEffectID(identifier float64) (RogueMagicLayerEffect, error) {
	if a._dataLayerEffectID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueMagicLayerEffect{}, err
			}
		}
		a.GroupData()
	}
	return a._dataLayerEffectID[identifier], nil
}
