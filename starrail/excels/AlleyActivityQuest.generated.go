package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type AlleyActivityQuest struct {
	ID           float64   `json:"ID"`
	MainTabTitle hash.Hash `json:"MainTabTitle"`
	QuestList    []float64 `json:"QuestList"`
	SubTab       float64   `json:"SubTab"`
}
type AlleyActivityQuestAccessor struct {
	_data   []AlleyActivityQuest
	_dataID map[float64]AlleyActivityQuest
}

// LoadData retrieves the data. Must be called before AlleyActivityQuest.GroupData
func (a *AlleyActivityQuestAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AlleyActivityQuest.json")
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
func (a *AlleyActivityQuestAccessor) Raw() ([]AlleyActivityQuest, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AlleyActivityQuest{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AlleyActivityQuestAccessor.LoadData to preload everything
func (a *AlleyActivityQuestAccessor) GroupData() {
	a._dataID = map[float64]AlleyActivityQuest{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the AlleyActivityQuest uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *AlleyActivityQuestAccessor) ByID(identifier float64) (AlleyActivityQuest, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AlleyActivityQuest{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
