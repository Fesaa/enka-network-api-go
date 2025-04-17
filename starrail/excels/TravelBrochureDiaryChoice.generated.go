package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type TravelBrochureDiaryChoice struct {
	ChoiceMessage hash.Hash `json:"ChoiceMessage"`
	DetailMessage hash.Hash `json:"DetailMessage"`
	ID            float64   `json:"ID"`
}
type TravelBrochureDiaryChoiceAccessor struct {
	_data   []TravelBrochureDiaryChoice
	_dataID map[float64]TravelBrochureDiaryChoice
}

// LoadData retrieves the data. Must be called before TravelBrochureDiaryChoice.GroupData
func (a *TravelBrochureDiaryChoiceAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TravelBrochureDiaryChoice.json")
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
func (a *TravelBrochureDiaryChoiceAccessor) Raw() ([]TravelBrochureDiaryChoice, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TravelBrochureDiaryChoice{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with TravelBrochureDiaryChoiceAccessor.LoadData to preload everything
func (a *TravelBrochureDiaryChoiceAccessor) GroupData() {
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the TravelBrochureDiaryChoice uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *TravelBrochureDiaryChoiceAccessor) ByID(identifier float64) (TravelBrochureDiaryChoice, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TravelBrochureDiaryChoice{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
