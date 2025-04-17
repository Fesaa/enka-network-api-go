package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type GuideRogueData struct {
	ID               float64                        `json:"ID"`
	IconPath         string                         `json:"IconPath"`
	Name             hash.Hash                      `json:"Name"`
	OpenConditions   []GuideRogueDataOpenConditions `json:"OpenConditions"`
	Priority         float64                        `json:"Priority"`
	RelatedID        float64                        `json:"RelatedID"`
	TabID            float64                        `json:"TabID"`
	TabIconPath      string                         `json:"TabIconPath"`
	UnlockConditions []interface{}                  `json:"UnlockConditions"`
}
type GuideRogueDataOpenConditions struct {
	Param string `json:"Param"`
	Type  string `json:"Type"`
}
type GuideRogueDataAccessor struct {
	_data          []GuideRogueData
	_dataID        map[float64]GuideRogueData
	_dataRelatedID map[float64]GuideRogueData
}

// LoadData retrieves the data. Must be called before GuideRogueData.GroupData
func (a *GuideRogueDataAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/GuideRogueData.json")
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
func (a *GuideRogueDataAccessor) Raw() ([]GuideRogueData, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []GuideRogueData{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with GuideRogueDataAccessor.LoadData to preload everything
func (a *GuideRogueDataAccessor) GroupData() {
	for _, d := range a._data {
		a._dataID[d.ID] = d
		a._dataRelatedID[d.RelatedID] = d
	}
}

// ByID returns the GuideRogueData uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *GuideRogueDataAccessor) ByID(identifier float64) (GuideRogueData, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return GuideRogueData{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByRelatedID returns the GuideRogueData uniquely identified by RelatedID
//
// Error is only non-nil if the source errors out
func (a *GuideRogueDataAccessor) ByRelatedID(identifier float64) (GuideRogueData, error) {
	if a._dataRelatedID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return GuideRogueData{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRelatedID[identifier], nil
}
