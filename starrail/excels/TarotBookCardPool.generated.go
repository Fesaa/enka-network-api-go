package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type TarotBookCardPool struct {
	CardList  []float64 `json:"CardList"`
	ClueList  []float64 `json:"ClueList"`
	ID        float64   `json:"ID"`
	StoryList []float64 `json:"StoryList"`
}
type TarotBookCardPoolAccessor struct {
	_data   []TarotBookCardPool
	_dataID map[float64]TarotBookCardPool
}

// LoadData retrieves the data. Must be called before TarotBookCardPool.GroupData
func (a *TarotBookCardPoolAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TarotBookCardPool.json")
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
func (a *TarotBookCardPoolAccessor) Raw() ([]TarotBookCardPool, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TarotBookCardPool{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with TarotBookCardPoolAccessor.LoadData to preload everything
func (a *TarotBookCardPoolAccessor) GroupData() {
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the TarotBookCardPool uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *TarotBookCardPoolAccessor) ByID(identifier float64) (TarotBookCardPool, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TarotBookCardPool{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
