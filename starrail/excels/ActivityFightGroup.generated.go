package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type ActivityFightGroup struct {
	ActivityFightGroupID       float64   `json:"ActivityFightGroupID"`
	ActivityFightGroupIconPath string    `json:"ActivityFightGroupIconPath"`
	BattleAreaGroupID          float64   `json:"BattleAreaGroupID"`
	BattleAreaID               float64   `json:"BattleAreaID"`
	FightStageDesc             hash.Hash `json:"FightStageDesc"`
	FightStageTitle            hash.Hash `json:"FightStageTitle"`
	FloorID                    float64   `json:"FloorID"`
	PlaneID                    float64   `json:"PlaneID"`
	SpecialAvatarID            float64   `json:"SpecialAvatarID"`
}
type ActivityFightGroupAccessor struct {
	_data                     []ActivityFightGroup
	_dataActivityFightGroupID map[float64]ActivityFightGroup
}

// LoadData retrieves the data. Must be called before ActivityFightGroup.GroupData
func (a *ActivityFightGroupAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActivityFightGroup.json")
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
func (a *ActivityFightGroupAccessor) Raw() ([]ActivityFightGroup, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActivityFightGroup{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ActivityFightGroupAccessor.LoadData to preload everything
func (a *ActivityFightGroupAccessor) GroupData() {
	a._dataActivityFightGroupID = map[float64]ActivityFightGroup{}
	for _, d := range a._data {
		a._dataActivityFightGroupID[d.ActivityFightGroupID] = d
	}
}

// ByActivityFightGroupID returns the ActivityFightGroup uniquely identified by ActivityFightGroupID
//
// Error is only non-nil if the source errors out
func (a *ActivityFightGroupAccessor) ByActivityFightGroupID(identifier float64) (ActivityFightGroup, error) {
	if a._dataActivityFightGroupID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityFightGroup{}, err
			}
		}
		a.GroupData()
	}
	return a._dataActivityFightGroupID[identifier], nil
}
