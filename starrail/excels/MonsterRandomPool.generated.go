package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type MonsterRandomPool struct {
	ElitePool    []json.Number `json:"ElitePool"`
	MinionPool   []json.Number `json:"MinionPool"`
	RandomPoolID json.Number   `json:"RandomPoolID"`
}
type MonsterRandomPoolAccessor struct {
	_data             []MonsterRandomPool
	_dataRandomPoolID map[json.Number]MonsterRandomPool
}

// LoadData retrieves the data. Must be called before MonsterRandomPool.GroupData
func (a *MonsterRandomPoolAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MonsterRandomPool.json")
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
func (a *MonsterRandomPoolAccessor) Raw() ([]MonsterRandomPool, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MonsterRandomPool{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MonsterRandomPoolAccessor.LoadData to preload everything
func (a *MonsterRandomPoolAccessor) GroupData() {
	a._dataRandomPoolID = map[json.Number]MonsterRandomPool{}
	for _, d := range a._data {
		a._dataRandomPoolID[d.RandomPoolID] = d
	}
}

// ByRandomPoolID returns the MonsterRandomPool uniquely identified by RandomPoolID
//
// Error is only non-nil if the source errors out
func (a *MonsterRandomPoolAccessor) ByRandomPoolID(identifier json.Number) (MonsterRandomPool, error) {
	if a._dataRandomPoolID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MonsterRandomPool{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRandomPoolID[identifier], nil
}
