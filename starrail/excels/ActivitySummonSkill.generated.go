package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type ActivitySummonSkill struct {
	SimpleSkillDesc     hash.Hash       `json:"SimpleSkillDesc"`
	SimpleSkillParmList []hash.IntValue `json:"SimpleSkillParmList"`
	SkillDesc           hash.Hash       `json:"SkillDesc"`
	SkillID             float64         `json:"SkillID"`
	SkillIconPath       string          `json:"SkillIconPath"`
	SkillName           hash.Hash       `json:"SkillName"`
	SkillParmList       []hash.IntValue `json:"SkillParmList"`
	SkillTriggerKey     string          `json:"SkillTriggerKey"`
	SkillType           string          `json:"SkillType"`
}
type ActivitySummonSkillAccessor struct {
	_data        []ActivitySummonSkill
	_dataSkillID map[float64]ActivitySummonSkill
}

// LoadData retrieves the data. Must be called before ActivitySummonSkill.GroupData
func (a *ActivitySummonSkillAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActivitySummonSkill.json")
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
func (a *ActivitySummonSkillAccessor) Raw() ([]ActivitySummonSkill, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActivitySummonSkill{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ActivitySummonSkillAccessor.LoadData to preload everything
func (a *ActivitySummonSkillAccessor) GroupData() {
	for _, d := range a._data {
		a._dataSkillID[d.SkillID] = d
	}
}

// BySkillID returns the ActivitySummonSkill uniquely identified by SkillID
//
// Error is only non-nil if the source errors out
func (a *ActivitySummonSkillAccessor) BySkillID(identifier float64) (ActivitySummonSkill, error) {
	if a._dataSkillID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivitySummonSkill{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSkillID[identifier], nil
}
