package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type ActivityTelevisionQuest struct {
	OriginalTabName hash.Hash `json:"OriginalTabName"`
	QuestGroupID    float64   `json:"QuestGroupID"`
	QuestIDList     []float64 `json:"QuestIDList"`
	TabName         hash.Hash `json:"TabName"`
}
type ActivityTelevisionQuestAccessor struct {
	_data             []ActivityTelevisionQuest
	_dataQuestGroupID map[float64]ActivityTelevisionQuest
}

// LoadData retrieves the data. Must be called before ActivityTelevisionQuest.GroupData
func (a *ActivityTelevisionQuestAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActivityTelevisionQuest.json")
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
func (a *ActivityTelevisionQuestAccessor) Raw() ([]ActivityTelevisionQuest, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActivityTelevisionQuest{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ActivityTelevisionQuestAccessor.LoadData to preload everything
func (a *ActivityTelevisionQuestAccessor) GroupData() {
	a._dataQuestGroupID = map[float64]ActivityTelevisionQuest{}
	for _, d := range a._data {
		a._dataQuestGroupID[d.QuestGroupID] = d
	}
}

// ByQuestGroupID returns the ActivityTelevisionQuest uniquely identified by QuestGroupID
//
// Error is only non-nil if the source errors out
func (a *ActivityTelevisionQuestAccessor) ByQuestGroupID(identifier float64) (ActivityTelevisionQuest, error) {
	if a._dataQuestGroupID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityTelevisionQuest{}, err
			}
		}
		a.GroupData()
	}
	return a._dataQuestGroupID[identifier], nil
}
