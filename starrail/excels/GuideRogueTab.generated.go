package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type GuideRogueTab struct {
	GuideType string    `json:"GuideType"`
	ID        float64   `json:"ID"`
	IconPath  string    `json:"IconPath"`
	Name      hash.Hash `json:"Name"`
	Priority  float64   `json:"Priority"`
	ResBarKey string    `json:"ResBarKey"`
}
type GuideRogueTabAccessor struct {
	_data          []GuideRogueTab
	_dataGuideType map[string]GuideRogueTab
	_dataID        map[float64]GuideRogueTab
	_dataPriority  map[float64]GuideRogueTab
}

// LoadData retrieves the data. Must be called before GuideRogueTab.GroupData
func (a *GuideRogueTabAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/GuideRogueTab.json")
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
func (a *GuideRogueTabAccessor) Raw() ([]GuideRogueTab, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []GuideRogueTab{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with GuideRogueTabAccessor.LoadData to preload everything
func (a *GuideRogueTabAccessor) GroupData() {
	a._dataGuideType = map[string]GuideRogueTab{}
	a._dataID = map[float64]GuideRogueTab{}
	a._dataPriority = map[float64]GuideRogueTab{}
	for _, d := range a._data {
		a._dataGuideType[d.GuideType] = d
		a._dataID[d.ID] = d
		a._dataPriority[d.Priority] = d
	}
}

// ByGuideType returns the GuideRogueTab uniquely identified by GuideType
//
// Error is only non-nil if the source errors out
func (a *GuideRogueTabAccessor) ByGuideType(identifier string) (GuideRogueTab, error) {
	if a._dataGuideType == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return GuideRogueTab{}, err
			}
		}
		a.GroupData()
	}
	return a._dataGuideType[identifier], nil
}

// ByID returns the GuideRogueTab uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *GuideRogueTabAccessor) ByID(identifier float64) (GuideRogueTab, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return GuideRogueTab{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByPriority returns the GuideRogueTab uniquely identified by Priority
//
// Error is only non-nil if the source errors out
func (a *GuideRogueTabAccessor) ByPriority(identifier float64) (GuideRogueTab, error) {
	if a._dataPriority == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return GuideRogueTab{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPriority[identifier], nil
}
