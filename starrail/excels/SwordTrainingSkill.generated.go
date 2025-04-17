package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type SwordTrainingSkill struct {
	AvatarStatusAddList []interface{}          `json:"AvatarStatusAddList"`
	Condition           float64                `json:"Condition"`
	Cost                SwordTrainingSkillCost `json:"Cost"`
	MazeBuffID          float64                `json:"MazeBuffID"`
	NextSkillIDList     []float64              `json:"NextSkillIDList"`
	ParamList           []float64              `json:"ParamList"`
	Rare                float64                `json:"Rare"`
	SkillID             float64                `json:"SkillID"`
	SkillIcon           string                 `json:"SkillIcon"`
	SkillName           hash.Hash              `json:"SkillName"`
	SkillPower          float64                `json:"SkillPower"`
	SkillRank           float64                `json:"SkillRank"`
	SkillTag            hash.Hash              `json:"SkillTag"`
	SkillTypeID         float64                `json:"SkillTypeID"`
}
type SwordTrainingSkillCost struct {
	ItemID float64 `json:"ItemID"`
}
type SwordTrainingSkillAccessor struct {
	_data           []SwordTrainingSkill
	_dataCondition  map[float64]SwordTrainingSkill
	_dataMazeBuffID map[float64]SwordTrainingSkill
	_dataSkillID    map[float64]SwordTrainingSkill
}

// LoadData retrieves the data. Must be called before SwordTrainingSkill.GroupData
func (a *SwordTrainingSkillAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/SwordTrainingSkill.json")
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
func (a *SwordTrainingSkillAccessor) Raw() ([]SwordTrainingSkill, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []SwordTrainingSkill{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with SwordTrainingSkillAccessor.LoadData to preload everything
func (a *SwordTrainingSkillAccessor) GroupData() {
	a._dataCondition = map[float64]SwordTrainingSkill{}
	a._dataMazeBuffID = map[float64]SwordTrainingSkill{}
	a._dataSkillID = map[float64]SwordTrainingSkill{}
	for _, d := range a._data {
		a._dataCondition[d.Condition] = d
		a._dataMazeBuffID[d.MazeBuffID] = d
		a._dataSkillID[d.SkillID] = d
	}
}

// ByCondition returns the SwordTrainingSkill uniquely identified by Condition
//
// Error is only non-nil if the source errors out
func (a *SwordTrainingSkillAccessor) ByCondition(identifier float64) (SwordTrainingSkill, error) {
	if a._dataCondition == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return SwordTrainingSkill{}, err
			}
		}
		a.GroupData()
	}
	return a._dataCondition[identifier], nil
}

// ByMazeBuffID returns the SwordTrainingSkill uniquely identified by MazeBuffID
//
// Error is only non-nil if the source errors out
func (a *SwordTrainingSkillAccessor) ByMazeBuffID(identifier float64) (SwordTrainingSkill, error) {
	if a._dataMazeBuffID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return SwordTrainingSkill{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMazeBuffID[identifier], nil
}

// BySkillID returns the SwordTrainingSkill uniquely identified by SkillID
//
// Error is only non-nil if the source errors out
func (a *SwordTrainingSkillAccessor) BySkillID(identifier float64) (SwordTrainingSkill, error) {
	if a._dataSkillID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return SwordTrainingSkill{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSkillID[identifier], nil
}
