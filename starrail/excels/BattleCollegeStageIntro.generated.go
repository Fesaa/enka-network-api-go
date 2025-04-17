package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type BattleCollegeStageIntro struct {
	StageIntroDescID float64 `json:"StageIntroDescID"`
}
type BattleCollegeStageIntroAccessor struct {
	_data                 []BattleCollegeStageIntro
	_dataStageIntroDescID map[float64]BattleCollegeStageIntro
}

// LoadData retrieves the data. Must be called before BattleCollegeStageIntro.GroupData
func (a *BattleCollegeStageIntroAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/BattleCollegeStageIntro.json")
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
func (a *BattleCollegeStageIntroAccessor) Raw() ([]BattleCollegeStageIntro, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []BattleCollegeStageIntro{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with BattleCollegeStageIntroAccessor.LoadData to preload everything
func (a *BattleCollegeStageIntroAccessor) GroupData() {
	a._dataStageIntroDescID = map[float64]BattleCollegeStageIntro{}
	for _, d := range a._data {
		a._dataStageIntroDescID[d.StageIntroDescID] = d
	}
}

// ByStageIntroDescID returns the BattleCollegeStageIntro uniquely identified by StageIntroDescID
//
// Error is only non-nil if the source errors out
func (a *BattleCollegeStageIntroAccessor) ByStageIntroDescID(identifier float64) (BattleCollegeStageIntro, error) {
	if a._dataStageIntroDescID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return BattleCollegeStageIntro{}, err
			}
		}
		a.GroupData()
	}
	return a._dataStageIntroDescID[identifier], nil
}
