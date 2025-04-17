package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type MainMission struct {
	BeginOperation        string                  `json:"BeginOperation"`
	BeginParam            []MainMissionBeginParam `json:"BeginParam"`
	ChapterID             json.Number             `json:"ChapterID"`
	DisplayPriority       json.Number             `json:"DisplayPriority"`
	DisplayRewardID       json.Number             `json:"DisplayRewardID"`
	IsDisplayActivityIcon bool                    `json:"IsDisplayActivityIcon"`
	IsInRaid              bool                    `json:"IsInRaid"`
	MainMissionID         json.Number             `json:"MainMissionID"`
	MissionAdvance        json.Number             `json:"MissionAdvance"`
	MissionPack           json.Number             `json:"MissionPack"`
	Name                  map[string]json.Number  `json:"Name"`
	NextMainMissionList   []interface{}           `json:"NextMainMissionList"`
	NextTrackMainMission  json.Number             `json:"NextTrackMainMission"`
	RewardID              json.Number             `json:"RewardID"`
	SubRewardList         []json.Number           `json:"SubRewardList"`
	SubType               string                  `json:"SubType"`
	TakeOperation         string                  `json:"TakeOperation"`
	TakeParam             []MainMissionTakeParam  `json:"TakeParam"`
	TrackWeight           json.Number             `json:"TrackWeight"`
	Type                  string                  `json:"Type"`
	WorldID               json.Number             `json:"WorldID"`
}
type MainMissionBeginParam struct {
	Type  string      `json:"Type"`
	Value json.Number `json:"Value"`
}
type MainMissionTakeParam struct {
	Type  string      `json:"Type"`
	Value json.Number `json:"Value"`
}
type MainMissionAccessor struct {
	_data                []MainMission
	_dataDisplayPriority map[json.Number]MainMission
	_dataMainMissionID   map[json.Number]MainMission
}

// LoadData retrieves the data. Must be called before MainMission.GroupData
func (a *MainMissionAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MainMission.json")
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
func (a *MainMissionAccessor) Raw() ([]MainMission, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MainMission{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MainMissionAccessor.LoadData to preload everything
func (a *MainMissionAccessor) GroupData() {
	a._dataDisplayPriority = map[json.Number]MainMission{}
	a._dataMainMissionID = map[json.Number]MainMission{}
	for _, d := range a._data {
		a._dataDisplayPriority[d.DisplayPriority] = d
		a._dataMainMissionID[d.MainMissionID] = d
	}
}

// ByDisplayPriority returns the MainMission uniquely identified by DisplayPriority
//
// Error is only non-nil if the source errors out
func (a *MainMissionAccessor) ByDisplayPriority(identifier json.Number) (MainMission, error) {
	if a._dataDisplayPriority == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MainMission{}, err
			}
		}
		a.GroupData()
	}
	return a._dataDisplayPriority[identifier], nil
}

// ByMainMissionID returns the MainMission uniquely identified by MainMissionID
//
// Error is only non-nil if the source errors out
func (a *MainMissionAccessor) ByMainMissionID(identifier json.Number) (MainMission, error) {
	if a._dataMainMissionID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MainMission{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMainMissionID[identifier], nil
}
