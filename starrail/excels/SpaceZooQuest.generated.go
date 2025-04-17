package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type SpaceZooQuest struct {
	ID           float64   `json:"ID"`
	QuestList    []float64 `json:"QuestList"`
	QuestTabName hash.Hash `json:"QuestTabName"`
	Type         string    `json:"Type"`
}
type SpaceZooQuestAccessor struct {
	_data   []SpaceZooQuest
	_dataID map[float64]SpaceZooQuest
}

// LoadData retrieves the data. Must be called before SpaceZooQuest.GroupData
func (a *SpaceZooQuestAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/SpaceZooQuest.json")
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
func (a *SpaceZooQuestAccessor) Raw() ([]SpaceZooQuest, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []SpaceZooQuest{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with SpaceZooQuestAccessor.LoadData to preload everything
func (a *SpaceZooQuestAccessor) GroupData() {
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the SpaceZooQuest uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *SpaceZooQuestAccessor) ByID(identifier float64) (SpaceZooQuest, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return SpaceZooQuest{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
