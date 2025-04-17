package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type DrinkMakerCupData struct {
	AudioEvent     string                 `json:"AudioEvent"`
	Capacity       json.Number            `json:"Capacity"`
	CupID          json.Number            `json:"CupID"`
	CupName        map[string]json.Number `json:"CupName"`
	IceCount       []json.Number          `json:"IceCount"`
	IconPath       string                 `json:"IconPath"`
	IncludeTagList []json.Number          `json:"IncludeTagList"`
	PerLayerHeight []json.Number          `json:"PerLayerHeight"`
	PrefabPath     string                 `json:"PrefabPath"`
	Type           string                 `json:"Type"`
}
type DrinkMakerCupDataAccessor struct {
	_data           []DrinkMakerCupData
	_dataAudioEvent map[string]DrinkMakerCupData
	_dataCupID      map[json.Number]DrinkMakerCupData
	_dataIconPath   map[string]DrinkMakerCupData
	_dataPrefabPath map[string]DrinkMakerCupData
}

// LoadData retrieves the data. Must be called before DrinkMakerCupData.GroupData
func (a *DrinkMakerCupDataAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/DrinkMakerCupData.json")
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
func (a *DrinkMakerCupDataAccessor) Raw() ([]DrinkMakerCupData, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []DrinkMakerCupData{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with DrinkMakerCupDataAccessor.LoadData to preload everything
func (a *DrinkMakerCupDataAccessor) GroupData() {
	a._dataAudioEvent = map[string]DrinkMakerCupData{}
	a._dataCupID = map[json.Number]DrinkMakerCupData{}
	a._dataIconPath = map[string]DrinkMakerCupData{}
	a._dataPrefabPath = map[string]DrinkMakerCupData{}
	for _, d := range a._data {
		a._dataAudioEvent[d.AudioEvent] = d
		a._dataCupID[d.CupID] = d
		a._dataIconPath[d.IconPath] = d
		a._dataPrefabPath[d.PrefabPath] = d
	}
}

// ByAudioEvent returns the DrinkMakerCupData uniquely identified by AudioEvent
//
// Error is only non-nil if the source errors out
func (a *DrinkMakerCupDataAccessor) ByAudioEvent(identifier string) (DrinkMakerCupData, error) {
	if a._dataAudioEvent == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return DrinkMakerCupData{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAudioEvent[identifier], nil
}

// ByCupID returns the DrinkMakerCupData uniquely identified by CupID
//
// Error is only non-nil if the source errors out
func (a *DrinkMakerCupDataAccessor) ByCupID(identifier json.Number) (DrinkMakerCupData, error) {
	if a._dataCupID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return DrinkMakerCupData{}, err
			}
		}
		a.GroupData()
	}
	return a._dataCupID[identifier], nil
}

// ByIconPath returns the DrinkMakerCupData uniquely identified by IconPath
//
// Error is only non-nil if the source errors out
func (a *DrinkMakerCupDataAccessor) ByIconPath(identifier string) (DrinkMakerCupData, error) {
	if a._dataIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return DrinkMakerCupData{}, err
			}
		}
		a.GroupData()
	}
	return a._dataIconPath[identifier], nil
}

// ByPrefabPath returns the DrinkMakerCupData uniquely identified by PrefabPath
//
// Error is only non-nil if the source errors out
func (a *DrinkMakerCupDataAccessor) ByPrefabPath(identifier string) (DrinkMakerCupData, error) {
	if a._dataPrefabPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return DrinkMakerCupData{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPrefabPath[identifier], nil
}
