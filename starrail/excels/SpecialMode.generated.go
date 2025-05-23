package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type SpecialMode struct {
	Desc01        string  `json:"Desc01"`
	Desc02        string  `json:"Desc02"`
	Desc03        string  `json:"Desc03"`
	IsUImode      bool    `json:"IsUImode"`
	PuzzleType    float64 `json:"PuzzleType"`
	SpecialModeID float64 `json:"SpecialModeID"`
	Title         string  `json:"Title"`
}
type SpecialModeAccessor struct {
	_data              []SpecialMode
	_dataDesc01        map[string]SpecialMode
	_dataSpecialModeID map[float64]SpecialMode
	_dataTitle         map[string]SpecialMode
}

// LoadData retrieves the data. Must be called before SpecialMode.GroupData
func (a *SpecialModeAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/SpecialMode.json")
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
func (a *SpecialModeAccessor) Raw() ([]SpecialMode, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []SpecialMode{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with SpecialModeAccessor.LoadData to preload everything
func (a *SpecialModeAccessor) GroupData() {
	a._dataDesc01 = map[string]SpecialMode{}
	a._dataSpecialModeID = map[float64]SpecialMode{}
	a._dataTitle = map[string]SpecialMode{}
	for _, d := range a._data {
		a._dataDesc01[d.Desc01] = d
		a._dataSpecialModeID[d.SpecialModeID] = d
		a._dataTitle[d.Title] = d
	}
}

// ByDesc01 returns the SpecialMode uniquely identified by Desc01
//
// Error is only non-nil if the source errors out
func (a *SpecialModeAccessor) ByDesc01(identifier string) (SpecialMode, error) {
	if a._dataDesc01 == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return SpecialMode{}, err
			}
		}
		a.GroupData()
	}
	return a._dataDesc01[identifier], nil
}

// BySpecialModeID returns the SpecialMode uniquely identified by SpecialModeID
//
// Error is only non-nil if the source errors out
func (a *SpecialModeAccessor) BySpecialModeID(identifier float64) (SpecialMode, error) {
	if a._dataSpecialModeID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return SpecialMode{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSpecialModeID[identifier], nil
}

// ByTitle returns the SpecialMode uniquely identified by Title
//
// Error is only non-nil if the source errors out
func (a *SpecialModeAccessor) ByTitle(identifier string) (SpecialMode, error) {
	if a._dataTitle == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return SpecialMode{}, err
			}
		}
		a.GroupData()
	}
	return a._dataTitle[identifier], nil
}
