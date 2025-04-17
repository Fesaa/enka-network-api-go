package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type MuseumMission struct {
	MuseumMissionID json.Number   `json:"MuseumMissionID"`
	Type            string        `json:"Type"`
	TypeParameter   []json.Number `json:"TypeParameter"`
}
type MuseumMissionAccessor struct {
	_data []MuseumMission
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
