package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type BoxingClubPerformance struct {
	BubbleTalkEnemy   string      `json:"BubbleTalkEnemy"`
	BubbleTalkPlayer  string      `json:"BubbleTalkPlayer"`
	EnemyRank         string      `json:"EnemyRank"`
	ID                json.Number `json:"ID"`
	MonsterTemplateID json.Number `json:"MonsterTemplateID"`
	Name              string      `json:"Name"`
	PlayerRank        string      `json:"PlayerRank"`
}
type BoxingClubPerformanceAccessor struct {
	_data                 []BoxingClubPerformance
	_dataBubbleTalkEnemy  map[string]BoxingClubPerformance
	_dataBubbleTalkPlayer map[string]BoxingClubPerformance
	_dataName             map[string]BoxingClubPerformance
}

// LoadData retrieves the data. Must be called before BoxingClubPerformance.GroupData
func (a *BoxingClubPerformanceAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/BoxingClubPerformance.json")
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
func (a *BoxingClubPerformanceAccessor) Raw() ([]BoxingClubPerformance, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []BoxingClubPerformance{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with BoxingClubPerformanceAccessor.LoadData to preload everything
func (a *BoxingClubPerformanceAccessor) GroupData() {
	a._dataBubbleTalkEnemy = map[string]BoxingClubPerformance{}
	a._dataBubbleTalkPlayer = map[string]BoxingClubPerformance{}
	a._dataName = map[string]BoxingClubPerformance{}
	for _, d := range a._data {
		a._dataBubbleTalkEnemy[d.BubbleTalkEnemy] = d
		a._dataBubbleTalkPlayer[d.BubbleTalkPlayer] = d
		a._dataName[d.Name] = d
	}
}

// ByBubbleTalkEnemy returns the BoxingClubPerformance uniquely identified by BubbleTalkEnemy
//
// Error is only non-nil if the source errors out
func (a *BoxingClubPerformanceAccessor) ByBubbleTalkEnemy(identifier string) (BoxingClubPerformance, error) {
	if a._dataBubbleTalkEnemy == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return BoxingClubPerformance{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBubbleTalkEnemy[identifier], nil
}

// ByBubbleTalkPlayer returns the BoxingClubPerformance uniquely identified by BubbleTalkPlayer
//
// Error is only non-nil if the source errors out
func (a *BoxingClubPerformanceAccessor) ByBubbleTalkPlayer(identifier string) (BoxingClubPerformance, error) {
	if a._dataBubbleTalkPlayer == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return BoxingClubPerformance{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBubbleTalkPlayer[identifier], nil
}

// ByName returns the BoxingClubPerformance uniquely identified by Name
//
// Error is only non-nil if the source errors out
func (a *BoxingClubPerformanceAccessor) ByName(identifier string) (BoxingClubPerformance, error) {
	if a._dataName == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return BoxingClubPerformance{}, err
			}
		}
		a.GroupData()
	}
	return a._dataName[identifier], nil
}
