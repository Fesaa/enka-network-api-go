package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type GuideRogueTab struct {
	GuideType string                 `json:"GuideType"`
	ID        json.Number            `json:"ID"`
	IconPath  string                 `json:"IconPath"`
	Name      map[string]json.Number `json:"Name"`
	Priority  json.Number            `json:"Priority"`
	ResBarKey string                 `json:"ResBarKey"`
}
type GuideRogueTabAccessor struct {
	_data          []GuideRogueTab
	_dataGuideType map[string]GuideRogueTab
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
	for _, d := range a._data {
		a._dataGuideType[d.GuideType] = d
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
