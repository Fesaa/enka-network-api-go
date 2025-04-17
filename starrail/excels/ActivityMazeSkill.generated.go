package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type ActivityMazeSkill struct {
	MazeSkillDesc      hash.Hash `json:"MazeSkillDesc"`
	MazeSkillId        float64   `json:"MazeSkillId"`
	MazeSkillName      hash.Hash `json:"MazeSkillName"`
	MazeSkilltype      float64   `json:"MazeSkilltype"`
	RelatedAvatarSkill float64   `json:"RelatedAvatarSkill"`
	SkillTriggerKey    string    `json:"SkillTriggerKey"`
}
type ActivityMazeSkillAccessor struct {
	_data                   []ActivityMazeSkill
	_dataMazeSkillId        map[float64]ActivityMazeSkill
	_dataRelatedAvatarSkill map[float64]ActivityMazeSkill
}

// LoadData retrieves the data. Must be called before ActivityMazeSkill.GroupData
func (a *ActivityMazeSkillAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActivityMazeSkill.json")
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
func (a *ActivityMazeSkillAccessor) Raw() ([]ActivityMazeSkill, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActivityMazeSkill{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ActivityMazeSkillAccessor.LoadData to preload everything
func (a *ActivityMazeSkillAccessor) GroupData() {
	for _, d := range a._data {
		a._dataMazeSkillId[d.MazeSkillId] = d
		a._dataRelatedAvatarSkill[d.RelatedAvatarSkill] = d
	}
}

// ByMazeSkillId returns the ActivityMazeSkill uniquely identified by MazeSkillId
//
// Error is only non-nil if the source errors out
func (a *ActivityMazeSkillAccessor) ByMazeSkillId(identifier float64) (ActivityMazeSkill, error) {
	if a._dataMazeSkillId == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityMazeSkill{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMazeSkillId[identifier], nil
}

// ByRelatedAvatarSkill returns the ActivityMazeSkill uniquely identified by RelatedAvatarSkill
//
// Error is only non-nil if the source errors out
func (a *ActivityMazeSkillAccessor) ByRelatedAvatarSkill(identifier float64) (ActivityMazeSkill, error) {
	if a._dataRelatedAvatarSkill == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityMazeSkill{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRelatedAvatarSkill[identifier], nil
}
