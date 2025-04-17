package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type RogueAeonDisplay struct {
	AeonBuffIcon       string    `json:"AeonBuffIcon"`
	AeonFigure         string    `json:"AeonFigure"`
	AeonIcon           string    `json:"AeonIcon"`
	AeonImage          string    `json:"AeonImage"`
	DisplayID          float64   `json:"DisplayID"`
	RogueAeonName      hash.Hash `json:"RogueAeonName"`
	RogueAeonPathName  hash.Hash `json:"RogueAeonPathName"`
	RogueAeonPathName2 hash.Hash `json:"RogueAeonPathName2"`
}
type RogueAeonDisplayAccessor struct {
	_data             []RogueAeonDisplay
	_dataAeonImage    map[string]RogueAeonDisplay
	_dataAeonIcon     map[string]RogueAeonDisplay
	_dataAeonFigure   map[string]RogueAeonDisplay
	_dataAeonBuffIcon map[string]RogueAeonDisplay
	_dataDisplayID    map[float64]RogueAeonDisplay
}

// LoadData retrieves the data. Must be called before RogueAeonDisplay.GroupData
func (a *RogueAeonDisplayAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueAeonDisplay.json")
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
func (a *RogueAeonDisplayAccessor) Raw() ([]RogueAeonDisplay, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueAeonDisplay{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueAeonDisplayAccessor.LoadData to preload everything
func (a *RogueAeonDisplayAccessor) GroupData() {
	for _, d := range a._data {
		a._dataAeonImage[d.AeonImage] = d
		a._dataAeonIcon[d.AeonIcon] = d
		a._dataAeonFigure[d.AeonFigure] = d
		a._dataAeonBuffIcon[d.AeonBuffIcon] = d
		a._dataDisplayID[d.DisplayID] = d
	}
}

// ByAeonImage returns the RogueAeonDisplay uniquely identified by AeonImage
//
// Error is only non-nil if the source errors out
func (a *RogueAeonDisplayAccessor) ByAeonImage(identifier string) (RogueAeonDisplay, error) {
	if a._dataAeonImage == nil {
		err := a.LoadData()
		if err != nil {
			return RogueAeonDisplay{}, err
		}
		a.GroupData()
	}
	return a._dataAeonImage[identifier], nil
}

// ByAeonIcon returns the RogueAeonDisplay uniquely identified by AeonIcon
//
// Error is only non-nil if the source errors out
func (a *RogueAeonDisplayAccessor) ByAeonIcon(identifier string) (RogueAeonDisplay, error) {
	if a._dataAeonIcon == nil {
		err := a.LoadData()
		if err != nil {
			return RogueAeonDisplay{}, err
		}
		a.GroupData()
	}
	return a._dataAeonIcon[identifier], nil
}

// ByAeonFigure returns the RogueAeonDisplay uniquely identified by AeonFigure
//
// Error is only non-nil if the source errors out
func (a *RogueAeonDisplayAccessor) ByAeonFigure(identifier string) (RogueAeonDisplay, error) {
	if a._dataAeonFigure == nil {
		err := a.LoadData()
		if err != nil {
			return RogueAeonDisplay{}, err
		}
		a.GroupData()
	}
	return a._dataAeonFigure[identifier], nil
}

// ByAeonBuffIcon returns the RogueAeonDisplay uniquely identified by AeonBuffIcon
//
// Error is only non-nil if the source errors out
func (a *RogueAeonDisplayAccessor) ByAeonBuffIcon(identifier string) (RogueAeonDisplay, error) {
	if a._dataAeonBuffIcon == nil {
		err := a.LoadData()
		if err != nil {
			return RogueAeonDisplay{}, err
		}
		a.GroupData()
	}
	return a._dataAeonBuffIcon[identifier], nil
}

// ByDisplayID returns the RogueAeonDisplay uniquely identified by DisplayID
//
// Error is only non-nil if the source errors out
func (a *RogueAeonDisplayAccessor) ByDisplayID(identifier float64) (RogueAeonDisplay, error) {
	if a._dataDisplayID == nil {
		err := a.LoadData()
		if err != nil {
			return RogueAeonDisplay{}, err
		}
		a.GroupData()
	}
	return a._dataDisplayID[identifier], nil
}
