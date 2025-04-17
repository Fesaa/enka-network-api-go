package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type MarbleMatchInfo struct {
	AIRank               json.Number            `json:"AIRank"`
	ANpcIds              []json.Number          `json:"ANpcIds"`
	BNpcIds              []json.Number          `json:"BNpcIds"`
	BanSealList          []json.Number          `json:"BanSealList"`
	CanGoMatchSubMission json.Number            `json:"CanGoMatchSubMission"`
	CustomID             json.Number            `json:"CustomID"`
	FirstType            json.Number            `json:"FirstType"`
	ID                   json.Number            `json:"ID"`
	LevelID              json.Number            `json:"LevelID"`
	Name                 map[string]json.Number `json:"Name"`
	PerformanceID        json.Number            `json:"PerformanceID"`
	PhaseID              string                 `json:"PhaseID"`
	PlayerID             json.Number            `json:"PlayerID"`
	Reward               json.Number            `json:"Reward"`
	Round                json.Number            `json:"Round"`
}
type MarbleMatchInfoAccessor struct {
	_data         []MarbleMatchInfo
	_dataID       map[json.Number]MarbleMatchInfo
	_dataLevelID  map[json.Number]MarbleMatchInfo
	_dataPlayerID map[json.Number]MarbleMatchInfo
}

// LoadData retrieves the data. Must be called before MarbleMatchInfo.GroupData
func (a *MarbleMatchInfoAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MarbleMatchInfo.json")
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
func (a *MarbleMatchInfoAccessor) Raw() ([]MarbleMatchInfo, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MarbleMatchInfo{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MarbleMatchInfoAccessor.LoadData to preload everything
func (a *MarbleMatchInfoAccessor) GroupData() {
	a._dataID = map[json.Number]MarbleMatchInfo{}
	a._dataLevelID = map[json.Number]MarbleMatchInfo{}
	a._dataPlayerID = map[json.Number]MarbleMatchInfo{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
		a._dataLevelID[d.LevelID] = d
		a._dataPlayerID[d.PlayerID] = d
	}
}

// ByID returns the MarbleMatchInfo uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *MarbleMatchInfoAccessor) ByID(identifier json.Number) (MarbleMatchInfo, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MarbleMatchInfo{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByLevelID returns the MarbleMatchInfo uniquely identified by LevelID
//
// Error is only non-nil if the source errors out
func (a *MarbleMatchInfoAccessor) ByLevelID(identifier json.Number) (MarbleMatchInfo, error) {
	if a._dataLevelID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MarbleMatchInfo{}, err
			}
		}
		a.GroupData()
	}
	return a._dataLevelID[identifier], nil
}

// ByPlayerID returns the MarbleMatchInfo uniquely identified by PlayerID
//
// Error is only non-nil if the source errors out
func (a *MarbleMatchInfoAccessor) ByPlayerID(identifier json.Number) (MarbleMatchInfo, error) {
	if a._dataPlayerID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MarbleMatchInfo{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPlayerID[identifier], nil
}
