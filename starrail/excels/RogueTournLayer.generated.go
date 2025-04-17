package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type RogueTournLayer struct {
	LayerID    float64 `json:"LayerID"`
	LayerNumID float64 `json:"LayerNumID"`
}
type RogueTournLayerAccessor struct {
	_data        []RogueTournLayer
	_dataLayerID map[float64]RogueTournLayer
}

// LoadData retrieves the data. Must be called before RogueTournLayer.GroupData
func (a *RogueTournLayerAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueTournLayer.json")
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
func (a *RogueTournLayerAccessor) Raw() ([]RogueTournLayer, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueTournLayer{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueTournLayerAccessor.LoadData to preload everything
func (a *RogueTournLayerAccessor) GroupData() {
	for _, d := range a._data {
		a._dataLayerID[d.LayerID] = d
	}
}

// ByLayerID returns the RogueTournLayer uniquely identified by LayerID
//
// Error is only non-nil if the source errors out
func (a *RogueTournLayerAccessor) ByLayerID(identifier float64) (RogueTournLayer, error) {
	if a._dataLayerID == nil {
		err := a.LoadData()
		if err != nil {
			return RogueTournLayer{}, err
		}
		a.GroupData()
	}
	return a._dataLayerID[identifier], nil
}
