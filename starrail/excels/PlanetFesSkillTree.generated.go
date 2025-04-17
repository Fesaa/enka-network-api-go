package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type PlanetFesSkillTree struct {
	Icon            string        `json:"Icon"`
	IsImportant     bool          `json:"IsImportant"`
	LevelCostList   []float64     `json:"LevelCostList"`
	LevelSkillList  []float64     `json:"LevelSkillList"`
	MaxLevel        float64       `json:"MaxLevel"`
	Name            hash.Hash     `json:"Name"`
	NextSkillIDList []float64     `json:"NextSkillIDList"`
	Phase           float64       `json:"Phase"`
	SkillID         float64       `json:"SkillID"`
	UnlockIDList    []interface{} `json:"UnlockIDList"`
}
type PlanetFesSkillTreeAccessor struct {
	_data        []PlanetFesSkillTree
	_dataSkillID map[float64]PlanetFesSkillTree
}

// LoadData retrieves the data. Must be called before PlanetFesSkillTree.GroupData
func (a *PlanetFesSkillTreeAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PlanetFesSkillTree.json")
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
func (a *PlanetFesSkillTreeAccessor) Raw() ([]PlanetFesSkillTree, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PlanetFesSkillTree{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PlanetFesSkillTreeAccessor.LoadData to preload everything
func (a *PlanetFesSkillTreeAccessor) GroupData() {
	for _, d := range a._data {
		a._dataSkillID[d.SkillID] = d
	}
}

// BySkillID returns the PlanetFesSkillTree uniquely identified by SkillID
//
// Error is only non-nil if the source errors out
func (a *PlanetFesSkillTreeAccessor) BySkillID(identifier float64) (PlanetFesSkillTree, error) {
	if a._dataSkillID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlanetFesSkillTree{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSkillID[identifier], nil
}
