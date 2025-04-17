package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type StageInfiniteWaveConfig struct {
	Ability              string                   `json:"Ability"`
	ClearPreviousAbility bool                     `json:"ClearPreviousAbility"`
	InfiniteWaveID       json.Number              `json:"InfiniteWaveID"`
	MaxMonsterCount      json.Number              `json:"MaxMonsterCount"`
	MaxTeammateCount     json.Number              `json:"MaxTeammateCount"`
	MonsterGroupIDList   []json.Number            `json:"MonsterGroupIDList"`
	ParamList            []map[string]json.Number `json:"ParamList"`
}
type StageInfiniteWaveConfigAccessor struct {
	_data               []StageInfiniteWaveConfig
	_dataInfiniteWaveID map[json.Number]StageInfiniteWaveConfig
}

// LoadData retrieves the data. Must be called before StageInfiniteWaveConfig.GroupData
func (a *StageInfiniteWaveConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/StageInfiniteWaveConfig.json")
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
func (a *StageInfiniteWaveConfigAccessor) Raw() ([]StageInfiniteWaveConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []StageInfiniteWaveConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with StageInfiniteWaveConfigAccessor.LoadData to preload everything
func (a *StageInfiniteWaveConfigAccessor) GroupData() {
	a._dataInfiniteWaveID = map[json.Number]StageInfiniteWaveConfig{}
	for _, d := range a._data {
		a._dataInfiniteWaveID[d.InfiniteWaveID] = d
	}
}

// ByInfiniteWaveID returns the StageInfiniteWaveConfig uniquely identified by InfiniteWaveID
//
// Error is only non-nil if the source errors out
func (a *StageInfiniteWaveConfigAccessor) ByInfiniteWaveID(identifier json.Number) (StageInfiniteWaveConfig, error) {
	if a._dataInfiniteWaveID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return StageInfiniteWaveConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataInfiniteWaveID[identifier], nil
}
