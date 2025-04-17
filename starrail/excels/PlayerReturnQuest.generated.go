package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type PlayerReturnQuest struct {
	GroupID       float64 `json:"GroupID"`
	ID            float64 `json:"ID"`
	LinearQuestID float64 `json:"LinearQuestID"`
}
type PlayerReturnQuestAccessor struct {
	_data   []PlayerReturnQuest
	_dataID map[float64]PlayerReturnQuest
}

// LoadData retrieves the data. Must be called before PlayerReturnQuest.GroupData
func (a *PlayerReturnQuestAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PlayerReturnQuest.json")
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
func (a *PlayerReturnQuestAccessor) Raw() ([]PlayerReturnQuest, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PlayerReturnQuest{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PlayerReturnQuestAccessor.LoadData to preload everything
func (a *PlayerReturnQuestAccessor) GroupData() {
	a._dataID = map[float64]PlayerReturnQuest{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the PlayerReturnQuest uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *PlayerReturnQuestAccessor) ByID(identifier float64) (PlayerReturnQuest, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlayerReturnQuest{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
