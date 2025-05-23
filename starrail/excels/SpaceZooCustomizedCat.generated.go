package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type SpaceZooCustomizedCat struct {
	AddCatID       float64   `json:"AddCatID"`
	ChannelFeature []float64 `json:"ChannelFeature"`
	NotShowDialog  bool      `json:"NotShowDialog"`
}
type SpaceZooCustomizedCatAccessor struct {
	_data         []SpaceZooCustomizedCat
	_dataAddCatID map[float64]SpaceZooCustomizedCat
}

// LoadData retrieves the data. Must be called before SpaceZooCustomizedCat.GroupData
func (a *SpaceZooCustomizedCatAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/SpaceZooCustomizedCat.json")
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
func (a *SpaceZooCustomizedCatAccessor) Raw() ([]SpaceZooCustomizedCat, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []SpaceZooCustomizedCat{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with SpaceZooCustomizedCatAccessor.LoadData to preload everything
func (a *SpaceZooCustomizedCatAccessor) GroupData() {
	a._dataAddCatID = map[float64]SpaceZooCustomizedCat{}
	for _, d := range a._data {
		a._dataAddCatID[d.AddCatID] = d
	}
}

// ByAddCatID returns the SpaceZooCustomizedCat uniquely identified by AddCatID
//
// Error is only non-nil if the source errors out
func (a *SpaceZooCustomizedCatAccessor) ByAddCatID(identifier float64) (SpaceZooCustomizedCat, error) {
	if a._dataAddCatID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return SpaceZooCustomizedCat{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAddCatID[identifier], nil
}
