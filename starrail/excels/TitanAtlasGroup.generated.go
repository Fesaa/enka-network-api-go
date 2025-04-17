package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type TitanAtlasGroup struct {
	TitanGroupDesc hash.Hash `json:"TitanGroupDesc"`
	TitanGroupID   float64   `json:"TitanGroupID"`
	TitanGroupName hash.Hash `json:"TitanGroupName"`
	TitleBGColor   string    `json:"TitleBGColor"`
}
type TitanAtlasGroupAccessor struct {
	_data             []TitanAtlasGroup
	_dataTitanGroupID map[float64]TitanAtlasGroup
	_dataTitleBGColor map[string]TitanAtlasGroup
}

// LoadData retrieves the data. Must be called before TitanAtlasGroup.GroupData
func (a *TitanAtlasGroupAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TitanAtlasGroup.json")
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
func (a *TitanAtlasGroupAccessor) Raw() ([]TitanAtlasGroup, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TitanAtlasGroup{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with TitanAtlasGroupAccessor.LoadData to preload everything
func (a *TitanAtlasGroupAccessor) GroupData() {
	for _, d := range a._data {
		a._dataTitanGroupID[d.TitanGroupID] = d
		a._dataTitleBGColor[d.TitleBGColor] = d
	}
}

// ByTitanGroupID returns the TitanAtlasGroup uniquely identified by TitanGroupID
//
// Error is only non-nil if the source errors out
func (a *TitanAtlasGroupAccessor) ByTitanGroupID(identifier float64) (TitanAtlasGroup, error) {
	if a._dataTitanGroupID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TitanAtlasGroup{}, err
			}
		}
		a.GroupData()
	}
	return a._dataTitanGroupID[identifier], nil
}

// ByTitleBGColor returns the TitanAtlasGroup uniquely identified by TitleBGColor
//
// Error is only non-nil if the source errors out
func (a *TitanAtlasGroupAccessor) ByTitleBGColor(identifier string) (TitanAtlasGroup, error) {
	if a._dataTitleBGColor == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TitanAtlasGroup{}, err
			}
		}
		a.GroupData()
	}
	return a._dataTitleBGColor[identifier], nil
}
