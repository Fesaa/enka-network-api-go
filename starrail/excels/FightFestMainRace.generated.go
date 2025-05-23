package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type FightFestMainRace struct {
	BlueAvatarID     float64   `json:"BlueAvatarID"`
	EventID          float64   `json:"EventID"`
	FightPhaseID     float64   `json:"FightPhaseID"`
	MainRaceID       float64   `json:"MainRaceID"`
	RaceBgFigurePath string    `json:"RaceBgFigurePath"`
	RedAvatarID      float64   `json:"RedAvatarID"`
	RewardID         float64   `json:"RewardID"`
	StageEndDesc     hash.Hash `json:"StageEndDesc"`
	StageName        hash.Hash `json:"StageName"`
	TutorialID       float64   `json:"TutorialID"`
}
type FightFestMainRaceAccessor struct {
	_data            []FightFestMainRace
	_dataEventID     map[float64]FightFestMainRace
	_dataMainRaceID  map[float64]FightFestMainRace
	_dataRedAvatarID map[float64]FightFestMainRace
	_dataRewardID    map[float64]FightFestMainRace
	_dataTutorialID  map[float64]FightFestMainRace
}

// LoadData retrieves the data. Must be called before FightFestMainRace.GroupData
func (a *FightFestMainRaceAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/FightFestMainRace.json")
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
func (a *FightFestMainRaceAccessor) Raw() ([]FightFestMainRace, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []FightFestMainRace{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with FightFestMainRaceAccessor.LoadData to preload everything
func (a *FightFestMainRaceAccessor) GroupData() {
	a._dataEventID = map[float64]FightFestMainRace{}
	a._dataMainRaceID = map[float64]FightFestMainRace{}
	a._dataRedAvatarID = map[float64]FightFestMainRace{}
	a._dataRewardID = map[float64]FightFestMainRace{}
	a._dataTutorialID = map[float64]FightFestMainRace{}
	for _, d := range a._data {
		a._dataEventID[d.EventID] = d
		a._dataMainRaceID[d.MainRaceID] = d
		a._dataRedAvatarID[d.RedAvatarID] = d
		a._dataRewardID[d.RewardID] = d
		a._dataTutorialID[d.TutorialID] = d
	}
}

// ByEventID returns the FightFestMainRace uniquely identified by EventID
//
// Error is only non-nil if the source errors out
func (a *FightFestMainRaceAccessor) ByEventID(identifier float64) (FightFestMainRace, error) {
	if a._dataEventID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return FightFestMainRace{}, err
			}
		}
		a.GroupData()
	}
	return a._dataEventID[identifier], nil
}

// ByMainRaceID returns the FightFestMainRace uniquely identified by MainRaceID
//
// Error is only non-nil if the source errors out
func (a *FightFestMainRaceAccessor) ByMainRaceID(identifier float64) (FightFestMainRace, error) {
	if a._dataMainRaceID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return FightFestMainRace{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMainRaceID[identifier], nil
}

// ByRedAvatarID returns the FightFestMainRace uniquely identified by RedAvatarID
//
// Error is only non-nil if the source errors out
func (a *FightFestMainRaceAccessor) ByRedAvatarID(identifier float64) (FightFestMainRace, error) {
	if a._dataRedAvatarID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return FightFestMainRace{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRedAvatarID[identifier], nil
}

// ByRewardID returns the FightFestMainRace uniquely identified by RewardID
//
// Error is only non-nil if the source errors out
func (a *FightFestMainRaceAccessor) ByRewardID(identifier float64) (FightFestMainRace, error) {
	if a._dataRewardID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return FightFestMainRace{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRewardID[identifier], nil
}

// ByTutorialID returns the FightFestMainRace uniquely identified by TutorialID
//
// Error is only non-nil if the source errors out
func (a *FightFestMainRaceAccessor) ByTutorialID(identifier float64) (FightFestMainRace, error) {
	if a._dataTutorialID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return FightFestMainRace{}, err
			}
		}
		a.GroupData()
	}
	return a._dataTutorialID[identifier], nil
}
