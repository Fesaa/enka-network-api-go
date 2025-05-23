package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type ParkourRailBallSkill struct {
	Desc           hash.Hash `json:"Desc"`
	ID             float64   `json:"ID"`
	IconPath       string    `json:"IconPath"`
	MiniIconBGPath string    `json:"MiniIconBGPath"`
	MiniIconPath   string    `json:"MiniIconPath"`
	Name           hash.Hash `json:"Name"`
	TutorialID     float64   `json:"TutorialID"`
	VideoID        float64   `json:"VideoID"`
}
type ParkourRailBallSkillAccessor struct {
	_data   []ParkourRailBallSkill
	_dataID map[float64]ParkourRailBallSkill
}

// LoadData retrieves the data. Must be called before ParkourRailBallSkill.GroupData
func (a *ParkourRailBallSkillAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ParkourRailBallSkill.json")
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
func (a *ParkourRailBallSkillAccessor) Raw() ([]ParkourRailBallSkill, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ParkourRailBallSkill{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ParkourRailBallSkillAccessor.LoadData to preload everything
func (a *ParkourRailBallSkillAccessor) GroupData() {
	a._dataID = map[float64]ParkourRailBallSkill{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the ParkourRailBallSkill uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *ParkourRailBallSkillAccessor) ByID(identifier float64) (ParkourRailBallSkill, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ParkourRailBallSkill{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
