package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type BookSeriesConfig struct {
	BookSeries         hash.Hash `json:"BookSeries"`
	BookSeriesComments hash.Hash `json:"BookSeriesComments"`
	BookSeriesID       float64   `json:"BookSeriesID"`
	BookSeriesNum      float64   `json:"BookSeriesNum"`
	BookSeriesWorld    float64   `json:"BookSeriesWorld"`
	IsShowInBookshelf  bool      `json:"IsShowInBookshelf"`
}
type BookSeriesConfigAccessor struct {
	_data             []BookSeriesConfig
	_dataBookSeriesID map[float64]BookSeriesConfig
}

// LoadData retrieves the data. Must be called before BookSeriesConfig.GroupData
func (a *BookSeriesConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/BookSeriesConfig.json")
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
func (a *BookSeriesConfigAccessor) Raw() ([]BookSeriesConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []BookSeriesConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with BookSeriesConfigAccessor.LoadData to preload everything
func (a *BookSeriesConfigAccessor) GroupData() {
	a._dataBookSeriesID = map[float64]BookSeriesConfig{}
	for _, d := range a._data {
		a._dataBookSeriesID[d.BookSeriesID] = d
	}
}

// ByBookSeriesID returns the BookSeriesConfig uniquely identified by BookSeriesID
//
// Error is only non-nil if the source errors out
func (a *BookSeriesConfigAccessor) ByBookSeriesID(identifier float64) (BookSeriesConfig, error) {
	if a._dataBookSeriesID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return BookSeriesConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBookSeriesID[identifier], nil
}
