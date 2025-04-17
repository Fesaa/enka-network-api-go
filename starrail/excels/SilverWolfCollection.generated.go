package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type SilverWolfCollection struct {
	PositionID json.Number `json:"PositionID"`
	QuestID    json.Number `json:"QuestID"`
	Type       string      `json:"Type"`
	TypeParam  json.Number `json:"TypeParam"`
}
type SilverWolfCollectionAccessor struct {
	_data          []SilverWolfCollection
	_dataQuestID   map[json.Number]SilverWolfCollection
	_dataTypeParam map[json.Number]SilverWolfCollection
}

// LoadData retrieves the data. Must be called before SilverWolfCollection.GroupData
func (a *SilverWolfCollectionAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/SilverWolfCollection.json")
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
func (a *SilverWolfCollectionAccessor) Raw() ([]SilverWolfCollection, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []SilverWolfCollection{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with SilverWolfCollectionAccessor.LoadData to preload everything
func (a *SilverWolfCollectionAccessor) GroupData() {
	a._dataQuestID = map[json.Number]SilverWolfCollection{}
	a._dataTypeParam = map[json.Number]SilverWolfCollection{}
	for _, d := range a._data {
		a._dataQuestID[d.QuestID] = d
		a._dataTypeParam[d.TypeParam] = d
	}
}

// ByQuestID returns the SilverWolfCollection uniquely identified by QuestID
//
// Error is only non-nil if the source errors out
func (a *SilverWolfCollectionAccessor) ByQuestID(identifier json.Number) (SilverWolfCollection, error) {
	if a._dataQuestID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return SilverWolfCollection{}, err
			}
		}
		a.GroupData()
	}
	return a._dataQuestID[identifier], nil
}

// ByTypeParam returns the SilverWolfCollection uniquely identified by TypeParam
//
// Error is only non-nil if the source errors out
func (a *SilverWolfCollectionAccessor) ByTypeParam(identifier json.Number) (SilverWolfCollection, error) {
	if a._dataTypeParam == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return SilverWolfCollection{}, err
			}
		}
		a.GroupData()
	}
	return a._dataTypeParam[identifier], nil
}
