package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ActivityPanelCondition struct {
	ActivityGoto                 float64                                 `json:"ActivityGoto"`
	ActivityGotoStoryLineRestore bool                                    `json:"ActivityGotoStoryLineRestore"`
	ActivityOpenActivityModule   float64                                 `json:"ActivityOpenActivityModule"`
	GuideConditions              []ActivityPanelConditionGuideConditions `json:"GuideConditions"`
	GuideGoto                    float64                                 `json:"GuideGoto"`
	GuideTakeMission             float64                                 `json:"GuideTakeMission"`
	PanelID                      float64                                 `json:"PanelID"`
	PreConditions                []ActivityPanelConditionPreConditions   `json:"PreConditions"`
	ShopOnlyActivityModule       float64                                 `json:"ShopOnlyActivityModule"`
}
type ActivityPanelConditionGuideConditions struct {
	Param string `json:"Param"`
	Type  string `json:"Type"`
}
type ActivityPanelConditionPreConditions struct {
	Param string `json:"Param"`
	Type  string `json:"Type"`
}
type ActivityPanelConditionAccessor struct {
	_data        []ActivityPanelCondition
	_dataPanelID map[float64]ActivityPanelCondition
}

// LoadData retrieves the data. Must be called before ActivityPanelCondition.GroupData
func (a *ActivityPanelConditionAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActivityPanelCondition.json")
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
func (a *ActivityPanelConditionAccessor) Raw() ([]ActivityPanelCondition, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActivityPanelCondition{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ActivityPanelConditionAccessor.LoadData to preload everything
func (a *ActivityPanelConditionAccessor) GroupData() {
	a._dataPanelID = map[float64]ActivityPanelCondition{}
	for _, d := range a._data {
		a._dataPanelID[d.PanelID] = d
	}
}

// ByPanelID returns the ActivityPanelCondition uniquely identified by PanelID
//
// Error is only non-nil if the source errors out
func (a *ActivityPanelConditionAccessor) ByPanelID(identifier float64) (ActivityPanelCondition, error) {
	if a._dataPanelID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityPanelCondition{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPanelID[identifier], nil
}
