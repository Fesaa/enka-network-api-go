package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type ActivityRaidCollectionQuest struct {
	QuestList    []float64 `json:"QuestList"`
	QuestTabID   float64   `json:"QuestTabID"`
	QuestTabName hash.Hash `json:"QuestTabName"`
}
type ActivityRaidCollectionQuestAccessor struct {
	_data           []ActivityRaidCollectionQuest
	_dataQuestTabID map[float64]ActivityRaidCollectionQuest
}

// LoadData retrieves the data. Must be called before ActivityRaidCollectionQuest.GroupData
func (a *ActivityRaidCollectionQuestAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActivityRaidCollectionQuest.json")
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
func (a *ActivityRaidCollectionQuestAccessor) Raw() ([]ActivityRaidCollectionQuest, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActivityRaidCollectionQuest{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ActivityRaidCollectionQuestAccessor.LoadData to preload everything
func (a *ActivityRaidCollectionQuestAccessor) GroupData() {
	a._dataQuestTabID = map[float64]ActivityRaidCollectionQuest{}
	for _, d := range a._data {
		a._dataQuestTabID[d.QuestTabID] = d
	}
}

// ByQuestTabID returns the ActivityRaidCollectionQuest uniquely identified by QuestTabID
//
// Error is only non-nil if the source errors out
func (a *ActivityRaidCollectionQuestAccessor) ByQuestTabID(identifier float64) (ActivityRaidCollectionQuest, error) {
	if a._dataQuestTabID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityRaidCollectionQuest{}, err
			}
		}
		a.GroupData()
	}
	return a._dataQuestTabID[identifier], nil
}
