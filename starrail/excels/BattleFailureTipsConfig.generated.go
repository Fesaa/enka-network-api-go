package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type BattleFailureTipsConfig struct {
	BattleFailureTipID       float64       `json:"BattleFailureTipID"`
	CustomStringList         []string      `json:"CustomStringList"`
	GameModeList             []string      `json:"GameModeList"`
	MainMissionFinishForce   []interface{} `json:"MainMissionFinishForce"`
	MainMissionTakenForce    []interface{} `json:"MainMissionTakenForce"`
	MainMissionUnfinishForce []interface{} `json:"MainMissionUnfinishForce"`
	MazebuffIDList           []interface{} `json:"MazebuffIDList"`
	MonsterTemplateIDList    []float64     `json:"MonsterTemplateIDList"`
	PlayerLevel              []float64     `json:"PlayerLevel"`
	Priority                 float64       `json:"Priority"`
	StageIDForce             []float64     `json:"StageIDForce"`
	StageTypeForce           []string      `json:"StageTypeForce"`
	TipContent               hash.Hash     `json:"TipContent"`
	Type                     string        `json:"Type"`
	WorldList                []interface{} `json:"WorldList"`
}
type BattleFailureTipsConfigAccessor struct {
	_data                   []BattleFailureTipsConfig
	_dataBattleFailureTipID map[float64]BattleFailureTipsConfig
}

// LoadData retrieves the data. Must be called before BattleFailureTipsConfig.GroupData
func (a *BattleFailureTipsConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/BattleFailureTipsConfig.json")
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
func (a *BattleFailureTipsConfigAccessor) Raw() ([]BattleFailureTipsConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []BattleFailureTipsConfig{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with BattleFailureTipsConfigAccessor.LoadData to preload everything
func (a *BattleFailureTipsConfigAccessor) GroupData() {
	for _, d := range a._data {
		a._dataBattleFailureTipID[d.BattleFailureTipID] = d
	}
}

// ByBattleFailureTipID returns the BattleFailureTipsConfig uniquely identified by BattleFailureTipID
//
// Error is only non-nil if the source errors out
func (a *BattleFailureTipsConfigAccessor) ByBattleFailureTipID(identifier float64) (BattleFailureTipsConfig, error) {
	if a._dataBattleFailureTipID == nil {
		err := a.LoadData()
		if err != nil {
			return BattleFailureTipsConfig{}, err
		}
		a.GroupData()
	}
	return a._dataBattleFailureTipID[identifier], nil
}
