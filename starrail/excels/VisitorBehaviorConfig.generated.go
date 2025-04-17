package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type VisitorBehaviorConfig struct {
	AnchorID                     float64 `json:"AnchorID"`
	BehaviorID                   float64 `json:"BehaviorID"`
	DefaultIdleFreeStyleMotionID float64 `json:"DefaultIdleFreeStyleMotionID"`
	DefaultPerformanceID         float64 `json:"DefaultPerformanceID"`
	NPCGroupID                   float64 `json:"NPCGroupID"`
	NPCID                        float64 `json:"NPCID"`
	NPCRotationYInfo             float64 `json:"NPCRotationYInfo"`
	NpcBubbleTalkSentenceID      float64 `json:"NpcBubbleTalkSentenceID"`
	PerformanceID                float64 `json:"PerformanceID"`
	RewardID                     float64 `json:"RewardID"`
	VisitorID                    float64 `json:"VisitorID"`
}
type VisitorBehaviorConfigAccessor struct {
	_data []VisitorBehaviorConfig
}

// LoadData retrieves the data. Must be called before VisitorBehaviorConfig.GroupData
func (a *VisitorBehaviorConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/VisitorBehaviorConfig.json")
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
func (a *VisitorBehaviorConfigAccessor) Raw() ([]VisitorBehaviorConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []VisitorBehaviorConfig{}, err
		}
	}
	return a._data, nil
}
