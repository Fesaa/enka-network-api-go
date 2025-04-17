package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type PlanetFesRaiseConfig struct {
	GoldCost     float64 `json:"GoldCost"`
	RaiseCurveID float64 `json:"RaiseCurveID"`
	RaiseValue   float64 `json:"RaiseValue"`
}
type PlanetFesRaiseConfigAccessor struct {
	_data []PlanetFesRaiseConfig
}

// LoadData retrieves the data. Must be called before PlanetFesRaiseConfig.GroupData
func (a *PlanetFesRaiseConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PlanetFesRaiseConfig.json")
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
func (a *PlanetFesRaiseConfigAccessor) Raw() ([]PlanetFesRaiseConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PlanetFesRaiseConfig{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}
