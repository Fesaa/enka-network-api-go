package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type RogueTournExhibitionConfig struct {
	ConfigID   float64 `json:"ConfigID"`
	GroupID    float64 `json:"GroupID"`
	PaintingID float64 `json:"PaintingID"`
	Type       string  `json:"Type"`
}
type RogueTournExhibitionConfigAccessor struct {
	_data           []RogueTournExhibitionConfig
	_dataPaintingID map[float64]RogueTournExhibitionConfig
}

// LoadData retrieves the data. Must be called before RogueTournExhibitionConfig.GroupData
func (a *RogueTournExhibitionConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueTournExhibitionConfig.json")
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
func (a *RogueTournExhibitionConfigAccessor) Raw() ([]RogueTournExhibitionConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueTournExhibitionConfig{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueTournExhibitionConfigAccessor.LoadData to preload everything
func (a *RogueTournExhibitionConfigAccessor) GroupData() {
	for _, d := range a._data {
		a._dataPaintingID[d.PaintingID] = d
	}
}

// ByPaintingID returns the RogueTournExhibitionConfig uniquely identified by PaintingID
//
// Error is only non-nil if the source errors out
func (a *RogueTournExhibitionConfigAccessor) ByPaintingID(identifier float64) (RogueTournExhibitionConfig, error) {
	if a._dataPaintingID == nil {
		err := a.LoadData()
		if err != nil {
			return RogueTournExhibitionConfig{}, err
		}
		a.GroupData()
	}
	return a._dataPaintingID[identifier], nil
}
