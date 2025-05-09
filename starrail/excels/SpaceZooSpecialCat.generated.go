package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type SpaceZooSpecialCat struct {
	ColorBar              []string  `json:"ColorBar"`
	ImagePath             string    `json:"ImagePath"`
	IsHide                bool      `json:"IsHide"`
	LargeImagePath        string    `json:"LargeImagePath"`
	MatPath               string    `json:"MatPath"`
	MatchedChannelFeature []float64 `json:"MatchedChannelFeature"`
	Name                  hash.Hash `json:"Name"`
	PhotoSubmissionID     float64   `json:"PhotoSubmissionID"`
	ResearchPointSSR      float64   `json:"ResearchPointSSR"`
	SpecialCatID          float64   `json:"SpecialCatID"`
	SpecialItem           float64   `json:"SpecialItem"`
	TipsCustomizedCat     []float64 `json:"TipsCustomizedCat"`
	TipsMissionID         float64   `json:"TipsMissionID"`
}
type SpaceZooSpecialCatAccessor struct {
	_data               []SpaceZooSpecialCat
	_dataImagePath      map[string]SpaceZooSpecialCat
	_dataLargeImagePath map[string]SpaceZooSpecialCat
	_dataMatPath        map[string]SpaceZooSpecialCat
	_dataSpecialCatID   map[float64]SpaceZooSpecialCat
	_dataSpecialItem    map[float64]SpaceZooSpecialCat
}

// LoadData retrieves the data. Must be called before SpaceZooSpecialCat.GroupData
func (a *SpaceZooSpecialCatAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/SpaceZooSpecialCat.json")
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
func (a *SpaceZooSpecialCatAccessor) Raw() ([]SpaceZooSpecialCat, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []SpaceZooSpecialCat{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with SpaceZooSpecialCatAccessor.LoadData to preload everything
func (a *SpaceZooSpecialCatAccessor) GroupData() {
	a._dataImagePath = map[string]SpaceZooSpecialCat{}
	a._dataLargeImagePath = map[string]SpaceZooSpecialCat{}
	a._dataMatPath = map[string]SpaceZooSpecialCat{}
	a._dataSpecialCatID = map[float64]SpaceZooSpecialCat{}
	a._dataSpecialItem = map[float64]SpaceZooSpecialCat{}
	for _, d := range a._data {
		a._dataImagePath[d.ImagePath] = d
		a._dataLargeImagePath[d.LargeImagePath] = d
		a._dataMatPath[d.MatPath] = d
		a._dataSpecialCatID[d.SpecialCatID] = d
		a._dataSpecialItem[d.SpecialItem] = d
	}
}

// ByImagePath returns the SpaceZooSpecialCat uniquely identified by ImagePath
//
// Error is only non-nil if the source errors out
func (a *SpaceZooSpecialCatAccessor) ByImagePath(identifier string) (SpaceZooSpecialCat, error) {
	if a._dataImagePath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return SpaceZooSpecialCat{}, err
			}
		}
		a.GroupData()
	}
	return a._dataImagePath[identifier], nil
}

// ByLargeImagePath returns the SpaceZooSpecialCat uniquely identified by LargeImagePath
//
// Error is only non-nil if the source errors out
func (a *SpaceZooSpecialCatAccessor) ByLargeImagePath(identifier string) (SpaceZooSpecialCat, error) {
	if a._dataLargeImagePath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return SpaceZooSpecialCat{}, err
			}
		}
		a.GroupData()
	}
	return a._dataLargeImagePath[identifier], nil
}

// ByMatPath returns the SpaceZooSpecialCat uniquely identified by MatPath
//
// Error is only non-nil if the source errors out
func (a *SpaceZooSpecialCatAccessor) ByMatPath(identifier string) (SpaceZooSpecialCat, error) {
	if a._dataMatPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return SpaceZooSpecialCat{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMatPath[identifier], nil
}

// BySpecialCatID returns the SpaceZooSpecialCat uniquely identified by SpecialCatID
//
// Error is only non-nil if the source errors out
func (a *SpaceZooSpecialCatAccessor) BySpecialCatID(identifier float64) (SpaceZooSpecialCat, error) {
	if a._dataSpecialCatID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return SpaceZooSpecialCat{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSpecialCatID[identifier], nil
}

// BySpecialItem returns the SpaceZooSpecialCat uniquely identified by SpecialItem
//
// Error is only non-nil if the source errors out
func (a *SpaceZooSpecialCatAccessor) BySpecialItem(identifier float64) (SpaceZooSpecialCat, error) {
	if a._dataSpecialItem == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return SpaceZooSpecialCat{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSpecialItem[identifier], nil
}
