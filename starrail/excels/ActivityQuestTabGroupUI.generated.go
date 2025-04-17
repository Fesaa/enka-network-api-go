package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ActivityQuestTabGroupUI struct {
	BgPrefabPath      string  `json:"BgPrefabPath"`
	QuestTabGroupID   float64 `json:"QuestTabGroupID"`
	TabItemPrefabPath string  `json:"TabItemPrefabPath"`
}
type ActivityQuestTabGroupUIAccessor struct {
	_data                []ActivityQuestTabGroupUI
	_dataBgPrefabPath    map[string]ActivityQuestTabGroupUI
	_dataQuestTabGroupID map[float64]ActivityQuestTabGroupUI
}

// LoadData retrieves the data. Must be called before ActivityQuestTabGroupUI.GroupData
func (a *ActivityQuestTabGroupUIAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActivityQuestTabGroupUI.json")
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
func (a *ActivityQuestTabGroupUIAccessor) Raw() ([]ActivityQuestTabGroupUI, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActivityQuestTabGroupUI{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ActivityQuestTabGroupUIAccessor.LoadData to preload everything
func (a *ActivityQuestTabGroupUIAccessor) GroupData() {
	a._dataBgPrefabPath = map[string]ActivityQuestTabGroupUI{}
	a._dataQuestTabGroupID = map[float64]ActivityQuestTabGroupUI{}
	for _, d := range a._data {
		a._dataBgPrefabPath[d.BgPrefabPath] = d
		a._dataQuestTabGroupID[d.QuestTabGroupID] = d
	}
}

// ByBgPrefabPath returns the ActivityQuestTabGroupUI uniquely identified by BgPrefabPath
//
// Error is only non-nil if the source errors out
func (a *ActivityQuestTabGroupUIAccessor) ByBgPrefabPath(identifier string) (ActivityQuestTabGroupUI, error) {
	if a._dataBgPrefabPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityQuestTabGroupUI{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBgPrefabPath[identifier], nil
}

// ByQuestTabGroupID returns the ActivityQuestTabGroupUI uniquely identified by QuestTabGroupID
//
// Error is only non-nil if the source errors out
func (a *ActivityQuestTabGroupUIAccessor) ByQuestTabGroupID(identifier float64) (ActivityQuestTabGroupUI, error) {
	if a._dataQuestTabGroupID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityQuestTabGroupUI{}, err
			}
		}
		a.GroupData()
	}
	return a._dataQuestTabGroupID[identifier], nil
}
