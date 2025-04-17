package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type MuseumStats struct {
	AreaID     float64 `json:"AreaID"`
	FundCost   float64 `json:"FundCost"`
	Level      float64 `json:"Level"`
	PhaseLimit float64 `json:"PhaseLimit"`
	StatsType  float64 `json:"StatsType"`
	StatsValue float64 `json:"StatsValue"`
}
type MuseumStatsAccessor struct {
	_data []MuseumStats
}

// LoadData retrieves the data. Must be called before MuseumStats.GroupData
func (a *MuseumStatsAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MuseumStats.json")
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
func (a *MuseumStatsAccessor) Raw() ([]MuseumStats, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MuseumStats{}, err
		}
	}
	return a._data, nil
}
