package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type AetherDivideMonster struct {
	MonsterID                 float64       `json:"MonsterID"`
	MonsterType               string        `json:"MonsterType"`
	SPMax                     hash.IntValue `json:"SPMax"`
	UltraSkillCutInPrefabPath string        `json:"UltraSkillCutInPrefabPath"`
}
type AetherDivideMonsterAccessor struct {
	_data          []AetherDivideMonster
	_dataMonsterID map[float64]AetherDivideMonster
}

// LoadData retrieves the data. Must be called before AetherDivideMonster.GroupData
func (a *AetherDivideMonsterAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AetherDivideMonster.json")
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
func (a *AetherDivideMonsterAccessor) Raw() ([]AetherDivideMonster, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AetherDivideMonster{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AetherDivideMonsterAccessor.LoadData to preload everything
func (a *AetherDivideMonsterAccessor) GroupData() {
	for _, d := range a._data {
		a._dataMonsterID[d.MonsterID] = d
	}
}

// ByMonsterID returns the AetherDivideMonster uniquely identified by MonsterID
//
// Error is only non-nil if the source errors out
func (a *AetherDivideMonsterAccessor) ByMonsterID(identifier float64) (AetherDivideMonster, error) {
	if a._dataMonsterID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDivideMonster{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMonsterID[identifier], nil
}
