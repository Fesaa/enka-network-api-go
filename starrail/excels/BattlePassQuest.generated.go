package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type BattlePassQuest struct {
	ID       float64 `json:"ID"`
	ShowTime bool    `json:"ShowTime"`
}
type BattlePassQuestAccessor struct {
	_data   []BattlePassQuest
	_dataID map[float64]BattlePassQuest
}

// LoadData retrieves the data. Must be called before BattlePassQuest.GroupData
func (a *BattlePassQuestAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/BattlePassQuest.json")
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
func (a *BattlePassQuestAccessor) Raw() ([]BattlePassQuest, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []BattlePassQuest{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with BattlePassQuestAccessor.LoadData to preload everything
func (a *BattlePassQuestAccessor) GroupData() {
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the BattlePassQuest uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *BattlePassQuestAccessor) ByID(identifier float64) (BattlePassQuest, error) {
	if a._dataID == nil {
		err := a.LoadData()
		if err != nil {
			return BattlePassQuest{}, err
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
