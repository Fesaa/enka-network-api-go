package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type TutorialGuideGroup struct {
	CanReview           bool                                    `json:"CanReview"`
	FinishTriggerParams []TutorialGuideGroupFinishTriggerParams `json:"FinishTriggerParams"`
	GroupID             float64                                 `json:"GroupID"`
	MessageText         hash.Hash                               `json:"MessageText"`
	Order               float64                                 `json:"Order"`
	RewardID            float64                                 `json:"RewardID"`
	TriggerParams       []TutorialGuideGroupTriggerParams       `json:"TriggerParams"`
	TutorialGuideIDList []float64                               `json:"TutorialGuideIDList"`
	TutorialShowType    string                                  `json:"TutorialShowType"`
	TutorialType        float64                                 `json:"TutorialType"`
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
	_dataGroupID map[float64]TutorialGuideGroup
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
	for _, d := range a._data {
		a._dataGroupID[d.GroupID] = d
	}
}

// ByGroupID returns the TutorialGuideGroup uniquely identified by GroupID
//
// Error is only non-nil if the source errors out
func (a *TutorialGuideGroupAccessor) ByGroupID(identifier float64) (TutorialGuideGroup, error) {
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
