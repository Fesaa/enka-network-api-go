package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type ChimeraGalleryAct struct {
	ActID float64   `json:"ActID"`
	Icon  string    `json:"Icon"`
	Name  hash.Hash `json:"Name"`
	Sort  float64   `json:"Sort"`
}
type ChimeraGalleryActAccessor struct {
	_data      []ChimeraGalleryAct
	_dataActID map[float64]ChimeraGalleryAct
	_dataIcon  map[string]ChimeraGalleryAct
	_dataSort  map[float64]ChimeraGalleryAct
}

// LoadData retrieves the data. Must be called before ChimeraGalleryAct.GroupData
func (a *ChimeraGalleryActAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ChimeraGalleryAct.json")
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
func (a *ChimeraGalleryActAccessor) Raw() ([]ChimeraGalleryAct, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ChimeraGalleryAct{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ChimeraGalleryActAccessor.LoadData to preload everything
func (a *ChimeraGalleryActAccessor) GroupData() {
	a._dataActID = map[float64]ChimeraGalleryAct{}
	a._dataIcon = map[string]ChimeraGalleryAct{}
	a._dataSort = map[float64]ChimeraGalleryAct{}
	for _, d := range a._data {
		a._dataActID[d.ActID] = d
		a._dataIcon[d.Icon] = d
		a._dataSort[d.Sort] = d
	}
}

// ByActID returns the ChimeraGalleryAct uniquely identified by ActID
//
// Error is only non-nil if the source errors out
func (a *ChimeraGalleryActAccessor) ByActID(identifier float64) (ChimeraGalleryAct, error) {
	if a._dataActID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ChimeraGalleryAct{}, err
			}
		}
		a.GroupData()
	}
	return a._dataActID[identifier], nil
}

// ByIcon returns the ChimeraGalleryAct uniquely identified by Icon
//
// Error is only non-nil if the source errors out
func (a *ChimeraGalleryActAccessor) ByIcon(identifier string) (ChimeraGalleryAct, error) {
	if a._dataIcon == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ChimeraGalleryAct{}, err
			}
		}
		a.GroupData()
	}
	return a._dataIcon[identifier], nil
}

// BySort returns the ChimeraGalleryAct uniquely identified by Sort
//
// Error is only non-nil if the source errors out
func (a *ChimeraGalleryActAccessor) BySort(identifier float64) (ChimeraGalleryAct, error) {
	if a._dataSort == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ChimeraGalleryAct{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSort[identifier], nil
}
