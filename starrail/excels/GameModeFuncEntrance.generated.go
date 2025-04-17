package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type GameModeFuncEntrance struct {
	BranchLineFuncEntranceListID float64 `json:"BranchLineFuncEntranceListID"`
	GameModeType                 float64 `json:"GameModeType"`
	MainLineFuncEntranceListID   float64 `json:"MainLineFuncEntranceListID"`
}
type GameModeFuncEntranceAccessor struct {
	_data             []GameModeFuncEntrance
	_dataGameModeType map[float64]GameModeFuncEntrance
}

// LoadData retrieves the data. Must be called before GameModeFuncEntrance.GroupData
func (a *GameModeFuncEntranceAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/GameModeFuncEntrance.json")
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
func (a *GameModeFuncEntranceAccessor) Raw() ([]GameModeFuncEntrance, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []GameModeFuncEntrance{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with GameModeFuncEntranceAccessor.LoadData to preload everything
func (a *GameModeFuncEntranceAccessor) GroupData() {
	a._dataGameModeType = map[float64]GameModeFuncEntrance{}
	for _, d := range a._data {
		a._dataGameModeType[d.GameModeType] = d
	}
}

// ByGameModeType returns the GameModeFuncEntrance uniquely identified by GameModeType
//
// Error is only non-nil if the source errors out
func (a *GameModeFuncEntranceAccessor) ByGameModeType(identifier float64) (GameModeFuncEntrance, error) {
	if a._dataGameModeType == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return GameModeFuncEntrance{}, err
			}
		}
		a.GroupData()
	}
	return a._dataGameModeType[identifier], nil
}
