package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type NormalMode struct {
	Desc01       string  `json:"Desc01"`
	Desc02       string  `json:"Desc02"`
	Desc03       string  `json:"Desc03"`
	NormalModeID float64 `json:"NormalModeID"`
	Title        string  `json:"Title"`
}
type NormalModeAccessor struct {
	_data             []NormalMode
	_dataDesc01       map[string]NormalMode
	_dataNormalModeID map[float64]NormalMode
	_dataTitle        map[string]NormalMode
}

// LoadData retrieves the data. Must be called before NormalMode.GroupData
func (a *NormalModeAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/NormalMode.json")
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
func (a *NormalModeAccessor) Raw() ([]NormalMode, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []NormalMode{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with NormalModeAccessor.LoadData to preload everything
func (a *NormalModeAccessor) GroupData() {
	a._dataDesc01 = map[string]NormalMode{}
	a._dataNormalModeID = map[float64]NormalMode{}
	a._dataTitle = map[string]NormalMode{}
	for _, d := range a._data {
		a._dataDesc01[d.Desc01] = d
		a._dataNormalModeID[d.NormalModeID] = d
		a._dataTitle[d.Title] = d
	}
}

// ByDesc01 returns the NormalMode uniquely identified by Desc01
//
// Error is only non-nil if the source errors out
func (a *NormalModeAccessor) ByDesc01(identifier string) (NormalMode, error) {
	if a._dataDesc01 == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return NormalMode{}, err
			}
		}
		a.GroupData()
	}
	return a._dataDesc01[identifier], nil
}

// ByNormalModeID returns the NormalMode uniquely identified by NormalModeID
//
// Error is only non-nil if the source errors out
func (a *NormalModeAccessor) ByNormalModeID(identifier float64) (NormalMode, error) {
	if a._dataNormalModeID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return NormalMode{}, err
			}
		}
		a.GroupData()
	}
	return a._dataNormalModeID[identifier], nil
}

// ByTitle returns the NormalMode uniquely identified by Title
//
// Error is only non-nil if the source errors out
func (a *NormalModeAccessor) ByTitle(identifier string) (NormalMode, error) {
	if a._dataTitle == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return NormalMode{}, err
			}
		}
		a.GroupData()
	}
	return a._dataTitle[identifier], nil
}
