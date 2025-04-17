package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type TravelBrochureDiaryGroup struct {
	ChoiceIDList     []json.Number          `json:"ChoiceIDList"`
	DiaryDescription map[string]json.Number `json:"DiaryDescription"`
	ID               json.Number            `json:"ID"`
	TextIDList       []json.Number          `json:"TextIDList"`
}
type TravelBrochureDiaryGroupAccessor struct {
	_data   []TravelBrochureDiaryGroup
	_dataID map[json.Number]TravelBrochureDiaryGroup
}

// LoadData retrieves the data. Must be called before TravelBrochureDiaryGroup.GroupData
func (a *TravelBrochureDiaryGroupAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TravelBrochureDiaryGroup.json")
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
func (a *TravelBrochureDiaryGroupAccessor) Raw() ([]TravelBrochureDiaryGroup, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TravelBrochureDiaryGroup{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with TravelBrochureDiaryGroupAccessor.LoadData to preload everything
func (a *TravelBrochureDiaryGroupAccessor) GroupData() {
	a._dataID = map[json.Number]TravelBrochureDiaryGroup{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the TravelBrochureDiaryGroup uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *TravelBrochureDiaryGroupAccessor) ByID(identifier json.Number) (TravelBrochureDiaryGroup, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TravelBrochureDiaryGroup{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
