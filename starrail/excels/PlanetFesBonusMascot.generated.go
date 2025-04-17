package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type PlanetFesBonusMascot struct {
	ID         json.Number `json:"ID"`
	SourcePath string      `json:"SourcePath"`
}
type PlanetFesBonusMascotAccessor struct {
	_data []PlanetFesBonusMascot
}

// LoadData retrieves the data. Must be called before PlanetFesBonusMascot.GroupData
func (a *PlanetFesBonusMascotAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PlanetFesBonusMascot.json")
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
func (a *PlanetFesBonusMascotAccessor) Raw() ([]PlanetFesBonusMascot, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PlanetFesBonusMascot{}, err
		}
	}
	return a._data, nil
}
