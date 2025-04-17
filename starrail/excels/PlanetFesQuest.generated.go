package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type PlanetFesQuest struct {
	Description    hash.Hash                      `json:"Description"`
	FinishwayID    float64                        `json:"FinishwayID"`
	ID             float64                        `json:"ID"`
	IconPath       string                         `json:"IconPath"`
	Name           hash.Hash                      `json:"Name"`
	QuestType      string                         `json:"QuestType"`
	RewardItemList []PlanetFesQuestRewardItemList `json:"RewardItemList"`
}
type PlanetFesQuestRewardItemList struct {
	ItemID  float64 `json:"ItemID"`
	ItemNum float64 `json:"ItemNum"`
}
type PlanetFesQuestAccessor struct {
	_data            []PlanetFesQuest
	_dataFinishwayID map[float64]PlanetFesQuest
	_dataID          map[float64]PlanetFesQuest
}

// LoadData retrieves the data. Must be called before PlanetFesQuest.GroupData
func (a *PlanetFesQuestAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PlanetFesQuest.json")
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
func (a *PlanetFesQuestAccessor) Raw() ([]PlanetFesQuest, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PlanetFesQuest{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PlanetFesQuestAccessor.LoadData to preload everything
func (a *PlanetFesQuestAccessor) GroupData() {
	for _, d := range a._data {
		a._dataFinishwayID[d.FinishwayID] = d
		a._dataID[d.ID] = d
	}
}

// ByFinishwayID returns the PlanetFesQuest uniquely identified by FinishwayID
//
// Error is only non-nil if the source errors out
func (a *PlanetFesQuestAccessor) ByFinishwayID(identifier float64) (PlanetFesQuest, error) {
	if a._dataFinishwayID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlanetFesQuest{}, err
			}
		}
		a.GroupData()
	}
	return a._dataFinishwayID[identifier], nil
}

// ByID returns the PlanetFesQuest uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *PlanetFesQuestAccessor) ByID(identifier float64) (PlanetFesQuest, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlanetFesQuest{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
