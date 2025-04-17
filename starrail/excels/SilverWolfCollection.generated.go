package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type SilverWolfCollection struct {
	PositionID json.Number `json:"PositionID"`
	QuestID    json.Number `json:"QuestID"`
	Type       string      `json:"Type"`
	TypeParam  json.Number `json:"TypeParam"`
}
type SilverWolfCollectionAccessor struct {
	_data []SilverWolfCollection
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
