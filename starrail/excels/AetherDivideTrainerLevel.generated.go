package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type AetherDivideTrainerLevel struct {
	ID               float64   `json:"ID"`
	IconPath         string    `json:"IconPath"`
	Name             hash.Hash `json:"Name"`
	QuestID          float64   `json:"QuestID"`
	QuestList        []float64 `json:"QuestList"`
	RareMonsterNumID float64   `json:"RareMonsterNumID"`
}
type AetherDivideTrainerLevelAccessor struct {
	_data         []AetherDivideTrainerLevel
	_dataID       map[float64]AetherDivideTrainerLevel
	_dataIconPath map[string]AetherDivideTrainerLevel
	_dataQuestID  map[float64]AetherDivideTrainerLevel
}

// LoadData retrieves the data. Must be called before AetherDivideTrainerLevel.GroupData
func (a *AetherDivideTrainerLevelAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AetherDivideTrainerLevel.json")
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
func (a *AetherDivideTrainerLevelAccessor) Raw() ([]AetherDivideTrainerLevel, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AetherDivideTrainerLevel{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AetherDivideTrainerLevelAccessor.LoadData to preload everything
func (a *AetherDivideTrainerLevelAccessor) GroupData() {
	a._dataID = map[float64]AetherDivideTrainerLevel{}
	a._dataIconPath = map[string]AetherDivideTrainerLevel{}
	a._dataQuestID = map[float64]AetherDivideTrainerLevel{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
		a._dataIconPath[d.IconPath] = d
		a._dataQuestID[d.QuestID] = d
	}
}

// ByID returns the AetherDivideTrainerLevel uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *AetherDivideTrainerLevelAccessor) ByID(identifier float64) (AetherDivideTrainerLevel, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDivideTrainerLevel{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByIconPath returns the AetherDivideTrainerLevel uniquely identified by IconPath
//
// Error is only non-nil if the source errors out
func (a *AetherDivideTrainerLevelAccessor) ByIconPath(identifier string) (AetherDivideTrainerLevel, error) {
	if a._dataIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDivideTrainerLevel{}, err
			}
		}
		a.GroupData()
	}
	return a._dataIconPath[identifier], nil
}

// ByQuestID returns the AetherDivideTrainerLevel uniquely identified by QuestID
//
// Error is only non-nil if the source errors out
func (a *AetherDivideTrainerLevelAccessor) ByQuestID(identifier float64) (AetherDivideTrainerLevel, error) {
	if a._dataQuestID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDivideTrainerLevel{}, err
			}
		}
		a.GroupData()
	}
	return a._dataQuestID[identifier], nil
}
