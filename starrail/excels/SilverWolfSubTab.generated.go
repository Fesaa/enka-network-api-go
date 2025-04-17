package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type SilverWolfSubTab struct {
	EntranceID    float64   `json:"EntranceID"`
	FinalQuest    float64   `json:"FinalQuest"`
	GroupID       float64   `json:"GroupID"`
	MappingInfoID float64   `json:"MappingInfoID"`
	QuestList     []float64 `json:"QuestList"`
	TabType       string    `json:"TabType"`
	UnlockMission float64   `json:"UnlockMission"`
}
type SilverWolfSubTabAccessor struct {
	_data           []SilverWolfSubTab
	_dataFinalQuest map[float64]SilverWolfSubTab
}

// LoadData retrieves the data. Must be called before SilverWolfSubTab.GroupData
func (a *SilverWolfSubTabAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/SilverWolfSubTab.json")
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
func (a *SilverWolfSubTabAccessor) Raw() ([]SilverWolfSubTab, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []SilverWolfSubTab{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with SilverWolfSubTabAccessor.LoadData to preload everything
func (a *SilverWolfSubTabAccessor) GroupData() {
	a._dataFinalQuest = map[float64]SilverWolfSubTab{}
	for _, d := range a._data {
		a._dataFinalQuest[d.FinalQuest] = d
	}
}

// ByFinalQuest returns the SilverWolfSubTab uniquely identified by FinalQuest
//
// Error is only non-nil if the source errors out
func (a *SilverWolfSubTabAccessor) ByFinalQuest(identifier float64) (SilverWolfSubTab, error) {
	if a._dataFinalQuest == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return SilverWolfSubTab{}, err
			}
		}
		a.GroupData()
	}
	return a._dataFinalQuest[identifier], nil
}
