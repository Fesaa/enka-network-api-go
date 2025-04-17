package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type ActivitySummonLevel struct {
	BattleTargetList        []json.Number `json:"BattleTargetList"`
	DifficultyLevel         string        `json:"DifficultyLevel"`
	EventID                 json.Number   `json:"EventID"`
	GroupID                 json.Number   `json:"GroupID"`
	ImagePath               string        `json:"ImagePath"`
	MasterAvatarList        []json.Number `json:"MasterAvatarList"`
	ReplaceMasterAvatarList []interface{} `json:"ReplaceMasterAvatarList"`
	ReplaceTrialAvatarList  []json.Number `json:"ReplaceTrialAvatarList"`
	TrialAvatarList         []json.Number `json:"TrialAvatarList"`
	UIEnterBattleAreaID     json.Number   `json:"UIEnterBattleAreaID"`
}
type ActivitySummonLevelAccessor struct {
	_data        []ActivitySummonLevel
	_dataEventID map[json.Number]ActivitySummonLevel
}

// LoadData retrieves the data. Must be called before ActivitySummonLevel.GroupData
func (a *ActivitySummonLevelAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActivitySummonLevel.json")
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
func (a *ActivitySummonLevelAccessor) Raw() ([]ActivitySummonLevel, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActivitySummonLevel{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ActivitySummonLevelAccessor.LoadData to preload everything
func (a *ActivitySummonLevelAccessor) GroupData() {
	a._dataEventID = map[json.Number]ActivitySummonLevel{}
	for _, d := range a._data {
		a._dataEventID[d.EventID] = d
	}
}

// ByEventID returns the ActivitySummonLevel uniquely identified by EventID
//
// Error is only non-nil if the source errors out
func (a *ActivitySummonLevelAccessor) ByEventID(identifier json.Number) (ActivitySummonLevel, error) {
	if a._dataEventID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivitySummonLevel{}, err
			}
		}
		a.GroupData()
	}
	return a._dataEventID[identifier], nil
}
