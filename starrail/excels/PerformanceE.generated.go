package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type PerformanceE struct {
	ChangePlayerType     string  `json:"ChangePlayerType"`
	EndBlack             string  `json:"EndBlack"`
	EndWithCrack         bool    `json:"EndWithCrack"`
	FloorID              float64 `json:"FloorID"`
	IsIntroDialogue      bool    `json:"IsIntroDialogue"`
	IsSkip               string  `json:"IsSkip"`
	PerformanceCharacter string  `json:"PerformanceCharacter"`
	PerformanceID        float64 `json:"PerformanceID"`
	PerformancePath      string  `json:"PerformancePath"`
	PlaneID              float64 `json:"PlaneID"`
	StartBlack           string  `json:"StartBlack"`
}
type PerformanceEAccessor struct {
	_data              []PerformanceE
	_dataPerformanceID map[float64]PerformanceE
}

// LoadData retrieves the data. Must be called before PerformanceE.GroupData
func (a *PerformanceEAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PerformanceE.json")
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
func (a *PerformanceEAccessor) Raw() ([]PerformanceE, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PerformanceE{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PerformanceEAccessor.LoadData to preload everything
func (a *PerformanceEAccessor) GroupData() {
	for _, d := range a._data {
		a._dataPerformanceID[d.PerformanceID] = d
	}
}

// ByPerformanceID returns the PerformanceE uniquely identified by PerformanceID
//
// Error is only non-nil if the source errors out
func (a *PerformanceEAccessor) ByPerformanceID(identifier float64) (PerformanceE, error) {
	if a._dataPerformanceID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PerformanceE{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPerformanceID[identifier], nil
}
