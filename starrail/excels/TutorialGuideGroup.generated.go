package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type TutorialGuideGroup struct {
	CanReview           bool                                    `json:"CanReview"`
	FinishTriggerParams []TutorialGuideGroupFinishTriggerParams `json:"FinishTriggerParams"`
	GroupID             json.Number                             `json:"GroupID"`
	MessageText         map[string]json.Number                  `json:"MessageText"`
	Order               json.Number                             `json:"Order"`
	RewardID            json.Number                             `json:"RewardID"`
	TriggerParams       []TutorialGuideGroupTriggerParams       `json:"TriggerParams"`
	TutorialGuideIDList []json.Number                           `json:"TutorialGuideIDList"`
	TutorialShowType    string                                  `json:"TutorialShowType"`
	TutorialType        json.Number                             `json:"TutorialType"`
}
type TutorialGuideGroupFinishTriggerParams struct {
	TriggerParam string `json:"TriggerParam"`
	TriggerType  string `json:"TriggerType"`
}
type TutorialGuideGroupTriggerParams struct {
	TriggerParam string `json:"TriggerParam"`
	TriggerType  string `json:"TriggerType"`
}
type TutorialGuideGroupAccessor struct {
	_data        []TutorialGuideGroup
	_dataGroupID map[json.Number]TutorialGuideGroup
}

// LoadData retrieves the data. Must be called before TutorialGuideGroup.GroupData
func (a *TutorialGuideGroupAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TutorialGuideGroup.json")
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
func (a *TutorialGuideGroupAccessor) Raw() ([]TutorialGuideGroup, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TutorialGuideGroup{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with TutorialGuideGroupAccessor.LoadData to preload everything
func (a *TutorialGuideGroupAccessor) GroupData() {
	a._dataGroupID = map[json.Number]TutorialGuideGroup{}
	for _, d := range a._data {
		a._dataGroupID[d.GroupID] = d
	}
}

// ByGroupID returns the TutorialGuideGroup uniquely identified by GroupID
//
// Error is only non-nil if the source errors out
func (a *TutorialGuideGroupAccessor) ByGroupID(identifier json.Number) (TutorialGuideGroup, error) {
	if a._dataGroupID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TutorialGuideGroup{}, err
			}
		}
		a.GroupData()
	}
	return a._dataGroupID[identifier], nil
}
