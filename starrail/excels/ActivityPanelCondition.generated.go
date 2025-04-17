package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type ActivityPanelCondition struct {
	ActivityGoto                 json.Number                             `json:"ActivityGoto"`
	ActivityGotoStoryLineRestore bool                                    `json:"ActivityGotoStoryLineRestore"`
	ActivityOpenActivityModule   json.Number                             `json:"ActivityOpenActivityModule"`
	GuideConditions              []ActivityPanelConditionGuideConditions `json:"GuideConditions"`
	GuideGoto                    json.Number                             `json:"GuideGoto"`
	GuideTakeMission             json.Number                             `json:"GuideTakeMission"`
	PanelID                      json.Number                             `json:"PanelID"`
	PreConditions                []ActivityPanelConditionPreConditions   `json:"PreConditions"`
	ShopOnlyActivityModule       json.Number                             `json:"ShopOnlyActivityModule"`
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
	_data []ActivityPanelCondition
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
