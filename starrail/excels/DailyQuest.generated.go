package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type DailyQuest struct {
	DailyID   float64   `json:"DailyID"`
	IsDelete  bool      `json:"IsDelete"`
	MaxLevel  float64   `json:"MaxLevel"`
	MinLevel  float64   `json:"MinLevel"`
	QuestList []float64 `json:"QuestList"`
}
type DailyQuestAccessor struct {
	_data        []DailyQuest
	_dataDailyID map[float64]DailyQuest
}

// LoadData retrieves the data. Must be called before DailyQuest.GroupData
func (a *DailyQuestAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/DailyQuest.json")
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
func (a *DailyQuestAccessor) Raw() ([]DailyQuest, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []DailyQuest{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with DailyQuestAccessor.LoadData to preload everything
func (a *DailyQuestAccessor) GroupData() {
	for _, d := range a._data {
		a._dataDailyID[d.DailyID] = d
	}
}

// ByDailyID returns the DailyQuest uniquely identified by DailyID
//
// Error is only non-nil if the source errors out
func (a *DailyQuestAccessor) ByDailyID(identifier float64) (DailyQuest, error) {
	if a._dataDailyID == nil {
		err := a.LoadData()
		if err != nil {
			return DailyQuest{}, err
		}
		a.GroupData()
	}
	return a._dataDailyID[identifier], nil
}
