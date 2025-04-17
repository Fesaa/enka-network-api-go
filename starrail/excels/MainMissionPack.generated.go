package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type MainMissionPack struct {
	MainMissionIdList []json.Number `json:"MainMissionIdList"`
	MissionPack       json.Number   `json:"MissionPack"`
}
type MainMissionPackAccessor struct {
	_data            []MainMissionPack
	_dataMissionPack map[json.Number]MainMissionPack
}

// LoadData retrieves the data. Must be called before MainMissionPack.GroupData
func (a *MainMissionPackAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MainMissionPack.json")
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
func (a *MainMissionPackAccessor) Raw() ([]MainMissionPack, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MainMissionPack{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MainMissionPackAccessor.LoadData to preload everything
func (a *MainMissionPackAccessor) GroupData() {
	a._dataMissionPack = map[json.Number]MainMissionPack{}
	for _, d := range a._data {
		a._dataMissionPack[d.MissionPack] = d
	}
}

// ByMissionPack returns the MainMissionPack uniquely identified by MissionPack
//
// Error is only non-nil if the source errors out
func (a *MainMissionPackAccessor) ByMissionPack(identifier json.Number) (MainMissionPack, error) {
	if a._dataMissionPack == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MainMissionPack{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMissionPack[identifier], nil
}
