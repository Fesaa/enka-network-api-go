package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type MuseumMission struct {
	MuseumMissionID float64   `json:"MuseumMissionID"`
	Type            string    `json:"Type"`
	TypeParameter   []float64 `json:"TypeParameter"`
}
type MuseumMissionAccessor struct {
	_data                []MuseumMission
	_dataMuseumMissionID map[float64]MuseumMission
}

// LoadData retrieves the data. Must be called before MuseumMission.GroupData
func (a *MuseumMissionAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MuseumMission.json")
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
func (a *MuseumMissionAccessor) Raw() ([]MuseumMission, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MuseumMission{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MuseumMissionAccessor.LoadData to preload everything
func (a *MuseumMissionAccessor) GroupData() {
	for _, d := range a._data {
		a._dataMuseumMissionID[d.MuseumMissionID] = d
	}
}

// ByMuseumMissionID returns the MuseumMission uniquely identified by MuseumMissionID
//
// Error is only non-nil if the source errors out
func (a *MuseumMissionAccessor) ByMuseumMissionID(identifier float64) (MuseumMission, error) {
	if a._dataMuseumMissionID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MuseumMission{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMuseumMissionID[identifier], nil
}
