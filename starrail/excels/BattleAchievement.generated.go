package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type BattleAchievement struct {
	AbilityName         string   `json:"AbilityName"`
	BattleAchievementID float64  `json:"BattleAchievementID"`
	ExcludeTagList      []string `json:"ExcludeTagList"`
	GameModeGroup       float64  `json:"GameModeGroup"`
	NeedTagList         []string `json:"NeedTagList"`
}
type BattleAchievementAccessor struct {
	_data                    []BattleAchievement
	_dataBattleAchievementID map[float64]BattleAchievement
}

// LoadData retrieves the data. Must be called before BattleAchievement.GroupData
func (a *BattleAchievementAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/BattleAchievement.json")
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
func (a *BattleAchievementAccessor) Raw() ([]BattleAchievement, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []BattleAchievement{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with BattleAchievementAccessor.LoadData to preload everything
func (a *BattleAchievementAccessor) GroupData() {
	a._dataBattleAchievementID = map[float64]BattleAchievement{}
	for _, d := range a._data {
		a._dataBattleAchievementID[d.BattleAchievementID] = d
	}
}

// ByBattleAchievementID returns the BattleAchievement uniquely identified by BattleAchievementID
//
// Error is only non-nil if the source errors out
func (a *BattleAchievementAccessor) ByBattleAchievementID(identifier float64) (BattleAchievement, error) {
	if a._dataBattleAchievementID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return BattleAchievement{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBattleAchievementID[identifier], nil
}
