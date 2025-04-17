package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type MazeFloor struct {
	BGMWorldState              string      `json:"BGMWorldState"`
	BaseFloorID                float64     `json:"BaseFloorID"`
	CombatBGMHigh              string      `json:"CombatBGMHigh"`
	CombatBGMLow               string      `json:"CombatBGMLow"`
	EnterAudioEvent            []string    `json:"EnterAudioEvent"`
	ExitAudioEvent             []string    `json:"ExitAudioEvent"`
	FloorBGMBusyStateName      string      `json:"FloorBGMBusyStateName"`
	FloorBGMGroupName          string      `json:"FloorBGMGroupName"`
	FloorBGMNormalStateName    string      `json:"FloorBGMNormalStateName"`
	FloorDefaultEmotion        string      `json:"FloorDefaultEmotion"`
	FloorID                    float64     `json:"FloorID"`
	FloorName                  string      `json:"FloorName"`
	FloorTag                   []string    `json:"FloorTag"`
	FloorType                  string      `json:"FloorType"`
	MapLayerNameList           []hash.Hash `json:"MapLayerNameList"`
	MunicipalConfigPath        string      `json:"MunicipalConfigPath"`
	OptionalLoadBlocksConfig   string      `json:"OptionalLoadBlocksConfig"`
	WalkingEffectAdditiveScale float64     `json:"WalkingEffectAdditiveScale"`
}
type MazeFloorAccessor struct {
	_data        []MazeFloor
	_dataFloorID map[float64]MazeFloor
}

// LoadData retrieves the data. Must be called before MazeFloor.GroupData
func (a *MazeFloorAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MazeFloor.json")
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
func (a *MazeFloorAccessor) Raw() ([]MazeFloor, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MazeFloor{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MazeFloorAccessor.LoadData to preload everything
func (a *MazeFloorAccessor) GroupData() {
	for _, d := range a._data {
		a._dataFloorID[d.FloorID] = d
	}
}

// ByFloorID returns the MazeFloor uniquely identified by FloorID
//
// Error is only non-nil if the source errors out
func (a *MazeFloorAccessor) ByFloorID(identifier float64) (MazeFloor, error) {
	if a._dataFloorID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MazeFloor{}, err
			}
		}
		a.GroupData()
	}
	return a._dataFloorID[identifier], nil
}
