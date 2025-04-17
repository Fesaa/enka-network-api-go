package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type MuseumArea struct {
	AreaID        float64 `json:"AreaID"`
	FundCost      float64 `json:"FundCost"`
	Level         float64 `json:"Level"`
	PhaseLimit    float64 `json:"PhaseLimit"`
	RenewPoint    float64 `json:"RenewPoint"`
	RequireStatsA float64 `json:"RequireStatsA"`
	RequireStatsB float64 `json:"RequireStatsB"`
	RequireStatsC float64 `json:"RequireStatsC"`
}
type MuseumAreaAccessor struct {
	_data []MuseumArea
}

// LoadData retrieves the data. Must be called before MuseumArea.GroupData
func (a *MuseumAreaAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MuseumArea.json")
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
func (a *MuseumAreaAccessor) Raw() ([]MuseumArea, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MuseumArea{}, err
		}
	}
	return a._data, nil
}
