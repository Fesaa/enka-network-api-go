package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type StarFightGoal struct {
	BattleTargetID float64 `json:"BattleTargetID"`
	ID             float64 `json:"ID"`
}
type StarFightGoalAccessor struct {
	_data   []StarFightGoal
	_dataID map[float64]StarFightGoal
}

// LoadData retrieves the data. Must be called before StarFightGoal.GroupData
func (a *StarFightGoalAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/StarFightGoal.json")
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
func (a *StarFightGoalAccessor) Raw() ([]StarFightGoal, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []StarFightGoal{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with StarFightGoalAccessor.LoadData to preload everything
func (a *StarFightGoalAccessor) GroupData() {
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the StarFightGoal uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *StarFightGoalAccessor) ByID(identifier float64) (StarFightGoal, error) {
	if a._dataID == nil {
		err := a.LoadData()
		if err != nil {
			return StarFightGoal{}, err
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
