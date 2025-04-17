package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type NounAtlas struct {
	ID           float64   `json:"ID"`
	IsIntroPage  bool      `json:"IsIntroPage"`
	NounDesc     hash.Hash `json:"NounDesc"`
	NounTitle    hash.Hash `json:"NounTitle"`
	RelatedTerms []float64 `json:"RelatedTerms"`
	SortID       float64   `json:"SortID"`
	Type         float64   `json:"Type"`
	Unlock       float64   `json:"Unlock"`
}
type NounAtlasAccessor struct {
	_data []NounAtlas
}

// LoadData retrieves the data. Must be called before NounAtlas.GroupData
func (a *NounAtlasAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/NounAtlas.json")
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
func (a *NounAtlasAccessor) Raw() ([]NounAtlas, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []NounAtlas{}, err
		}
	}
	return a._data, nil
}
