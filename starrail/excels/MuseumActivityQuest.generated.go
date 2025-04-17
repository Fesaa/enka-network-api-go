package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type MuseumActivityQuest struct {
	ID            json.Number   `json:"ID"`
	Name          string        `json:"Name"`
	QuestIconPath string        `json:"QuestIconPath"`
	QuestList     []json.Number `json:"QuestList"`
}
type MuseumActivityQuestAccessor struct {
	_data     []MuseumActivityQuest
	_dataName map[string]MuseumActivityQuest
}

// LoadData retrieves the data. Must be called before MuseumActivityQuest.GroupData
func (a *MuseumActivityQuestAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MuseumActivityQuest.json")
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
func (a *MuseumActivityQuestAccessor) Raw() ([]MuseumActivityQuest, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MuseumActivityQuest{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MuseumActivityQuestAccessor.LoadData to preload everything
func (a *MuseumActivityQuestAccessor) GroupData() {
	a._dataName = map[string]MuseumActivityQuest{}
	for _, d := range a._data {
		a._dataName[d.Name] = d
	}
}

// ByName returns the MuseumActivityQuest uniquely identified by Name
//
// Error is only non-nil if the source errors out
func (a *MuseumActivityQuestAccessor) ByName(identifier string) (MuseumActivityQuest, error) {
	if a._dataName == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MuseumActivityQuest{}, err
			}
		}
		a.GroupData()
	}
	return a._dataName[identifier], nil
}
