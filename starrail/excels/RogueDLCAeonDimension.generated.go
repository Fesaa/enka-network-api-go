package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type RogueDLCAeonDimension struct {
	AeonDimensionID       json.Number            `json:"AeonDimensionID"`
	AeonDimensionMaxPoint json.Number            `json:"AeonDimensionMaxPoint"`
	AeonIcon              string                 `json:"AeonIcon"`
	DimensionIcon         string                 `json:"DimensionIcon"`
	PlayShortDesc         map[string]json.Number `json:"PlayShortDesc"`
}
type RogueDLCAeonDimensionAccessor struct {
	_data              []RogueDLCAeonDimension
	_dataAeonIcon      map[string]RogueDLCAeonDimension
	_dataDimensionIcon map[string]RogueDLCAeonDimension
}

// LoadData retrieves the data. Must be called before RogueDLCAeonDimension.GroupData
func (a *RogueDLCAeonDimensionAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueDLCAeonDimension.json")
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
func (a *RogueDLCAeonDimensionAccessor) Raw() ([]RogueDLCAeonDimension, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueDLCAeonDimension{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueDLCAeonDimensionAccessor.LoadData to preload everything
func (a *RogueDLCAeonDimensionAccessor) GroupData() {
	a._dataAeonIcon = map[string]RogueDLCAeonDimension{}
	a._dataDimensionIcon = map[string]RogueDLCAeonDimension{}
	for _, d := range a._data {
		a._dataAeonIcon[d.AeonIcon] = d
		a._dataDimensionIcon[d.DimensionIcon] = d
	}
}

// ByAeonIcon returns the RogueDLCAeonDimension uniquely identified by AeonIcon
//
// Error is only non-nil if the source errors out
func (a *RogueDLCAeonDimensionAccessor) ByAeonIcon(identifier string) (RogueDLCAeonDimension, error) {
	if a._dataAeonIcon == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueDLCAeonDimension{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAeonIcon[identifier], nil
}

// ByDimensionIcon returns the RogueDLCAeonDimension uniquely identified by DimensionIcon
//
// Error is only non-nil if the source errors out
func (a *RogueDLCAeonDimensionAccessor) ByDimensionIcon(identifier string) (RogueDLCAeonDimension, error) {
	if a._dataDimensionIcon == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueDLCAeonDimension{}, err
			}
		}
		a.GroupData()
	}
	return a._dataDimensionIcon[identifier], nil
}
