package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type SpaceZooHandbookText struct {
	SpecialCatID float64 `json:"SpecialCatID"`
	UITextID     string  `json:"UITextID"`
}
type SpaceZooHandbookTextAccessor struct {
	_data             []SpaceZooHandbookText
	_dataSpecialCatID map[float64]SpaceZooHandbookText
}

// LoadData retrieves the data. Must be called before SpaceZooHandbookText.GroupData
func (a *SpaceZooHandbookTextAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/SpaceZooHandbookText.json")
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
func (a *SpaceZooHandbookTextAccessor) Raw() ([]SpaceZooHandbookText, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []SpaceZooHandbookText{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with SpaceZooHandbookTextAccessor.LoadData to preload everything
func (a *SpaceZooHandbookTextAccessor) GroupData() {
	a._dataSpecialCatID = map[float64]SpaceZooHandbookText{}
	for _, d := range a._data {
		a._dataSpecialCatID[d.SpecialCatID] = d
	}
}

// BySpecialCatID returns the SpaceZooHandbookText uniquely identified by SpecialCatID
//
// Error is only non-nil if the source errors out
func (a *SpaceZooHandbookTextAccessor) BySpecialCatID(identifier float64) (SpaceZooHandbookText, error) {
	if a._dataSpecialCatID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return SpaceZooHandbookText{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSpecialCatID[identifier], nil
}
