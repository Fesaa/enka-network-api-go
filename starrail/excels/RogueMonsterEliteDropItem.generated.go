package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type RogueMonsterEliteDropItem struct {
	MonsterEliteDropItemDisplayList []float64 `json:"MonsterEliteDropItemDisplayList"`
	MonsterEliteDropItemID          float64   `json:"MonsterEliteDropItemID"`
}
type RogueMonsterEliteDropItemAccessor struct {
	_data                       []RogueMonsterEliteDropItem
	_dataMonsterEliteDropItemID map[float64]RogueMonsterEliteDropItem
}

// LoadData retrieves the data. Must be called before RogueMonsterEliteDropItem.GroupData
func (a *RogueMonsterEliteDropItemAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueMonsterEliteDropItem.json")
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
func (a *RogueMonsterEliteDropItemAccessor) Raw() ([]RogueMonsterEliteDropItem, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueMonsterEliteDropItem{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueMonsterEliteDropItemAccessor.LoadData to preload everything
func (a *RogueMonsterEliteDropItemAccessor) GroupData() {
	a._dataMonsterEliteDropItemID = map[float64]RogueMonsterEliteDropItem{}
	for _, d := range a._data {
		a._dataMonsterEliteDropItemID[d.MonsterEliteDropItemID] = d
	}
}

// ByMonsterEliteDropItemID returns the RogueMonsterEliteDropItem uniquely identified by MonsterEliteDropItemID
//
// Error is only non-nil if the source errors out
func (a *RogueMonsterEliteDropItemAccessor) ByMonsterEliteDropItemID(identifier float64) (RogueMonsterEliteDropItem, error) {
	if a._dataMonsterEliteDropItemID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueMonsterEliteDropItem{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMonsterEliteDropItemID[identifier], nil
}
