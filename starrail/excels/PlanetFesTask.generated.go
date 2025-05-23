package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type PlanetFesTask struct {
	QuestID       float64   `json:"QuestID"`
	RandomGroupID float64   `json:"RandomGroupID"`
	TaskID        float64   `json:"TaskID"`
	TaskTips      hash.Hash `json:"TaskTips"`
}
type PlanetFesTaskAccessor struct {
	_data       []PlanetFesTask
	_dataTaskID map[float64]PlanetFesTask
}

// LoadData retrieves the data. Must be called before PlanetFesTask.GroupData
func (a *PlanetFesTaskAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PlanetFesTask.json")
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
func (a *PlanetFesTaskAccessor) Raw() ([]PlanetFesTask, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PlanetFesTask{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PlanetFesTaskAccessor.LoadData to preload everything
func (a *PlanetFesTaskAccessor) GroupData() {
	a._dataTaskID = map[float64]PlanetFesTask{}
	for _, d := range a._data {
		a._dataTaskID[d.TaskID] = d
	}
}

// ByTaskID returns the PlanetFesTask uniquely identified by TaskID
//
// Error is only non-nil if the source errors out
func (a *PlanetFesTaskAccessor) ByTaskID(identifier float64) (PlanetFesTask, error) {
	if a._dataTaskID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlanetFesTask{}, err
			}
		}
		a.GroupData()
	}
	return a._dataTaskID[identifier], nil
}
