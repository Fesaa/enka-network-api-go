package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type NPCMonsterMark struct {
	GroupID    float64 `json:"GroupID"`
	ID         float64 `json:"ID"`
	InstanceID float64 `json:"InstanceID"`
}
type NPCMonsterMarkAccessor struct {
	_data   []NPCMonsterMark
	_dataID map[float64]NPCMonsterMark
}

// LoadData retrieves the data. Must be called before NPCMonsterMark.GroupData
func (a *NPCMonsterMarkAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/NPCMonsterMark.json")
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
func (a *NPCMonsterMarkAccessor) Raw() ([]NPCMonsterMark, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []NPCMonsterMark{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with NPCMonsterMarkAccessor.LoadData to preload everything
func (a *NPCMonsterMarkAccessor) GroupData() {
	a._dataID = map[float64]NPCMonsterMark{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the NPCMonsterMark uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *NPCMonsterMarkAccessor) ByID(identifier float64) (NPCMonsterMark, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return NPCMonsterMark{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
