package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type TreasureDungeoActivityQuest struct {
	DungeonGroupID float64   `json:"DungeonGroupID"`
	ID             float64   `json:"ID"`
	Name           hash.Hash `json:"Name"`
	QuestList      []float64 `json:"QuestList"`
}
type TreasureDungeoActivityQuestAccessor struct {
	_data []TreasureDungeoActivityQuest
}

// LoadData retrieves the data. Must be called before TreasureDungeoActivityQuest.GroupData
func (a *TreasureDungeoActivityQuestAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TreasureDungeoActivityQuest.json")
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
func (a *TreasureDungeoActivityQuestAccessor) Raw() ([]TreasureDungeoActivityQuest, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TreasureDungeoActivityQuest{}, err
		}
	}
	return a._data, nil
}
